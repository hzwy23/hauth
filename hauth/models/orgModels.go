package models

import (

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"time"
	"errors"
)

type OrgModel struct {
}

type SysOrgInfo struct {
	Org_unit_id     string `json:"org_id"`
	Org_unit_desc   string `json:"org_desc"`
	Up_org_id       string `json:"up_org_id"`
	Org_status_id   string `json:"status_id"`
	Org_status_desc string `json:"status_desc"`
	Domain_id       string `json:"domain_id"`
	Create_date     string `json:"create_date"`
	Maintance_date  string `json:"modify_date"`
	Create_user     string `json:"create_user"`
	Maintance_user  string `json:"modify_user"`
	Code_number     string `json:"code_number"`
	Org_dept        string `json:"org_dept,omitempty"`
}

//获取域下边所有机构号
func (OrgModel) Get(domain_id string) ([]SysOrgInfo, error) {
	key := hcache.GenKey("ASOFDATEORGINFO",domain_id)
	if hcache.IsExist(key) {
		logs.Debug("get org info fromo cache.")
		rst,_ := hcache.Get(key).([]SysOrgInfo)
		return rst,nil
	}

	var rst []SysOrgInfo
	rows, err := dbobj.Query(sys_rdbms_041, domain_id)
	if err != nil {
		return nil, err
	}

	err = dbobj.Scan(rows, &rst)
	if err != nil {
		return nil, err
	}

	hcache.Put(key,rst,720*time.Minute)

	return rst, nil
}

func (OrgModel) Delete(mjs []SysOrgInfo) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return errors.New("开启事物处理失败,请联系管理员")
	}

	for _, val := range mjs {
		_, err := tx.Exec(sys_rdbms_044, val.Org_unit_id, val.Domain_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return errors.New("这个机构下边存在用户信息,请先清除用户信息,再来删除机构")
		}
		hcache.Delete(hcache.GenKey("ASOFDATEORGINFO",val.Domain_id))
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return errors.New("提交删除请求失败,请联系管理员")
	}
	return nil
}

func (OrgModel) Update(org_unit_desc, up_org_id, org_status_id, maintance_user, org_unit_id,domain_id string) error {
	defer hcache.Delete(hcache.GenKey("ASOFDATEORGINFO",domain_id))
	return dbobj.Exec(sys_rdbms_069, org_unit_desc, up_org_id, org_status_id, maintance_user, org_unit_id)
}

func (OrgModel) Post(org_unit_id, org_unit_desc, up_org_id, org_status_id, domain_id, create_user, maintance_user, id string) error {
	defer hcache.Delete(hcache.GenKey("ASOFDATEORGINFO",domain_id))
	return dbobj.Exec(sys_rdbms_043, org_unit_id, org_unit_desc, up_org_id, org_status_id,
		domain_id, create_user, maintance_user, id)
}

func (this OrgModel) GetSubOrgInfo(domain_id string,org_id string) ([]SysOrgInfo, error) {
	var rst []SysOrgInfo

	all, err := this.Get(domain_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	// 将自身机构加入到结果中.
	for _,val:=range all {
		if val.Org_unit_id == org_id {
			rst = append(rst,val)
			break
		}
	}

	this.dfs(all,org_id,&rst)

	return rst, nil
}

func (this OrgModel)dfs(node []SysOrgInfo,org_id string ,rst *[]SysOrgInfo) {
	for _, val := range node {
		if val.Up_org_id == org_id {
			*rst = append(*rst,val)
			this.dfs(node,val.Org_unit_id,rst)
		}
	}
}
