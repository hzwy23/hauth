package models

import (
	"errors"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

type ProjectMgr struct {
	Project_id            string `json:"domain_id"`
	Project_name          string `json:"domain_desc"`
	Project_status        string `json:"domain_status"`
	Maintance_date        string `json:"maintance_date"`
	User_id               string `json:"create_user_id"`
	Domain_maintance_date string `json:"domain_modify_date"`
	Domain_maintance_user string `json:"domain_modify_user"`
	Domain_dept           string `json:"domain_dept"`
	Domain_up_id          string `json:"domain_up_id"`
}

type domainDataModel struct {
	Domain_id  string       `json:"domain_id"`
	Owner_list []ProjectMgr `json:"owner_list"`
}

func (ProjectMgr) GetAll(offset, limit string) ([]ProjectMgr, int64, error) {
	rows, err := dbobj.Query(sys_rdbms_082, offset, limit)
	defer rows.Close()
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}

	//	var oneLine ProjectMgr
	var rst []ProjectMgr
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(sys_rdbms_081).Scan(&total)

	return rst, total, nil
}

func (ProjectMgr) GetRow(domain_id string) (ProjectMgr, error) {
	var rst ProjectMgr
	err := dbobj.QueryRow(sys_rdbms_084, domain_id).Scan(&rst.Project_id,
		&rst.Project_name, &rst.Project_status, &rst.Maintance_date, &rst.User_id, &rst.Domain_maintance_date, &rst.Domain_maintance_user)
	return rst, err
}

func (ProjectMgr) Get(domain_id string) ([]ProjectMgr, error) {
	rows, err := dbobj.Query(sys_rdbms_034, domain_id, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return nil, err
	}

	//	var oneLine ProjectMgr
	var rst []ProjectMgr
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return nil, err
	}

	return rst, nil
}

func (ProjectMgr) GetOwner(domain_id string) (domainDataModel, error) {
	var ret domainDataModel
	rows, err := dbobj.Query(sys_rdbms_034, domain_id, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return ret, err
	}

	//	var oneLine ProjectMgr
	var rst []ProjectMgr
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error("query data error.", dbobj.GetErrorMsg(err))
		return ret, err
	}

	ret.Domain_id = domain_id
	ret.Owner_list = rst
	return ret, nil
}

func (ProjectMgr) Post(domain_id, domain_desc, domain_status, user_id, did string) error {
	tx, err := dbobj.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(sys_rdbms_036, domain_id, domain_desc, domain_status, user_id, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sys_rdbms_086, domain_id, did, 2, user_id, user_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ProjectMgr) Delete(js []ProjectMgr, user_id string, domain_id string) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range js {
		if domain_id == val.Project_id {
			tx.Rollback()
			logs.Error(err)
			return errors.New("您不能删除自己所处的域。")
		}

		if user_id != "admin" {
			level := CheckDomainRights(user_id, val.Project_id)
			if level != 2 {
				tx.Rollback()
				logs.Error("您没有权限删除这个域")
				return errors.New("您没有权限删除这个域。")
			}
		}
		_, err := tx.Exec(sys_rdbms_037, val.Project_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (ProjectMgr) Update(domainDesc, domainStatus, user_id, domainId string) error {
	return dbobj.Exec(sys_rdbms_038, domainDesc, domainStatus, user_id, domainId)
}
