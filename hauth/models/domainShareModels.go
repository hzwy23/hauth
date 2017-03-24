package models

import (
	"encoding/json"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

type DomainShareModel struct {
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
	return dbobj.Exec(sys_rdbms_086, domain_id, target_domain_id, auth_level, user_id, user_id)
}

func (DomainShareModel) Update(uuid, user_id, auth_level string) error {
	return dbobj.Exec(sys_rdbms_088, auth_level, user_id, uuid)
}

func (DomainShareModel) Delete(js string) error {
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
		err := dbobj.Exec(sys_rdbms_087, val.Uuid)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			return err
		}
	}
	return tx.Commit()
}
