package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/asofdate/utils/logs"
)

type RoleAndResourceModel struct {
	mres ResourceModel
}

type RoleResourceRelData struct {
	Role_id string    `json:"role_id"`
	Res_id  string	  `json:"res_id"`
}

func (this RoleAndResourceModel) Delete(role_id, res_id string) error {

	var rst []resData
	var load []resData
	rst = append(rst, resData{Res_id: res_id})

	// 获取已经拥有的角色
	all, err := this.Get(role_id)
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

func (this RoleAndResourceModel) Post(role_id, res_id string) error {

	var load []resData
	var rst map[string]resData = make(map[string]resData)
	var row []resData

	// 获取所有资源
	all, err := this.mres.Get()
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

	getted, err := this.Get(role_id)
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


// 查询没有获取到的资源信息
func (this RoleAndResourceModel) UnGetted(role_id string) ([]resData, error) {

	var rst []resData

	// 获取角色已经拥有了的资源id
	role_res,err := this.get(role_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	// 获取所有的用户信息
	rst_res,err := this.mres.Get()
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	var diff = make(map[string]resData)
	for _,val := range rst_res {
		diff[val.Res_id] = val
	}

	for _, val := range role_res {
		delete(diff,val.Res_id)
	}

	// 修复差异项父节点
	tmp := this.searchParent(diff, rst_res)
	for len(tmp) != 0 {
		for _, val := range tmp {
			diff[val.Res_id] = val
		}
		tmp = this.searchParent(diff, rst_res)
	}
	for _, val := range diff {
		rst = append(rst, val)
	}
	return rst,nil
}


// 查询角色已经拥有的资源信息
func (this RoleAndResourceModel) Get(role_id string) ([]resData, error) {


	var rst []resData

	role_res,err := this.get(role_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	rst_res,err := this.mres.Get()
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	for _,val:= range role_res {
		for _,res := range rst_res {
			if val.Res_id == res.Res_id {
				var one resData
				one.Res_id = res.Res_id
				one.Res_name = res.Res_name
				one.Res_up_id = res.Res_up_id
				rst = append(rst,one)
				break
			}
		}
	}
	return rst,nil
}

// 获取某些角色,指定资源的所有下级资源
func (this RoleAndResourceModel) Gets(roles []string,res_id ...string)([]resData,error){

	var rst []resData
	var role_res map[string]string = make(map[string]string)
	for _,val := range roles {
		tmp,err := this.get(val)
		if err != nil {
			logs.Error(err)
			return nil,err
		}
		for _, p :=range tmp {
			role_res[p.Res_id] = ""
		}
	}

	var rst_res []resData
	if len(res_id) == 1{
		var err error
		rst_res,err = this.mres.GetChildren(res_id[0])
		if err != nil {
			logs.Error(err)
			return nil,err
		}
	} else if len(res_id) == 2 {
		tmp,err := this.mres.GetChildren(res_id[0])
		if err != nil {
			logs.Error(err)
			return nil,err
		}
		for _, val := range tmp {
			if val.Res_type == res_id[1] {
				rst_res = append(rst_res,val)
			}
		}
	} else {
		var err error
		rst_res,err = this.mres.Get()
		if err != nil {
			logs.Error(err)
			return nil,err
		}
	}


	for _,res := range rst_res {
		if _,ok:=role_res[res.Res_id]; ok{
			var one resData
			one.Res_id = res.Res_id
			one.Res_name = res.Res_name
			one.Res_up_id = res.Res_up_id
			rst = append(rst,one)
		}
	}

	return rst,nil
}

// 查找所有的父级资源信息
func (this RoleAndResourceModel) searchParent(diff map[string]resData, all []resData) []resData {
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

// 获取指定角色拥有的资源ID列表
func (this RoleAndResourceModel) get(role_id string )([]RoleResourceRelData,error){

	var rst []RoleResourceRelData
	rows,err := dbobj.Query(sys_rdbms_100,role_id)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	err = dbobj.Scan(rows,&rst)
	if err != nil {
		logs.Error(err)
		return nil,err
	}

	return rst,nil
}
