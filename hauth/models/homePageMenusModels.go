package models

import (
	"encoding/json"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"

)

type HomePageMenusModel struct {
	Res_id       string
	Res_name     string
	Res_url      string
	Res_bg_color string
	Res_class    string
	Res_img      string
	Group_id     string
	Res_up_id    string
}

func (this HomePageMenusModel) Get(id, typeId, useId string) ([]byte, error) {

	var rst []HomePageMenusModel
	rows, err := dbobj.Query(sys_rdbms_012, id, typeId, useId)
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
	return json.Marshal(rst)

}


func (this HomePageMenusModel)dfs(node []HomePageMenusModel,up_id string,rst *[]HomePageMenusModel){
	for _,val:=range node{
		if val.Res_up_id == up_id {
			*rst = append(*rst,val)
			this.dfs(node,val.Res_id,rst)
		}
	}
}

func (this HomePageMenusModel) GetUrl(user_id, id string) string {
	row := dbobj.QueryRow(sys_rdbms_011, user_id, id)
	var url string
	err := row.Scan(&url)
	if err != nil {
		logs.Error("cant not fetch menu_url", err)
		url = "./views/hauth/theme/default/sysconfig.tpl"
	}
	return url
}
