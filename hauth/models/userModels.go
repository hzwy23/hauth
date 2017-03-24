package models

import (
	"encoding/json"
	"errors"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

const (
	error_user_forbid_delete_admin = "admin 是系统超级管理员，无法被删除"
	error_user_forbid_yourself     = "禁止将自己从系统中删除"
	error_user_json                = "解析JSON数据失败"
	error_user_query_org           = "查询用户机构所属域信息失败"
	error_user_no_auth             = "您没有权限操作这个域中的用户"
	error_user_exec                = "执行SQL语句失败"
	error_user_commit              = "提交事务初始失败"
	error_user_modify_status       = "修改用户状态信息失败"
	error_user_modify_passwd       = "修改用户密码失败"
	error_user_modify_info         = "修改用户信息失败"
)

type UserModel struct {
}

type userInfo struct {
	User_id             string `json:"user_id"`
	User_name           string `json:"user_name"`
	User_status_desc    string `json:"status_desc"`
	User_create_date    string `json:"create_date"`
	User_owner          string `json:"create_user"`
	User_email          string `json:"user_email"`
	User_phone          string `json:"user_phone"`
	Org_unit_id         string `json:"org_unit_id"`
	Org_unit_desc       string `json:"org_unit_desc"`
	Domain_id           string `json:"domain_id"`
	Domain_name         string `json:"domain_name"`
	User_maintance_date string `json:"modify_date"`
	User_maintance_user string `json:"modify_user"`
	User_status_id      string `json:"status_cd"`
}

// 查询用户自己的详细信息
func (UserModel) GetOwnerDetails(user_id string) ([]userInfo, error) {
	row, err := dbobj.Query(sys_rdbms_023, user_id)
	defer row.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []userInfo
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 查询怒
func (UserModel) GetDefault(domain_id string) ([]userInfo, error) {

	row, err := dbobj.Query(sys_rdbms_017, domain_id)
	defer row.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []userInfo
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 新增用户信息
func (UserModel) Post(userId, userPasswd, userDesc, userStatus, id, userEmail, userPhone, userOrgUnitId string) error {
	tx, err := dbobj.Begin()
	// insert user details
	//
	_, err = tx.Exec(sys_rdbms_018, userId, userDesc, id, userEmail, userPhone, userOrgUnitId, id)
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return err
	}

	// insert user passwd
	//
	_, err = tx.Exec(sys_rdbms_019, userId, userPasswd, userStatus)
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return err
	}

	// insert theme info
	//

	stheme := `insert into sys_user_theme(user_id,theme_id) values(?,?)`

	_, err = tx.Exec(stheme, userId, "1001")
	if err != nil {
		tx.Rollback()
		logs.Error(err.Error())
		return err
	}

	return tx.Commit()
}

// 删除用户信息
func (UserModel) Delete(ijs []byte, user_id, domain_id string) (string, error) {
	var js []userInfo
	err := json.Unmarshal(ijs, &js)
	if err != nil {
		logs.Error(err)
		return error_user_json, err
	}

	tx, _ := dbobj.Begin()
	for _, val := range js {
		//判断用户是否在线
		//如果在线,则不允许删除用户
		if val.User_id == "admin" {
			tx.Rollback()
			return error_user_forbid_delete_admin, errors.New(error_user_forbid_delete_admin)
		}

		// check user
		// can't delete yourself
		if user_id == val.User_id {
			tx.Rollback()
			return error_user_forbid_yourself, errors.New("禁止将自己删除。")
		}

		// query domain_id by org_unit_id
		did, err := CheckDomainByOrgId(val.Org_unit_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_user_query_org, errors.New(error_user_query_org)
		}

		if user_id != "admin" && domain_id != did {
			level := CheckDomainRights(user_id, did)
			if level != 2 {
				logs.Error("没有被授权删除这个域中的信息")
				tx.Rollback()
				return error_user_no_auth, errors.New(error_user_no_auth)
			}
		}

		_, err = tx.Exec(sys_rdbms_007, val.User_id)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return error_user_exec, err
		}
	}
	return error_user_commit, tx.Commit()
}

// 搜索用户信息
func (this UserModel) Search(org_id string, status_id string, domain_id string) ([]userInfo, error) {
	var rst []userInfo
	var err error
	// 如果机构号为空
	// 直接查询指定域中所有的用户
	if org_id == "" {
		rst, err = this.GetDefault(domain_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else {
		rows, err := dbobj.Query(sys_rdbms_090, org_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}

		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	}

	if status_id == "" {
		return rst, nil
	} else if status_id == "0" {
		var ret []userInfo
		for _, val := range rst {
			if val.User_status_id == "0" {
				ret = append(ret, val)
			}
		}
		return ret, nil
	} else {
		var ret []userInfo
		for _, val := range rst {
			if val.User_status_id == "1" {
				ret = append(ret, val)
			}
		}
		return ret, nil
	}
}

func (this UserModel) ModifyStatus(status_id, user_id string) (string, error) {
	err := dbobj.Exec(sys_rdbms_016, status_id, user_id)
	return error_user_modify_status, err
}

func (this UserModel) ModifyPasswd(passwd, user_id string) (string, error) {
	err := dbobj.Exec(sys_rdbms_020, passwd, user_id)
	return error_user_modify_passwd, err
}

// 修改用户信息
func (this UserModel) Put(user_name, org_id, phone, email, uid, user_id string) (string, error) {
	err := dbobj.Exec(sys_rdbms_021, user_name, phone, email, uid, org_id, user_id)
	return error_user_modify_info, err
}
