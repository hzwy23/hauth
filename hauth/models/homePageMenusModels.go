package models

import (
	"encoding/json"

	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/dbobj"
)

type HomePageMenusModel struct {
	mur UserRolesModel
	mut UserThemeModel
	mts ThemeResourceModel
	mrs RoleAndResourceModel
}

type homePageMenuData struct {
	Res_id        string
	Res_name      string
	Res_url       string
	Res_bg_color  string
	Res_class     string
	Res_img       string
	Group_id      string
	Res_up_id     string
	Res_open_type string
}

func (this HomePageMenusModel) Get(id, typeId, useId string) ([]byte, error) {

	// 首先获取用户主题信息
	theme, err := this.mut.Get(useId)
	if err != nil || len(theme) != 1 {
		logs.Error(err)
		return nil, err
	}

	// 获取这个主题的所有资源信息
	theme_resource, err := this.mts.Get(theme[0].Theme_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 获取这个用户的角色信息
	roles, err := this.mur.GetRolesByUser(useId)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var role_list []string
	for _, val := range roles {
		role_list = append(role_list, val.Role_id)
	}

	// 获取角色拥有的资源信息
	role_resource, err := this.mrs.Gets(role_list, id, typeId)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var mres = make(map[string]resData)
	for _, val := range role_resource {
		mres[val.Res_id] = val
	}

	// 获取角色拥有的资源信息
	var rst []homePageMenuData
	for _, t_res := range theme_resource {
		if val, ok := mres[t_res.Res_id]; ok {
			var one homePageMenuData
			one.Res_id = t_res.Res_id
			one.Res_up_id = val.Res_up_id
			one.Res_name = val.Res_name
			one.Group_id = t_res.Group_id
			one.Res_bg_color = t_res.Res_bg_color
			one.Res_class = t_res.Res_class
			one.Res_img = t_res.Res_img
			one.Res_url = t_res.Res_url
			one.Res_open_type = t_res.Res_type
			rst = append(rst, one)
		}
	}
	return json.Marshal(rst)
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
