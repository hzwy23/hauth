package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/hzwy23/dbobj"
)

type UserThemeModel struct {
}

type userThemeData struct {
	User_id    string `json:"user_id"`
	Theme_id   string `json:"theme_id"`
	Theme_desc string `json:"theme_desc"`
}

func (this UserThemeModel) Get(user_id string) ([]userThemeData, error) {
	var rst []userThemeData
	rows, err := dbobj.Query(sys_rdbms_102, user_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this UserThemeModel) Put(user_id, theme_id string) error {
	_, err := dbobj.Exec(sys_rdbms_024, theme_id, user_id)
	return err
}
