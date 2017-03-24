package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/astaxie/beego/context"

	"github.com/hzwy23/asofdate/hauth/models"

	"fmt"
	"strings"

	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/tealeg/xlsx"
)

type OrgController struct {
	models *models.OrgModel
}

var OrgCtl = &OrgController{
	models: new(models.OrgModel),
}

func (OrgController) GetOrgPage(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}
	hz, _ := template.ParseFiles("./views/hauth/org_page.tpl")
	hz.Execute(ctx.ResponseWriter, nil)
}

func (this OrgController) GetSysOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	if domain_id == "" {
		domain_id = jclaim.Domain_id
	}

	if jclaim.User_id != "admin" && domain_id != jclaim.Domain_id {
		level := models.CheckDomainRights(jclaim.User_id, domain_id)
		if level != 1 && level != 2 {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "没有权限访问这个域中的信息。")
			return
		}
	}

	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 417, "操作数据库失败")
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this OrgController) DeleteOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	orgList := ctx.Request.FormValue("JSON")
	var mjs []models.SysOrgInfo
	err := json.Unmarshal([]byte(orgList), &mjs)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "delete org info failed.", err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	err = this.models.Delete(mjs, jclaim.Org_id, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 418, err.Error())
		return
	} else {
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, "delete org info successfully.")
		return
	}
}

func (this OrgController) UpdateOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}
	org_unit_id := ctx.Request.FormValue("Id")
	org_unit_desc := ctx.Request.FormValue("Org_unit_desc")
	up_org_id := ctx.Request.FormValue("Up_org_id")
	org_status_id := ctx.Request.FormValue("Status_cd")

	did, err := models.CheckDomainByOrgId(org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限更新这个域中的机构信息")
		return
	}
	if jclaim.User_id != "admin" && jclaim.Domain_id != did {
		level := models.CheckDomainRights(jclaim.User_id, did)
		if level != 2 {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限更新这个域中的机构信息", level)
			return
		}
	}

	check, err := this.models.GetSubOrgInfo(org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "操作数据库失败。")
		return
	}

	for _, val := range check {
		if val.Org_unit_id == up_org_id {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "上级机构号不能是自己的下属机构")
			return
		}
	}

	err = this.models.Update(org_unit_desc, up_org_id, org_status_id, jclaim.User_id, org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "modify org info failed.", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "modify org info successfully")
}

func (this OrgController) InsertOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}
	org_unit_id := ctx.Request.FormValue("Org_unit_id")
	org_unit_desc := ctx.Request.FormValue("Org_unit_desc")
	up_org_id := ctx.Request.FormValue("Up_org_id")
	domain_id := ctx.Request.FormValue("Domain_id")

	if domain_id != jclaim.Domain_id && jclaim.User_id != "admin" {
		level := models.CheckDomainRights(jclaim.User_id, domain_id)
		if level != 2 {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限在这个中新增机构信息", level)
			return
		}
	}

	id := domain_id + "_join_" + org_unit_id
	create_user := jclaim.User_id
	maintance_user := jclaim.User_id
	org_status_id := "0"

	if !utils.ValidAlphaNumber(org_unit_id, 1, 30) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "机构编码必须有1,30位字母或数字组成")
		return
	}

	if strings.TrimSpace(org_unit_desc) == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "机构名称不能为空，请输入机构名称")
		return
	}

	if strings.TrimSpace(domain_id) == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择所属域，所属域不能为空")
		return
	}

	if strings.TrimSpace(up_org_id) == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择上级机构号，上级机构号不能为空")
		return
	}

	err = this.models.Post(org_unit_id, org_unit_desc, up_org_id, org_status_id,
		domain_id, create_user, maintance_user, id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "add new org info failed.", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "add new org info successfully")
}

func (OrgController) getOrgTops(node []models.SysOrgInfo) []models.SysOrgInfo {
	var ret []models.SysOrgInfo
	for _, val := range node {
		flag := true
		for _, iv := range node {
			if val.Up_org_id == iv.Org_unit_id {
				flag = false
			}
		}
		if flag {
			ret = append(ret, val)
		}
	}
	return ret
}

func (this OrgController) orgTree(node []models.SysOrgInfo, id string, d int, result *[]models.SysOrgInfo) {
	var oneline models.SysOrgInfo
	for _, val := range node {
		if val.Up_org_id == id {
			oneline = val
			oneline.Org_dept = strconv.Itoa(d)
			*result = append(*result, oneline)
			this.orgTree(node, val.Org_unit_id, d+1, result)
		}
	}
}

func (this OrgController) GetSubOrgInfo(ctx *context.Context) {

	ctx.Request.ParseForm()

	org_unit_id := ctx.Request.FormValue("org_unit_id")

	rst, err := this.models.GetSubOrgInfo(org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "操作数据库失败")
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this OrgController) Download(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")
	domain_id := ctx.Request.FormValue("domain_id")
	fmt.Println(domain_id)
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}
	if domain_id == "" {
		domain_id = jclaim.Domain_id
	}
	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 417, "操作数据库失败")
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	{
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = "机构编码"
		cell2 := row.AddCell()
		cell2.Value = "机构名称"
		cell3 := row.AddCell()
		cell3.Value = "上级用户编码"
		cell4 := row.AddCell()
		cell4.Value = "机构状态"
		cell5 := row.AddCell()
		cell5.Value = "创建日期"
		cell6 := row.AddCell()
		cell6.Value = "创建人"
		cell7 := row.AddCell()
		cell7.Value = "维护日期"
		cell8 := row.AddCell()
		cell8.Value = "维护人"
		cell9 := row.AddCell()
		cell9.Value = "所属域"
	}

	for _, v := range rst {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = v.Code_number
		cell2 := row.AddCell()
		cell2.Value = v.Org_unit_desc
		cell3 := row.AddCell()
		uplist := strings.Split(v.Up_org_id, "_join_")
		if len(uplist) == 2 {
			cell3.Value = uplist[1]
		} else {
			cell3.Value = v.Up_org_id
		}
		cell4 := row.AddCell()
		cell4.Value = v.Org_status_desc
		cell5 := row.AddCell()
		cell5.Value = v.Create_date
		cell6 := row.AddCell()
		cell6.Value = v.Create_user
		cell7 := row.AddCell()
		cell7.Value = v.Maintance_date
		cell8 := row.AddCell()
		cell8.Value = v.Maintance_user
		cell9 := row.AddCell()
		cell9.Value = v.Domain_desc
	}
	file.Write(ctx.ResponseWriter)
}
