package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"net/http"
	"strconv"

	"github.com/hzwy23/asofdate/hauth/models"

	"fmt"
	"strings"

	"github.com/hzwy23/asofdate/hauth/hcache"

	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/tealeg/xlsx"
	"github.com/asaskevich/govalidator"
)

type orgController struct {
	models *models.OrgModel
}

var OrgCtl = &orgController{
	models: new(models.OrgModel),
}

func (orgController) GetOrgPage(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		return
	}
	rst, err := hcache.GetStaticFile("AsofdateOrgPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}
	ctx.ResponseWriter.Write(rst)
}

func (this orgController) GetSysOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
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

	if !models.CheckDomain(ctx,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "没有权限访问这个域中的信息。")
		return
	}

	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 417, "操作数据库失败")
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this orgController) DeleteOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
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

func (this orgController) UpdateOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
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

	if !models.CheckDomain(ctx,did,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限更新这个域中的机构信息")
		return
	}


	// 校验输入信息
	if govalidator.IsNull(org_unit_desc){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"机构描述信息不能为空.")
		return
	}

	if !govalidator.IsIn(org_status_id,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择机构状态.")
		return
	}

	if !govalidator.IsWord(org_unit_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"机构编码不正确")
		return
	}

	if !govalidator.IsWord(up_org_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择上级机构.")
		return
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

func (this orgController) InsertOrgInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
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

	if !models.CheckDomain(ctx,domain_id,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限在这个中新增机构信息")
		return
	}

	id := domain_id + "_join_" + org_unit_id
	create_user := jclaim.User_id
	maintance_user := jclaim.User_id
	org_status_id := "0"

	if !govalidator.IsWord(org_unit_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "机构编码必须有1,30位字母或数字组成")
		return
	}

	if govalidator.IsNull(org_unit_desc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "机构名称不能为空，请输入机构名称")
		return
	}

	if !govalidator.IsWord(domain_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择所属域，所属域不能为空")
		return
	}

	if !govalidator.IsWord(up_org_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择上级机构号，上级机构号不能为空")
		return
	}

	if !govalidator.IsIn(org_status_id,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"机构状态不正确.")
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

func (orgController) getOrgTops(node []models.SysOrgInfo) []models.SysOrgInfo {
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

func (this orgController) orgTree(node []models.SysOrgInfo, id string, d int, result *[]models.SysOrgInfo) {
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

func (this orgController) GetSubOrgInfo(ctx *context.Context) {

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

func (this orgController) Download(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")
	domain_id := ctx.Request.FormValue("domain_id")

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

	if !models.CheckDomain(ctx,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,403,"您没有权限导出这个域中的机构信息.")
		return
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

func init() {
	hcache.Register("AsofdateOrgPage", "./views/hauth/org_page.tpl")
}
