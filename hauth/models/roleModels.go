package models

import (
	"errors"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

type RoleModel struct {
}

type RoleInfo struct {
	Code_number         string `json:"code_number"`
	Role_name           string `json:"role_name"`
	Role_owner          string `json:"create_user"`
	Role_create_date    string `json:"create_date"`
	Role_status_desc    string `json:"role_status_desc"`
	Role_status         string `json:"role_status_code"`
	Domain_id           string `json:"domain_id"`
	Domain_desc         string `json:"domain_desc"`
	Role_maintance_date string `json:"modify_date"`
	Role_maintance_user string `json:"modify_user"`
	Role_id             string `json:"role_id"`
}

func (RoleModel) GetRow(role_id string) ([]RoleInfo, error) {
	var rst []RoleInfo
	rows, err := dbobj.Query(sys_rdbms_091, role_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (RoleModel) Get(domain_id string) ([]RoleInfo, error) {
	rows, err := dbobj.Query(sys_rdbms_028, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []RoleInfo
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (RoleModel) Post(id, rolename, user_id, rolestatus, domainid, roleid string) error {
	return dbobj.Exec(sys_rdbms_026, id, rolename, user_id, rolestatus, domainid, user_id, roleid)
}

func (RoleModel) Delete(allrole []RoleInfo, user_id, domain_id string) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}

	for _, val := range allrole {

		if val.Domain_id != domain_id && user_id != "admin" {
			level := CheckDomainRights(user_id, val.Domain_id)
			if level != 2 {
				tx.Rollback()
				return errors.New("您没有权限删除这个域中的角色信息")
			}
		}

		_, err := tx.Exec(sys_rdbms_027, val.Role_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
		logs.Info("delete role info successfully. role id is :", val.Role_id)
	}
	return tx.Commit()
}

func (RoleModel) Update(Role_name, Role_status, Role_id, User_id string) error {
	return dbobj.Exec(sys_rdbms_050, Role_name, Role_status, User_id, Role_id)
}
