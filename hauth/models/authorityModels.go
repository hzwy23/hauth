package models

import (
	"encoding/json"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
)

type AuthorityModel struct {
}

var AuthModel = new(AuthorityModel)

type mbatchUsers struct {
	User_id string `json:"user_id"`
	Role_id string `json:"role_id"`
}

type gettedRoles struct {
	Role_id     string `json:"role_id"`
	Role_name   string `json:"role_name"`
	Code_number string `json:"code_number"`
}

func (AuthorityModel) Grants(users, roles, user_id string) ([]mbatchUsers, error) {
	var user []mbatchUsers
	var role []mbatchUsers
	var ret []mbatchUsers
	err := json.Unmarshal([]byte(users), &user)
	err = json.Unmarshal([]byte(roles), &role)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(roles), &role)
	if err != nil {
		return nil, err
	}
	var tmp mbatchUsers
	tx, _ := dbobj.Begin()
	for _, u := range user {
		for _, r := range role {
			_, err := tx.Exec(sys_rdbms_048, u.User_id+"-"+r.Role_id, r.Role_id, u.User_id, user_id)
			if err != nil {
				tmp.Role_id = r.Role_id
				tmp.User_id = u.User_id
				ret = append(ret, tmp)
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (AuthorityModel) GetOwnerRoles(user_id string) ([]gettedRoles, error) {
	var rst []gettedRoles
	rows, err := dbobj.Query(sys_rdbms_046, user_id, user_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return rst, nil
}

func (AuthorityModel) GetGrantRoles(user_id, uid string) ([]gettedRoles, error) {
	var rst []gettedRoles

	rows, err := dbobj.Query(sys_rdbms_047, user_id, user_id, uid)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return rst, nil
}
