package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/logs"
)

type UserRolesModel struct {
	User_id     string `json:"user_id"`
	Role_id     string `json:"role_id"`
	Code_number string `json:"code_number"`
	Role_name   string `json:"role_name"`
	Role_status string `json:"role_status"`
}

// 根据用户id,获取这个用户已经拥有的角色
func (UserRolesModel) GetRolesByUser(user_id string) ([]UserRolesModel, error) {
	rows, err := dbobj.Query(sys_rdbms_094, user_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []UserRolesModel
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 获取这个用户id,还没有获取的角色信息
func (UserRolesModel) GetOtherRoles(user_id string) ([]UserRolesModel, error) {
	rows, err := dbobj.Query(sys_rdbms_095, user_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []UserRolesModel
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 对这个域中的用户进行授权
func (UserRolesModel) Auth(data []UserRolesModel, user_id string) (string, error) {

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_user_role_begin", err
	}

	for _, val := range data {
		uuid := utils.JoinCode(val.User_id, val.Role_id)
		_, err = tx.Exec(sys_rdbms_096, uuid, val.Role_id, val.User_id, user_id)
		if err != nil {
			logs.Info("用户【", val.User_id, "】已经拥有了角色【", val.Role_id, "】，无需重复授权。")
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_user_role_commit", err
	}
	return "success", nil
}

// 移除这个用户拥有的角色信息
func (UserRolesModel) Revoke(rst []UserRolesModel) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		return "error_sql_begin", err
	}

	for _, val := range rst {
		_, err := dbobj.Exec(sys_rdbms_097, val.User_id, val.Role_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_user_role_commit", err
		}
	}

	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_user_role", err
	}
	return "success", nil
}
