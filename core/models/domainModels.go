package models

import (
	"database/sql"
	"errors"
	"net/url"

	"github.com/hzwy23/hauth/utils/logs"
	"github.com/hzwy23/hauth/utils/validator"
	"github.com/hzwy23/dbobj"
)

type DomainMmodel struct {
	Project_id            string `json:"domain_id"`
	Project_name          string `json:"domain_desc"`
	Project_status        string `json:"domain_status"`
	Maintance_date        string `json:"maintance_date" dateType:"YYYY-MM-DD HH24:MM:SS"`
	User_id               string `json:"create_user_id"`
	Domain_maintance_date string `json:"domain_modify_date" dateType:"YYYY-MM-DD HH24:MM:SS"`
	Domain_maintance_user string `json:"domain_modify_user"`
	Domain_status_cd      string `json:"domain_status_cd"`
}

type domainDataSet struct {
	Domain_id  string         `json:"domain_id"`
	Owner_list []DomainMmodel `json:"owner_list"`
}

func (this DomainMmodel) Get() ([]DomainMmodel, error) {
	var rst []DomainMmodel
	rows, err := dbobj.Query(sys_rdbms_025)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	//	var oneLine DomainMmodel
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return rst, nil
}

// 获取指定域的详细信息
func (DomainMmodel) GetRow(domain_id string) (DomainMmodel, error) {
	var rst DomainMmodel

	row, err := dbobj.Query(sys_rdbms_084, domain_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	for row.Next() {
		var domain_id sql.NullString
		var domain_desc sql.NullString
		var domain_status sql.NullString
		var modify_date sql.NullString
		var modify_user sql.NullString
		var create_date sql.NullString
		var create_user sql.NullString
		err := row.Scan(&domain_id,
			&domain_desc,
			&domain_status,
			&create_date,
			&create_user,
			&modify_date,
			&modify_user)
		if err != nil {
			logs.Error(err)
			return rst, err
		}

		rst.Project_id = domain_id.String
		rst.Project_name = domain_desc.String
		rst.Project_status = domain_status.String
		rst.Maintance_date = create_date.String
		rst.User_id = create_user.String
		rst.Domain_maintance_date = modify_date.String
		rst.Domain_maintance_user = modify_user.String
		return rst, nil
	}
	return rst, errors.New("no value")
}

// 新增域信息
// 并将新增的域授权个创建人
func (DomainMmodel) Post(data url.Values, user_id string, user_domain_id string) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		return "error_sql_begin", err
	}

	// Get the form data
	domain_id := data.Get("domainId")
	domain_desc := data.Get("domainDesc")
	domain_status := data.Get("domainStatus")

	// validator domain id format
	if !validator.IsAlnum(domain_id) {
		return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
	}

	// validator domain describe format. It does not allow null values
	if validator.IsEmpty(domain_desc) {
		return "as_of_date_domain_isempty", errors.New("as_of_date_domain_isempty")
	}

	// validator domain status format
	// It must be in the 0 and 1
	if !validator.IsIn(domain_status, "0", "1") {
		return "as_of_date_domain_status_check", errors.New("as_of_date_domain_status_check")
	}

	_, err = tx.Exec(sys_rdbms_036, domain_id, domain_desc, domain_status, user_id, user_id)
	if err != nil {
		tx.Rollback()
		return "as_of_date_domain_add_failed", err
	}

	_, err = tx.Exec(sys_rdbms_086, domain_id, user_domain_id, 2, user_id, user_id)
	if err != nil {
		tx.Rollback()
		return "as_of_date_domain_add_failed", err
	}

	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "as_of_date_domain_add_failed", err
	}
	return "success", nil
}

// 删除域信息
// 在controller中校验权限
func (DomainMmodel) Delete(js []DomainMmodel) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range js {
		_, err := tx.Exec(sys_rdbms_037, val.Project_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 更新域信息
// 只能更新名称和状态
func (DomainMmodel) Update(data url.Values, user_id string) (string, error) {

	domainId := data.Get("domainId")
	domainDesc := data.Get("domainDesc")
	domainStatus := data.Get("domainStatus")

	// 校验域名称,不能为空
	if validator.IsEmpty(domainDesc) {
		return "as_of_date_domain_isempty", errors.New("as_of_date_domain_isempty")
	}

	// 校验域状态编码,必须是0或者1
	if !validator.IsIn(domainStatus, "0", "1") {
		return "as_of_date_domain_status_check", errors.New("as_of_date_domain_status_check")
	}

	_, err := dbobj.Exec(sys_rdbms_038, domainDesc, domainStatus, user_id, domainId)
	if err != nil {
		logs.Error(err)
		return "as_of_date_domain_update", err
	}
	return "success", nil
}
