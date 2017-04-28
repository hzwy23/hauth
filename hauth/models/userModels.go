package models

import (
	"errors"
	"net/url"
	"strings"

	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/crypto/haes"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/validator"
)

type UserModel struct {
	morg OrgModel
}

type UserInfo struct {
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
func (UserModel) GetOwnerDetails(user_id string) ([]UserInfo, error) {
	var rst []UserInfo
	row, err := dbobj.Query(sys_rdbms_023, user_id)
	defer row.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 查询域中所有的用户信息
func (UserModel) GetDefault(domain_id string) ([]UserInfo, error) {
	row, err := dbobj.Query(sys_rdbms_017, domain_id)
	defer row.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []UserInfo
	err = dbobj.Scan(row, &rst)
	return rst, err
}

// 新增用户信息
func (UserModel) Post(data url.Values, user_id string) (string, error) {
	userId := data.Get("userId")
	userDesc := data.Get("userDesc")
	password := data.Get("userPasswd")
	surepassword := data.Get("userPasswdConfirm")
	userStatus := data.Get("userStatus")
	userEmail := data.Get("userEmail")
	userPhone := data.Get("userPhone")
	userOrgUnitId := data.Get("userOrgUnitId")
	domain_id := data.Get("domainId")

	if !validator.IsWord(userId) {
		return "error_user_id_check", errors.New("error_user_id_check")
	}
	//

	if validator.IsEmpty(userDesc) {
		return "error_user_name_check", errors.New("error_user_name_check")
	}

	if validator.IsEmpty(password) {
		return "error_user_passwd_check", errors.New("error_user_passwd_check")
	}

	if validator.IsEmpty(surepassword) {
		return "error_passwd_empty", errors.New("error_passwd_empty")
	}

	if password != surepassword {
		return "error_passwd_confirm_failed", errors.New("error_passwd_confirm_failed")
	}

	if len(strings.TrimSpace(password)) < 6 {
		return "error_passwd_short", errors.New("error_passwd_short")
	}

	userPasswd, err := haes.Encrypt(password)
	if err != nil {
		logs.Error(err)
		return "error_user_passwd_encrypt", errors.New("error_user_passwd_encrypt")
	}

	//
	if !validator.IsEmail(userEmail) {
		return "error_user_email_check", errors.New("error_user_email_check")
	}

	if !validator.IsWord(userOrgUnitId) {
		return "error_user_role_org", errors.New("error_user_role_org")
	}

	//
	if !validator.IsMobilePhone(userPhone) {
		return "error_user_phone_check", errors.New("error_user_phone_check")
	}

	org_domain_id, _ := utils.SplitDomain(userOrgUnitId)
	if domain_id != org_domain_id {
		return "error_user_org_format", errors.New("error_user_org_format")
	}

	tx, err := dbobj.Begin()
	// insert user details
	//
	_, err = tx.Exec(sys_rdbms_018, userId, userDesc, user_id, userEmail, userPhone, userOrgUnitId, user_id)
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return "error_user_post", err
	}

	// insert user passwd
	_, err = tx.Exec(sys_rdbms_019, userId, userPasswd, userStatus)
	if err != nil {
		tx.Rollback()
		logs.Error(err)
		return "error_user_post", err
	}

	// insert theme info
	_, err = tx.Exec(sys_rdbms_045, userId, "1001")
	if err != nil {
		tx.Rollback()
		logs.Error(err.Error())
		return "error_user_post", err
	}

	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_user_post", err
	}
	return "success", nil
}

// 删除用户信息
func (UserModel) Delete(data []UserInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		return "error_sql_begin", err
	}

	for _, val := range data {
		_, err = tx.Exec(sys_rdbms_007, val.User_id, val.Org_unit_id)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return "error_user_exec", err
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_user_commit", err
	}

	return "success", nil
}

// 搜索用户信息
func (this UserModel) Search(org_id string, status_id string, domain_id string) ([]UserInfo, error) {
	var rst []UserInfo

	ret, err := this.GetDefault(domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if !validator.IsEmpty(org_id) {

		orglist, err := this.morg.GetSubOrgInfo(domain_id, org_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}

		var orgmap map[string]string = make(map[string]string)
		for _, val := range orglist {
			orgmap[val.Org_unit_id] = ""
		}

		for _, val := range ret {
			if _, ok := orgmap[val.Org_unit_id]; ok {
				if !validator.IsEmpty(status_id) {
					if val.User_status_id == status_id {
						rst = append(rst, val)
					}
				} else {
					rst = append(rst, val)
				}
			}
		}
	} else {
		for _, val := range ret {
			if !validator.IsEmpty(status_id) {
				if val.User_status_id == status_id {
					rst = append(rst, val)
				}
			} else {
				rst = append(rst, val)
			}
		}
	}
	return rst, nil
}

func (this UserModel) ModifyStatus(status_id, user_id string) (string, error) {
	if !validator.IsIn(status_id, "0", "1") {
		return "error_user_status_empty", errors.New("error_user_status_empty")
	}

	_, err := dbobj.Exec(sys_rdbms_016, status_id, user_id)
	return "error_user_modify_status", err
}

func (this UserModel) ModifyPasswd(data url.Values) (string, error) {
	user_id := data.Get("userid")
	user_password := data.Get("newpasswd")
	confirm_password := data.Get("surepasswd")
	if user_password != confirm_password {
		return "error_passwd_confirm_failed", errors.New("error_passwd_confirm_failed")
	}

	if len(strings.TrimSpace(confirm_password)) < 6 || len(strings.TrimSpace(confirm_password)) > 30 {
		return "error_passwd_short", errors.New("error_passwd_short")
	}

	encry_passwd, err := haes.Encrypt(user_password)
	if err != nil {
		logs.Error(err)
		return "error_password_encrpty", errors.New("error_password_encrpty")
	}

	_, err = dbobj.Exec(sys_rdbms_020, encry_passwd, user_id)
	if err != nil {
		logs.Error(err)
		return "error_user_modify_passwd", err
	}
	return "success", nil
}

// 修改用户信息
func (this UserModel) Put(data url.Values, modify_user string) (string, error) {
	user_name := data.Get("userDesc")
	org_id := data.Get("orgId")
	phone := data.Get("userPhone")
	email := data.Get("userEmail")
	user_id := data.Get("userId")

	if !validator.IsWord(user_id) {
		return "error_user_id_empty", errors.New("error_user_id_empty")
	}

	if validator.IsEmpty(user_name) {
		return "error_user_desc_empty", errors.New("error_user_desc_empty")
	}

	if !validator.IsEmail(email) {
		return "error_user_email_format", errors.New("error_user_email_format")
	}

	if !validator.IsWord(org_id) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	if !validator.IsMobilePhone(phone) {
		return "error_user_phone_format", errors.New("error_user_phone_format")
	}

	_, err := dbobj.Exec(sys_rdbms_021, user_name, phone, email, modify_user, org_id, user_id)
	if err != nil {
		logs.Error(err)
		return "error_user_modify_info", err
	}
	return "success", nil
}
