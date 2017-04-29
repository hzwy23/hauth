package models

import (
	"errors"

	"net/url"

	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/validator"
)

type ResourceModel struct {
	Mtheme ThemeResourceModel
}

type resData struct {
	Res_id        string `json:"res_id"`
	Res_name      string `json:"res_name"`
	Res_attr      string `json:"res_attr"`
	Res_attr_desc string `json:"res_attr_desc"`
	Res_up_id     string `json:"res_up_id"`
	Res_type      string `json:"res_type"`
	Res_type_desc string `json:"res_type_desc"`
	Sys_flag      string `json:"sys_flag"`
}

// 查询所有的资源信息
func (this *ResourceModel) Get() ([]resData, error) {
	rows, err := dbobj.Query(sys_rdbms_071)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []resData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

func (this *ResourceModel) GetChildren(res_id string) ([]resData, error) {
	rst, err := this.Get()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var ret []resData
	this.dfs(rst, res_id, &ret)
	return ret, nil
}

// 所有指定资源的详细信息
func (this *ResourceModel) Query(res_id string) ([]resData, error) {
	rows, err := dbobj.Query(sys_rdbms_089, res_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []resData
	err = dbobj.Scan(rows, &rst)
	return rst, err
}

// 新增资源
func (this *ResourceModel) Post(data url.Values) (string, error) {

	theme_id := data.Get("theme_id")
	res_type := data.Get("res_type")
	res_id := data.Get("res_id")
	res_name := data.Get("res_name")
	res_up_id := data.Get("res_up_id")
	res_url := data.Get("res_url")
	res_class := data.Get("res_class")
	res_img := data.Get("res_img")
	res_bg_color := data.Get("res_bg_color")
	group_id := data.Get("group_id")
	sort_id := data.Get("sort_id")
	res_open_type := data.Get("res_open_type")

	res_attr := "1"
	if res_type == "0" || res_type == "4" {
		res_attr = "0"
	}
	if res_type == "0" {
		res_up_id = "-1"
	}

	if !validator.IsWord(res_id) {
		logs.Error("资源编码必须由1,30位字母或数字组成")
		return "error_resource_res_id", errors.New("error_resource_res_id")
	}

	if validator.IsEmpty(res_name) {
		logs.Error("菜单名称不能为空")
		return "error_resource_desc_empty", errors.New("error_resource_desc_empty")
	}

	if validator.IsEmpty(res_type) {
		logs.Error("菜单类别不能为空")
		return "error_resource_type", errors.New("error_resource_type")
	}

	switch res_type {
	case "0":
		// 首页主菜单信息
		if !validator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			return "error_resource_route_uri", errors.New("error_resource_route_uri")
		}

		if validator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			return "error_resource_class_style", errors.New("error_resource_class_style")
		}

		if !validator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			return "error_resource_icon", errors.New("error_resource_icon")
		}

		if !validator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			return "error_resource_group", errors.New("error_resource_group")
		}

		if !validator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			return "error_resource_sort", errors.New("error_resource_sort")
		}

		if validator.IsEmpty(res_open_type) {
			logs.Error("打开方式不能为空")
			return "error_resource_open_type", errors.New("error_resource_open_type")
		}

	case "1":
		// 子系统菜单信息
		if !validator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			return "error_resource_route_uri", errors.New("error_resource_route_uri")
		}

		if !validator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			return "error_resource_up_id", errors.New("error_resource_up_id")
		}

		if validator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			return "error_resource_class_style", errors.New("error_resource_class_style")
		}

		if !validator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			return "error_resource_icon", errors.New("error_resource_icon")
		}

		if !validator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			return "error_resource_group", errors.New("error_resource_group")
		}

		if !validator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			return "error_resource_sort", errors.New("error_resource_sort")
		}

		if validator.IsEmpty(res_open_type) {
			logs.Error("打开方式不能为空")
			return "error_resource_open_type", errors.New("error_resource_open_type")
		}

	case "2":
		// 功能按钮信息
		if !validator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			return "error_resource_up_id", errors.New("error_resource_up_id")
		}

		if !validator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			return "error_resource_route_uri", errors.New("error_resource_route_uri")
		}
		sort_id = "0"
		res_img = ""
		group_id = ""
		res_class = ""
		res_bg_color = ""
		res_open_type = ""
		res_open_type = ""

	case "4":
		// 虚拟节点信息
		// 功能按钮信息
		if !validator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			return "error_resource_up_id", errors.New("error_resource_up_id")
		}

		err := this.addButtonMenu(res_id, res_name, res_attr, res_up_id, res_type)
		if err != nil {
			logs.Error(err)
			return "error_resource_exec", errors.New("error_resource_exec")
		}
		return "success", nil
	default:
		return "error_resource_type", errors.New("error_resource_type")
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_sql_begin", err
	}

	// update sys_resource_info
	_, err = tx.Exec(sys_rdbms_072, res_id, res_name, res_attr, res_up_id, res_type)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return "error_resource_add", err
	}

	// update sys_theme_value
	_, err = tx.Exec(sys_rdbms_073, theme_id, res_id, res_url, res_open_type, res_bg_color, res_class, group_id, res_img, sort_id)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return "error_resource_theme_add", err
	}

	// 将新增资源授权给admin
	_, err = tx.Exec(sys_rdbms_074, "vertex_root_join_sysadmin", res_id)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return "error_resource_auth_to_admin", err
	}
	if tx.Commit() != nil {
		logs.Error(err)
		return "error_resource_commit", err
	}
	return "success", nil
}

// 新增按钮
func (this ResourceModel) addButtonMenu(res_id, res_name, res_attr, res_up_id, res_type string) error {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	_, err = tx.Exec(sys_rdbms_072, res_id, res_name, res_attr, res_up_id, res_type)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(sys_rdbms_074, "vertex_root_join_sysadmin", res_id)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
	}
	return tx.Commit()
}

// 删除指定的资源
func (this *ResourceModel) Delete(res_id string) (string, error) {
	var rst []resData

	all, err := this.Get()
	if err != nil {
		logs.Error(err)
		return "error_resource_query", err
	}

	this.dfs(all, res_id, &rst)

	// add res_id
	for _, val := range all {
		if val.Res_id == res_id {
			rst = append(rst, val)
			break
		}
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_resource_begin", err
	}

	for _, val := range rst {

		if val.Sys_flag == "0" {
			tx.Rollback()
			return "error_resource_forbid_system_resource", errors.New("error_resource_forbid_system_resource")
		}

		_, err = tx.Exec(sys_rdbms_075, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_resource_role_relation", err
		}

		_, err = tx.Exec(sys_rdbms_076, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_resource_theme_relation", err
		}

		_, err = tx.Exec(sys_rdbms_077, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_resource_delete", err
		}

	}
	return "error_resource_commit", tx.Commit()
}

func (this *ResourceModel) Update(res_id, res_name string) (string, error) {

	if validator.IsEmpty(res_name) {
		return "error_resource_desc_empty", errors.New("error_resource_desc_empty")
	}

	_, err := dbobj.Exec(sys_rdbms_005, res_name, res_id)
	if err != nil {
		logs.Error(err)
		return "error_resource_update", err
	}
	return "success", nil
}

// 获取子资源信息
func (this *ResourceModel) dfs(all []resData, res_id string, rst *[]resData) {
	for _, val := range all {
		if val.Res_up_id == res_id {
			*rst = append(*rst, val)
			if val.Res_id == val.Res_up_id {
				logs.Error("层级关系错误,不允许上级菜单域当前菜单编码一致,当前菜单编码:", val.Res_id, "上级菜单编码:", val.Res_up_id)
				return
			}
			this.dfs(all, val.Res_id, rst)
		}
	}
}
