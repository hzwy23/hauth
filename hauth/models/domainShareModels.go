package models

import (
	"errors"

	"net/url"

	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logs"
)

type DomainShareModel struct {
	md DomainMmodel
}

type DomainShareData struct {
	Uuid                string `json:"uuid"`
	Target_domain_id    string `json:"target_domain_id"`
	Domain_name         string `json:"domain_name"`
	Authorization_level string `json:"auth_level"`
	Create_user         string `json:"create_user"`
	Create_date         string `json:"create_date"`
	Modify_user         string `json:"modify_user"`
	Modify_date         string `json:"modify_date"`
}

type dusModel struct {
	Domain_id   string `json:"domain_id"`
	Domain_name string `json:"domain_name"`
}

// 获取指定域共享给了哪些对象
func (DomainShareModel) Get(domain_id string) ([]DomainShareData, error) {

	rows, err := dbobj.Query(sys_rdbms_083, domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []DomainShareData

	err = dbobj.Scan(rows, &rst)

	return rst, err
}

func (DomainShareModel) UnAuth(domain_id string) ([]dusModel, error) {
	rows, err := dbobj.Query(sys_rdbms_085, domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []dusModel
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (DomainShareModel) Post(data url.Values, user_id string) (string, error) {

	domain_id := data.Get("domain_id")
	target_domain_id := data.Get("target_domain_id")
	auth_level := data.Get("auth_level")
	fmt.Println(domain_id)
	if !govalidator.IsAlnum(domain_id) {
		return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
	}

	if !govalidator.IsAlnum(target_domain_id) {
		return "as_of_date_domain_target", errors.New("as_of_date_domain_target")
	}

	if !govalidator.IsIn(auth_level, "1", "2") {
		return "as_of_date_domain_mode", errors.New("as_of_date_domain_mode")
	}
	_, err := dbobj.Exec(sys_rdbms_086, domain_id, target_domain_id, auth_level, user_id, user_id)
	if err != nil {
		logs.Error(err)
		return "as_of_date_domain_share_failed", err
	}
	return "success", nil
}

func (DomainShareModel) Update(data url.Values, user_id string) (string, error) {

	domain_id := data.Get("domain_id")
	uuid := data.Get("uuid")
	level := data.Get("auth_level")

	if !govalidator.IsWord(domain_id) {
		return "as_of_date_domain_target", errors.New("as_of_date_domain_target")
	}

	if !govalidator.IsIn(level, "1", "2") {
		return "as_of_date_domain_mode", errors.New("as_of_date_domain_mode")
	}

	_, err := dbobj.Exec(sys_rdbms_088, level, user_id, uuid)
	if err != nil {
		logs.Error(err)
		return "as_of_date_domain_share_update", errors.New("as_of_date_domain_share_update")
	}
	return "success", nil
}

func (DomainShareModel) Delete(data []DomainShareData, domain_id string) (string, error) {

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_sql_begin", err
	}

	for _, val := range data {
		_, err := dbobj.Exec(sys_rdbms_087, val.Uuid, domain_id)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return "as_of_date_domain_share_delete", err
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "as_of_date_domain_share_delete", errors.New("as_of_date_domain_share_delete")
	}
	return "success", nil
}

// 获取这个有哪些域把共享给了指定的这个域
func (this DomainShareModel) get(domain_id string) ([]DomainMmodel, error) {
	var rst []DomainMmodel
	rows, err := dbobj.Query(sys_rdbms_034, domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this DomainShareModel) GetList(domain_id string) ([]DomainMmodel, error) {
	// 获取所有的域信息
	rst, err := this.md.Get()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 获取指定域能够访问到的域信息
	ret, err := this.get(domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var dmap = make(map[string]bool)
	dmap[domain_id] = true
	for _, val := range ret {
		dmap[val.Project_id] = true
	}

	var dslice []DomainMmodel
	for _, val := range rst {
		if _, ok := dmap[val.Project_id]; ok && val.Domain_status_cd == "0" {
			dslice = append(dslice, val)
		}
	}
	return dslice, nil
}

func (this DomainShareModel) GetOwner(domain_id string) (domainDataSet, error) {

	var ret domainDataSet

	rst, err := this.GetList(domain_id)
	if err != nil {
		logs.Error(err)
		return ret, err
	}

	ret.Domain_id = domain_id
	ret.Owner_list = rst

	return ret, nil
}
