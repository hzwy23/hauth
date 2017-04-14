package models

import (
	"errors"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"time"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils"
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
	domain_id := utils.SplitDomain(role_id)
	ret,err := this.Get(domain_id)
	if err !=nil {
		logs.Error(err)
		return rst,err
	}

	for _, val := range ret {
		if val.Role_id == role_id {
			return val,nil
		}
	}
	return rst, errors.New("no found")
}

func (RoleModel) Get(domain_id string) ([]RoleInfo, error) {
	key := hcache.GenKey("ROLEMODELS",domain_id)
	if hcache.IsExist(key){
		logs.Debug("get data from cache.")
		rst , _ := hcache.Get(key).([]RoleInfo)
		return rst,nil
	}
	rows, err := dbobj.Query(sys_rdbms_028, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []RoleInfo
	err = dbobj.Scan(rows, &rst)
	hcache.Put(key,rst,720*time.Minute)
	return rst, err
}

func (RoleModel) Post(id, rolename, user_id, rolestatus, domainid, roleid string) error {
	defer hcache.Delete(hcache.GenKey("ROLEMODELS",domainid))
	_,err := dbobj.Exec(sys_rdbms_026, id, rolename, user_id, rolestatus, domainid, user_id, roleid)
	return err
}

func (RoleModel) Delete(allrole []RoleInfo, user_id, domain_id string) error {

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}

	for _, val := range allrole {

		if val.Domain_id != domain_id && user_id != "admin" {
			level := hrpc.CheckDomainRights(user_id, val.Domain_id)
			if level != 2 {
				tx.Rollback()
				return errors.New("您没有权限删除这个域中的角色信息")
			}
		}
		hcache.Delete(hcache.GenKey("ROLEMODELS",val.Domain_id))
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

func (RoleModel) Update(Role_name, Role_status, Role_id, User_id,domain_id string) error {
	defer hcache.Delete(hcache.GenKey("ROLEMODELS",domain_id))
	_,err := dbobj.Exec(sys_rdbms_050, Role_name, Role_status, User_id, Role_id)
	return err
}
