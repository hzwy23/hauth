package models

import (
	"encoding/json"

	"errors"

	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

const (
	error_user_role_no_auth string = "您没有权限删除这个域中用户的角色信息"
	error_user_role         string = "执行SQL失败"
	error_user_role_commit  string = "提交请求失败"
	error_user_role_begin   string = "初始化事物处理失败"
	error_user_role_json    string = "json数据解析失败"
)

type UserRolesModel struct {
}

type userRoleData struct {
	User_id     string `json:"user_id"`
	Role_id     string `json:"role_id"`
	Code_number string `json:"code_number"`
	Role_name   string `json:"role_name"`
	Role_status string `json:"role_status"`
}

// 根据用户id,获取这个用户已经拥有的角色
func (UserRolesModel) GetRolesByUser(user_id string) ([]userRoleData, error) {
	rows, err := dbobj.Query(sys_rdbms_094, user_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []userRoleData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 获取这个用户id,还没有获取的角色信息
func (UserRolesModel) GetOtherRoles(user_id string) ([]userRoleData, error) {
	rows, err := dbobj.Query(sys_rdbms_095, user_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []userRoleData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 对这个域中的用户进行授权
func (UserRolesModel) Auth(domain_id, user_id, ijs string) (string, error) {

	var rst []userRoleData

	err := json.Unmarshal([]byte(ijs), &rst)

	if err != nil {
		logs.Error(err)
		return error_user_role_json, err
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return error_user_role_begin, err
	}

	for _, val := range rst {
		did, err := hrpc.GetDomainId(val.User_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_user_role_no_auth, errors.New(error_user_role_no_auth)
		}

		if domain_id != did && user_id != "admin" {
			level := hrpc.GetDomainAuth(user_id, did)
			if level != 2 {
				tx.Rollback()
				return error_user_role_no_auth, errors.New(error_user_role_no_auth)
			}
		}

		_, err = tx.Exec(sys_rdbms_096, val.Role_id, val.User_id, user_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_user_role, err
		}
	}
	return error_user_role_commit, tx.Commit()
}

// 移除这个用户拥有的角色信息
func (UserRolesModel) Revoke(user_id string, role_id string, uid, did string) (string, error) {
	u_domain_id, err := hrpc.GetDomainId(user_id)
	if err != nil {
		return error_user_role_no_auth, err
	}
	if u_domain_id != did && uid != "admin" {
		level := hrpc.GetDomainAuth(uid, u_domain_id)
		if level != 2 {
			return error_user_role_no_auth, errors.New(error_user_role_no_auth)
		}
	}
	_, err = dbobj.Exec(sys_rdbms_097, user_id, role_id)
	return error_user_role_commit, err
}
