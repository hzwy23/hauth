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

// 查找所有的父级资源信息
func (this ResourceModel) searchParent(diff map[string]resData, all []resData) []resData {
	var ret []resData
	for _, val := range diff {
		if _, ok := diff[val.Res_up_id]; !ok {
			for _, vl := range all {
				if vl.Res_id == val.Res_up_id {
					ret = append(ret, vl)
				}
			}
		}
	}
	return ret
}

// 查询没有获取到的资源信息
func (this ResourceModel) UnGetted(role_id string) ([]resData, error) {

	// 获取已经拥有的角色信息
	rows, err := dbobj.Query(sys_rdbms_092, role_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var get = make(map[string]resData)
	for rows.Next() {
		var tmp = resData{}
		err := rows.Scan(&tmp.Res_id, &tmp.Res_name, &tmp.Res_up_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		get[tmp.Res_id] = tmp
	}

	// 获取所有的资源信息
	all, err := this.Get()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var diff = make(map[string]resData)
	for _, val := range all {
		if _, ok := get[val.Res_id]; !ok {
			diff[val.Res_id] = val
		}
	}
	// 修复差异项父节点
	tmp := this.searchParent(diff, all)
	for len(tmp) != 0 {
		for _, val := range tmp {
			diff[val.Res_id] = val
		}
		tmp = this.searchParent(diff, all)
	}
	var ret []resData
	for _, val := range diff {
		ret = append(ret, val)
	}
	return ret, nil
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

// 查询指定资源,指定主题的详细信息
func (ResourceModel) QueryTheme(res_id string, theme_id string) ([]themeData, error) {
	rows, err := dbobj.Query(sys_rdbms_070, theme_id, res_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []themeData
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
	_, err = tx.Exec(sys_rdbms_072, res_id, res_name, res_attr, res_up_id, res_type)
	if err != nil {
		logs.Error(err)
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(sys_rdbms_073, theme_id, res_id, res_url, res_type, res_bg_color, res_class, group_id, res_img, sort_id)
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

func (this ResourceModel) search(rst, all []resData) []resData {
	var tmp []resData
	for _, val := range rst {
		for _, v := range all {
			if val.Res_id == v.Res_up_id {
				tmp = append(tmp, v)
			}
		}
	}
	return tmp
}

// 删除指定的资源
func (this ResourceModel) Delete(res_id string) (string, error) {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
	var rst []resData
	var load []resData
	all, err := this.Get()
	if err != nil {
		logs.Error(err)
		return error_resource_query, err
	}

	for _, val := range all {
		if val.Res_id == res_id {
			rst = append(rst, val)
			break
		}
	}

	//获取第一层子节点
	tmp := this.search(rst, all)
	load = append(load, tmp...)
	for tmp != nil {
		tep := this.search(tmp, all)
		if tep == nil {
			break
		} else {
			load = append(load, tep...)
			tmp = tep
		}
	}
	load = append(load, rst...)

	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return error_resource_begin, err
	}

	for _, val := range load {

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
	return dbobj.Exec(sys_rdbms_005, res_name, res_id)
}

func (this ResourceModel) CheckThemeExists(theme_id string, res_id string) int {
	cnt := 0
	err := dbobj.QueryRow(sys_rdbms_006, theme_id, res_id).Scan(&cnt)
	if err != nil {
		return -1
	}
	return cnt
}

func (this ResourceModel) UpdateTheme(res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, theme_id, res_id string) error {
	defer hcache.Delete(hcache.GenKey("RESOURCEMODELS","ALLRES"))
	return dbobj.Exec(sys_rdbms_009, res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, theme_id, res_id)
}

func (this ResourceModel) AddThemeInfo(theme_id, res_id, res_url, res_class, res_img, res_by_color, res_group_id, res_sort_id string) (string, error) {
	res_type := "0"
	err := dbobj.QueryRow(sys_rdbms_013, res_id).Scan(&res_type)
	if err != nil {
		logs.Error(err)
		return error_resource_queryType, err
	}

	err = dbobj.Exec(sys_rdbms_008, theme_id, res_id, res_url, res_type, res_by_color, res_class, res_group_id, res_img, res_sort_id)

	return error_resource_addTheme, err

}
