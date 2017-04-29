package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logs"
)

type ThemeResourceModel struct {
}

type themeData struct {
	Theme_id     string `json:"theme_id"`
	Theme_desc   string `json:"theme_desc"`
	Res_id       string `json:"res_id"`
	Res_url      string `json:"res_url"`
	Res_type     string `json:"res_type"`
	Res_bg_color string `json:"res_bg_color"`
	Res_class    string `json:"res_class"`
	Group_id     string `json:"group_id"`
	Res_img      string `json:"res_img"`
	Sort_id      string `json:"sort_id"`
}

func (this ThemeResourceModel) Get(theme_id string) ([]themeData, error) {
	var rst []themeData
	rows, err := dbobj.Query(sys_rdbms_101, theme_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 查询指定资源,指定主题的详细信息
func (this ThemeResourceModel) GetDetails(res_id string, theme_id string) ([]themeData, error) {
	rows, err := dbobj.Query(sys_rdbms_070, theme_id, res_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []themeData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this ThemeResourceModel) Update(res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, theme_id, res_id, res_open_type string) error {
	_, err := dbobj.Exec(sys_rdbms_009, res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, res_open_type, theme_id, res_id)
	return err
}

func (this ThemeResourceModel) Post(theme_id, res_id, res_url, res_class, res_img, res_by_color, res_group_id, res_sort_id, res_type string) (string, error) {

	_, err := dbobj.Exec(sys_rdbms_008, theme_id, res_id, res_url, res_type, res_by_color, res_class, res_group_id, res_img, res_sort_id)

	return "error_resource_addTheme", err

}

func (this ThemeResourceModel) CheckThemeExists(theme_id string, res_id string) (int, string) {
	cnt := -1
	err := dbobj.QueryRow(sys_rdbms_006, theme_id, res_id).Scan(&cnt)
	if err != nil {
		return -1, ""
	}
	res_type := "4"
	err = dbobj.QueryRow(sys_rdbms_013, res_id).Scan(&res_type)
	if err != nil {
		return -1, ""
	}
	return cnt, res_type
}
