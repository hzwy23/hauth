package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/asofdate/utils/logs"
)

type RoleAndResourceModel struct {
	ResourceModel
}


func (this RoleAndResourceModel) search(rst, all []resData) []resData {
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

func (this RoleAndResourceModel) Revoke(role_id, res_id string) error {

	var rst []resData
	var load []resData
	rst = append(rst, resData{Res_id: res_id})

	// 获取已经拥有的角色
	all, err := this.GetByRoleId(role_id)
	if err != nil {
		logs.Error(err)
		return err
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
		return err
	}
	for _, val := range load {
		_, err = tx.Exec(sys_rdbms_093, role_id, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (this RoleAndResourceModel) Auth(role_id, res_id string) error {

	var load []resData
	var rst map[string]resData = make(map[string]resData)
	var row []resData

	// 获取所有资源
	all, err := this.Get()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range all {
		if val.Res_id == res_id {
			rst[res_id] = val
			row = append(row, val)
			break
		}
	}

	// 修复差异项父节点
	tmp := this.searchParent(rst, all)
	for len(tmp) != 0 {
		for _, val := range tmp {
			rst[val.Res_id] = val
		}
		tmp = this.searchParent(rst, all)
	}
	for _, val := range rst {
		load = append(load, val)
	}

	// 获取子菜单
	//获取第一层子节点
	tmp = this.search(row, all)
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

	getted, err := this.GetByRoleId(role_id)
	if err != nil {
		logs.Error(err)
		return err
	}
	var diff map[string]resData = make(map[string]resData)

	for _, val := range load {
		diff[val.Res_id] = val
	}

	for _, val := range getted {
		if v, ok := diff[val.Res_id]; ok {
			delete(diff, v.Res_id)
		}
	}
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, val := range diff {
		_, err = tx.Exec(sys_rdbms_074, role_id, val.Res_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// 查询角色已经拥有的资源信息
func (RoleAndResourceModel) GetByRoleId(role_id string) ([]resData, error) {
	rows, err := dbobj.Query(sys_rdbms_092, role_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []resData
	for rows.Next() {
		var tmp = resData{}
		err := rows.Scan(&tmp.Res_id, &tmp.Res_name, &tmp.Res_up_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		rst = append(rst, tmp)
	}
	return rst, err
}

