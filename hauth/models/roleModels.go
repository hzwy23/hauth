package models

import (
	"errors"
	"net/url"

	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/validator"
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

// 查询某一个角色的具体信息
func (this RoleModel) GetRow(role_id string) (RoleInfo, error) {
	var rst RoleInfo
	domain_id, err := utils.SplitDomain(role_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	ret, err := this.Get(domain_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}

	for _, val := range ret {
		if val.Role_id == role_id {
			return val, nil
		}
	}
	return rst, errors.New("no found")
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

func (RoleModel) Post(data url.Values, user_id string) (string, error) {
	domainid := data.Get("domain_id")
	roleid := data.Get("role_id")
	rolename := data.Get("role_name")
	rolestatus := data.Get("role_status")
	id := utils.JoinCode(domainid, roleid)

	//校验
	if !validator.IsAlnum(roleid) {
		return "error_role_id_format", errors.New("error_role_id_format")
	}
	//
	if validator.IsEmpty(rolename) {
		return "error_role_desc_empty", errors.New("error_role_desc_empty")
	}

	if !validator.IsWord(domainid) {
		return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
	}

	if !validator.IsIn(rolestatus, "0", "1") {
		return "error_role_status", errors.New("error_role_status")
	}

	_, err := dbobj.Exec(sys_rdbms_026, id, rolename, user_id, rolestatus, domainid, user_id, roleid)
	if err != nil {
		logs.Error(err)
		return "error_role_add_failed", err
	}
	return "success", nil
}

func (RoleModel) Delete(allrole []RoleInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_sql_begin", err
	}

	for _, val := range allrole {
		_, err := tx.Exec(sys_rdbms_027, val.Role_id, val.Domain_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_role_delete_failed", err
		}
		logs.Info("delete role info successfully. role id is :", val.Role_id)
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_role_delete_failed", err
	}
	return "success", nil
}

func (RoleModel) Update(data url.Values, user_id string) (string, error) {
	Role_id := data.Get("Role_id")
	Role_name := data.Get("Role_name")
	Role_status := data.Get("Role_status")

	if !validator.IsWord(Role_id) {
		return "error_role_id_format", errors.New("error_role_id_format")
	}

	if validator.IsEmpty(Role_name) {
		return "error_role_desc_empty", errors.New("error_role_desc_empty")
	}

	if !validator.IsIn(Role_status, "0", "1") {
		return "error_role_status", errors.New("error_role_status")
	}

	_, err := dbobj.Exec(sys_rdbms_050, Role_name, Role_status, user_id, Role_id)
	if err != nil {
		logs.Error(err)
		return "error_role_update_failed", errors.New("error_role_update_failed")
	}
	return "success", nil
}
