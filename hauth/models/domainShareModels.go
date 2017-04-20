package models

import (
	"encoding/json"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

type DomainShareModel struct {
	md  ProjectMgr
}

type dsModel struct {
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
func (DomainShareModel) Get(domain_id string) ([]dsModel, error) {

	rows, err := dbobj.Query(sys_rdbms_083, domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rst []dsModel

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

func (DomainShareModel) Post(domain_id, target_domain_id, auth_level, user_id string) error {
	_,err := dbobj.Exec(sys_rdbms_086, domain_id, target_domain_id, auth_level, user_id, user_id)
	return err
}

func (DomainShareModel) Update(uuid, user_id, auth_level string) error {
	_,err := dbobj.Exec(sys_rdbms_088, auth_level, user_id, uuid)
	return err
}

func (DomainShareModel) Delete(js string,domain_id string) error {
	var rst []dsModel

	err := json.Unmarshal([]byte(js), &rst)
	if err != nil {
		logs.Error(err)
		return err
	}
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}

	for _, val := range rst {
		_,err := dbobj.Exec(sys_rdbms_087, val.Uuid,domain_id)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return err
		}
	}
	return tx.Commit()
}

// 获取这个有哪些域把共享给了指定的这个域
func (this DomainShareModel)get(domain_id string)([]ProjectMgr,error){
	var rst []ProjectMgr
	rows,err := dbobj.Query(sys_rdbms_034,domain_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}
	err = dbobj.Scan(rows,&rst)
	return rst,err
}

func (this DomainShareModel) GetList(domain_id string) ([]ProjectMgr, error) {
	// 获取所有的域信息
	rst,err := this.md.Get()
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	// 获取指定域能够访问到的域信息
	ret,err := this.get(domain_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	var dmap  = make(map[string]bool)
	dmap[domain_id]=true
	for _, val := range ret {
		dmap[val.Project_id]=true
	}

	var dslice []ProjectMgr
	for _,val := range rst {
		if _,ok := dmap[val.Project_id]; ok && val.Domain_status_cd == "0"{
			dslice = append(dslice,val)
		}
	}
	return dslice, nil
}

func (this DomainShareModel) GetOwner(domain_id string) (domainDataModel, error) {

	var ret domainDataModel

	rst,err := this.GetList(domain_id)
	if err != nil {
		logs.Error(err)
		return ret,err
	}

	ret.Domain_id = domain_id
	ret.Owner_list = rst

	return ret, nil
}
