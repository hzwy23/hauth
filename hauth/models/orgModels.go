package models

import (
	"errors"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/logs"
)

type OrgModel struct {
}

type SysOrgInfo struct {
	Org_unit_id    string `json:"org_id"`
	Org_unit_desc  string `json:"org_desc"`
	Up_org_id      string `json:"up_org_id"`
	Domain_id      string `json:"domain_id"`
	Create_date    string `json:"create_date"`
	Maintance_date string `json:"modify_date"`
	Create_user    string `json:"create_user"`
	Maintance_user string `json:"modify_user"`
	Code_number    string `json:"code_number"`
	Org_dept       string `json:"org_dept,omitempty"`
}

//获取域下边所有机构号
func (OrgModel) Get(domain_id string) ([]SysOrgInfo, error) {
	var rst []SysOrgInfo
	rows, err := dbobj.Query(sys_rdbms_041, domain_id)
	if err != nil {
		return nil, err
	}

	err = dbobj.Scan(rows, &rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (this OrgModel) Delete(mjs []SysOrgInfo, domain_id string) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_sql_begin", errors.New("error_sql_begin")
	}

	for _, val := range mjs {
		// 获取这个机构的所有下属机构信息
		sublist, err := this.GetSubOrgInfo(domain_id, val.Org_unit_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_org_sub_query", errors.New("error_org_sub_query")
		}
		for _, org := range sublist {
			_, err := tx.Exec(sys_rdbms_044, org.Org_unit_id, domain_id)
			if err != nil {
				logs.Error(err)
				tx.Rollback()
				return "error_org_delete", errors.New("error_org_delete")
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_org_delete_commit", errors.New("error_org_delete_commit")
	}
	return "success", nil
}

func (this OrgModel) Update(data url.Values, user_id string) (string, error) {
	org_unit_id := data.Get("Id")
	org_unit_desc := data.Get("Org_unit_desc")
	up_org_id := data.Get("Up_org_id")

	domain_id, err := utils.SplitDomain(org_unit_id)
	if err != nil {
		logs.Error(err)
		return "as_of_date_no_separator", errors.New("as_of_date_no_separator")
	}

	if !govalidator.IsWord(org_unit_id) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	// 校验输入信息
	if govalidator.IsEmpty(org_unit_desc) {
		return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
	}

	if !govalidator.IsWord(up_org_id) {
		return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
	}

	check, err := this.GetSubOrgInfo(domain_id, org_unit_id)
	if err != nil {
		logs.Error(err)
		return "error_org_sub_query", errors.New("error_org_sub_query")
	}

	for _, val := range check {
		if val.Org_unit_id == up_org_id {
			return "error_org_up_id_complex", errors.New("error_org_up_id_complex")
		}
	}

	_, err = dbobj.Exec(sys_rdbms_069, org_unit_desc, up_org_id, user_id, org_unit_id)
	if err != nil {
		logs.Error(err)
		return "error_org_modify", err
	}
	return "success", nil
}

func (OrgModel) Post(data url.Values, user_id string) (string, error) {

	code_number := data.Get("Org_unit_id")
	org_unit_desc := data.Get("Org_unit_desc")
	up_org_id := data.Get("Up_org_id")
	domain_id := data.Get("Domain_id")

	org_unit_id := utils.JoinCode(domain_id, code_number)

	if !govalidator.IsAlnum(code_number) {
		return "error_org_id_format", errors.New("error_org_id_format")
	}

	if govalidator.IsEmpty(org_unit_desc) {
		return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
	}

	if !govalidator.IsWord(domain_id) {
		return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
	}

	if !govalidator.IsWord(up_org_id) {
		return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
	}

	_, err := dbobj.Exec(sys_rdbms_043, code_number, org_unit_desc, up_org_id, domain_id, user_id, user_id, org_unit_id)
	if err != nil {
		logs.Error(err)
		return "error_org_add", errors.New("error_org_add")
	}
	return "success", nil
}

func (this OrgModel) GetSubOrgInfo(domain_id string, org_id string) ([]SysOrgInfo, error) {
	var rst []SysOrgInfo

	all, err := this.Get(domain_id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	// 将自身机构加入到结果中.
	for _, val := range all {
		if val.Org_unit_id == org_id {
			rst = append(rst, val)
			break
		}
	}

	this.dfs(all, org_id, &rst)

	return rst, nil
}

func (this OrgModel) dfs(node []SysOrgInfo, org_id string, rst *[]SysOrgInfo) {
	for _, val := range node {
		if val.Up_org_id == org_id {
			*rst = append(*rst, val)
			if val.Org_unit_id == val.Up_org_id {
				logs.Error("当前机构与上级机构编码一致,逻辑错误,退出递归")
				return
			}
			this.dfs(node, val.Org_unit_id, rst)
		}
	}
}

func (this OrgModel) Upload(data []SysOrgInfo) (string, error) {
	tx, err := dbobj.Begin()
	if err != nil {
		logs.Error(err)
		return "error_sql_begin", errors.New("error_sql_begin")
	}

	for _, val := range data {
		if !govalidator.IsAlnum(val.Code_number) {
			tx.Rollback()
			return "error_org_id_format", errors.New("机构编码必须由1-30位字母,数字组成")
		}

		if govalidator.IsEmpty(val.Org_unit_desc) {
			tx.Rollback()
			return "error_org_id_desc_empty", errors.New("error_org_id_desc_empty")
		}

		if govalidator.IsEmpty(val.Up_org_id) {
			tx.Rollback()
			return "error_org_up_id_empty", errors.New("error_org_up_id_empty")
		}

		if !govalidator.IsAlnum(val.Domain_id) {
			tx.Rollback()
			return "as_of_date_domain_id_check", errors.New("as_of_date_domain_id_check")
		}

		_, err = tx.Exec(sys_rdbms_043, val.Code_number, val.Org_unit_desc, val.Up_org_id, val.Domain_id, val.Create_user, val.Create_user, val.Org_unit_id)
		if err != nil {
			logs.Error(err)
			tx.Rollback()
			return "error_org_upload", errors.New("上传机构信息失败,机构号是:" + val.Code_number + ",机构名称是:" + val.Org_unit_desc)
		}
	}
	err = tx.Commit()
	if err != nil {
		logs.Error(err)
		return "error_org_submit", errors.New("error_org_submit")
	}
	return "success", nil
}
