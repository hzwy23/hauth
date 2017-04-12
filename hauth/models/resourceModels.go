package models

import (
	"errors"

	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"time"
)

const (
	error_resource_forbid_system_resource = "禁止删除系统内置菜单信息"
	error_resource_query                  = "查询菜单资源信息失败"
	error_resource_begin                  = "初始化事物处理失败"
	error_resource_role_relation          = "删除资源与角色关系失败"
	error_resource_theme_relation         = "删除资源与主题关系失败"
	error_resource_delete                 = "删除菜单资源信息失败"
	error_resource_commit                 = "提交事务处理失败"
	error_resource_queryType              = "获取资源类型信息失败"
	error_resource_addTheme               = "新增主题风格信息失败"
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
func (ResourceModel) Get() ([]resData, error) {
	key := hcache.GenKey("RESOURCEMODELS","ALLRES")
	if hcache.IsExist(key){
		logs.Debug("get data from cache")
		rst , _ := hcache.Get(key).([]resData)
		return rst,nil
	}
	rows, err := dbobj.Query(sys_rdbms_071)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []resData
	err = dbobj.Scan(rows, &rst)
	hcache.Put(key,rst,720*time.Minute)
	return rst, err
}

func (this ResourceModel)GetChildren(res_id string )([]resData,error){
	rst,err := this.Get()
	if err != nil {
		logs.Error(err)
		return nil,err
	}
	var ret []resData
	this.dfs(rst,res_id,&ret)
	return ret,nil
}

// 所有指定资源的详细信息
func (ResourceModel) Query(res_id string) ([]resData, error) {
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
func (ResourceModel) Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id string) error {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}

	// update sys_resource_info
	_, err = tx.Exec(sys_rdbms_072, res_id, res_name, res_attr, res_up_id, res_type)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	// update sys_theme_value
	_, err = tx.Exec(sys_rdbms_073, theme_id, res_id, res_url, res_type, res_bg_color, res_class, group_id, res_img, sort_id)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}

	// 将新增资源授权给admin
	_, err = tx.Exec(sys_rdbms_074, "vertex_root_join_sysadmin", res_id)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
	}
	return tx.Commit()
}

// 新增按钮
func (this ResourceModel) PostButton(res_id, res_name, res_attr, res_up_id, res_type string) error {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
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
func (this ResourceModel) Delete(res_id string) (string, error) {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
	var rst []resData

	all, err := this.Get()
	if err != nil {
		logs.Error(err)
		return error_resource_query, err
	}

	this.dfs(all,res_id,&rst)

	// add res_id
	for _,val:=range all {
		if val.Res_id == res_id {
			rst = append(rst,val)
			break
		}
	}

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return error_resource_begin, err
	}

	for _, val := range rst {

		if val.Sys_flag == "0" {
			tx.Rollback()
			return error_resource_forbid_system_resource, errors.New(error_resource_forbid_system_resource)
		}

		_, err = tx.Exec(sys_rdbms_075, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_resource_role_relation, err
		}

		_, err = tx.Exec(sys_rdbms_076, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_resource_theme_relation, err
		}

		_, err = tx.Exec(sys_rdbms_077, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return error_resource_delete, err
		}

	}
	return error_resource_commit, tx.Commit()
}

func (this ResourceModel) Update(res_id, res_name string) error {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
	_,err := dbobj.Exec(sys_rdbms_005, res_name, res_id)
	return err
}

// 获取子资源信息
func (this ResourceModel)dfs(all []resData,res_id string,rst *[]resData){
	for _, val:=range all {
		if val.Res_up_id == res_id {
			*rst = append(*rst,val)
			this.dfs(all,val.Res_id,rst)
		}
	}
}