package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"strconv"

	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"io/ioutil"
)

type orgController struct {
	models *models.OrgModel
	upload   chan int
}

var OrgCtl = &orgController{
	models: new(models.OrgModel),
	upload:make(chan int,1),
}

// swagger:operation GET /v1/auth/resource/org/page StaticFiles orgController
//
// 机构信息配置管理页面
//
// 首先系统会检查用户的连接信息,如果用户被授权访问这个页面,将会返回机构配置管理页面内容,否则返回响应的错误住状态.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (orgController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}
	rst, err := hcache.GetStaticFile("AsofdateOrgPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/resource/org/get orgController orgController
//
// 查询机构信息
//
// API将会返回指定域中的机构信息,用户在请求这个API时,需要传入domain_id这个字段值
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	if domain_id == "" {
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := hjwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 417, i18n.Get(ctx.Request,"error_query_org_info"))
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/resource/org/delete orgController orgController
//
// 删除机构信息
//
// 首先系统会校验用户的权限,如果用户拥有删除机构的权限,系统将会根据用户请求的参数,删除响应的机构信息.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: JSON
//   in: query
//   description: json format
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")
	orgList := ctx.Request.FormValue("JSON")
	var mjs []models.SysOrgInfo
	err := json.Unmarshal([]byte(orgList), &mjs)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_delete_org_info"), err)
		return
	}

	if govalidator.IsEmpty(domain_id) {
		cok, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := hjwt.ParseJwt(cok.Value)
		if err != nil {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request,domain_id,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}


	err = this.models.Delete(mjs,domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 418, err.Error())
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/resource/org/update orgController orgController
//
// 更新机构信息
//
// 系统将会更具用户传入的参数,修改指定机构信息.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}
	org_unit_id := ctx.Request.FormValue("Id")
	org_unit_desc := ctx.Request.FormValue("Org_unit_desc")
	up_org_id := ctx.Request.FormValue("Up_org_id")
	org_status_id := ctx.Request.FormValue("Status_cd")

	did, err := hrpc.CheckDomainByOrgId(org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_get_domain_by_orgid"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, did, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	// 校验输入信息
	if govalidator.IsEmpty(org_unit_desc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_id_desc_empty"))
		return
	}

	if !govalidator.IsIn(org_status_id, "0", "1") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_status_empty"))
		return
	}

	if !govalidator.IsWord(org_unit_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_id_format"))
		return
	}

	if !govalidator.IsWord(up_org_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_up_id_empty"))
		return
	}

	check, err := this.models.GetSubOrgInfo(did, org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_sub_query"))
		return
	}

	for _, val := range check {
		if val.Org_unit_id == up_org_id {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_up_id_complex"))
			return
		}
	}

	err = this.models.Update(org_unit_desc, up_org_id, org_status_id, jclaim.User_id, org_unit_id, did)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_modify"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
}


// swagger:operation POST /v1/auth/resource/org/post orgController orgController
//
// 新增机构信息
//
// 想指定域中新增机构信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: Org_unit_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// - name: Org_unit_desc
//   in: query
//   description: org desc
//   required: true
//   type: string
//   format:
// - name: Up_org_id
//   in: query
//   description: up org id
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	org_unit_id := ctx.Request.FormValue("Org_unit_id")
	org_unit_desc := ctx.Request.FormValue("Org_unit_desc")
	up_org_id := ctx.Request.FormValue("Up_org_id")
	domain_id := ctx.Request.FormValue("Domain_id")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	id := utils.JoinCode(domain_id, org_unit_id)
	org_status_id := "0"

	if !govalidator.IsWord(org_unit_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_id_format"))
		return
	}

	if govalidator.IsEmpty(org_unit_desc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_id_desc_empty"))
		return
	}

	if !govalidator.IsWord(domain_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_id_check"))
		return
	}

	if !govalidator.IsWord(up_org_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_up_id_empty"))
		return
	}

	if !govalidator.IsIn(org_status_id, "0", "1") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_status_empty"))
		return
	}

	err = this.models.Post(org_unit_id, org_unit_desc, up_org_id, org_status_id,
		domain_id, jclaim.User_id, jclaim.User_id, id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_add"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
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

// swagger:operation GET /v1/auth/relation/domain/org orgController orgController
//
// 返回某个机构的所有下级机构信息
//
// 根据客户端请求时指定的机构id,获取这个id所有的下属机构信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: org_unit_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) GetSubOrgInfo(ctx *context.Context) {

	ctx.Request.ParseForm()

	org_unit_id := ctx.Request.FormValue("org_unit_id")

	did, err := hrpc.CheckDomainByOrgId(org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_get_domain_by_orgid"))
		return
	}

	rst, err := this.models.GetSubOrgInfo(did, org_unit_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_org_sub_query"))
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/resource/org/download orgController orgController
//
// 下载机构信息
//
// 下载某个指定域的所有机构信息. 只能下载用户有权限访问的域中的机构
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController) Download(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")
	domain_id := ctx.Request.FormValue("domain_id")

	if govalidator.IsEmpty(domain_id) {
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := hjwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403,i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 417, i18n.Get(ctx.Request,"error_query_org_info"))
		return
	}

	var sheet *xlsx.Sheet
	HOME := os.Getenv("HBIGDATA_HOME")
	file, err := xlsx.OpenFile(filepath.Join(HOME, "upload","template","hauthOrgExportTemplate.xlsx"))
	if err != nil {
		file = xlsx.NewFile()
		sheet, err = file.AddSheet("机构信息")
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_sheet"))
			return
		}

		{
			row := sheet.AddRow()
			cell1 := row.AddCell()
			cell1.Value = "机构编码"
			cell2 := row.AddCell()
			cell2.Value = "机构名称"
			cell3 := row.AddCell()
			cell3.Value = "上级编码"
			cell9 := row.AddCell()
			cell9.Value = "所属域"
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

		}
	} else {
		sheet = file.Sheet["机构信息"]
		if sheet == nil {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_sheet"))
			return
		}
	}
	for _, v := range rst {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = v.Code_number
		cell1.SetStyle(sheet.Rows[1].Cells[0].GetStyle())

		cell2 := row.AddCell()
		cell2.Value = v.Org_unit_desc
		cell2.SetStyle(sheet.Rows[1].Cells[1].GetStyle())

		cell3 := row.AddCell()
		cell3.Value = utils.SplitCode(v.Up_org_id)
		cell3.SetStyle(sheet.Rows[1].Cells[2].GetStyle())

		cell9 := row.AddCell()
		cell9.Value = v.Domain_id
		cell9.SetStyle(sheet.Rows[1].Cells[3].GetStyle())

		cell4 := row.AddCell()
		cell4.Value = v.Org_status_desc
		cell4.SetStyle(sheet.Rows[1].Cells[4].GetStyle())

		cell5 := row.AddCell()
		cell5.Value = v.Create_date
		cell5.SetStyle(sheet.Rows[1].Cells[5].GetStyle())

		cell6 := row.AddCell()
		cell6.Value = v.Create_user
		cell6.SetStyle(sheet.Rows[1].Cells[6].GetStyle())

		cell7 := row.AddCell()
		cell7.Value = v.Maintance_date
		cell7.SetStyle(sheet.Rows[1].Cells[7].GetStyle())

		cell8 := row.AddCell()
		cell8.Value = v.Maintance_user
		cell8.SetStyle(sheet.Rows[1].Cells[8].GetStyle())

	}

	if len(sheet.Rows) >= 3 {
		sheet.Rows = append(sheet.Rows[0:1], sheet.Rows[2:]...)
	}

	file.Write(ctx.ResponseWriter)
}

// swagger:operation GET /v1/auth/resource/org/upload orgController orgController
//
// 上传机构信息
//
// 根据客户端导入的excel格式的数据,将机构信息写入到数据库中.
//
// 这个上传过程是:增量删除, 一旦出现重复的机构,将会中断上传过程,且数据库会立刻回滚.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this orgController)Upload(ctx *context.Context){
	if len(this.upload) != 0 {
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"error_org_upload_wait"))
		return
	}

	// 从cookies中获取用户连接信息
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// 同一个时间,只能有一个导入任务
	this.upload <- 1
	defer func() {
		<-this.upload
	}()

	ctx.Request.ParseForm()
	fd, _, err := ctx.Request.FormFile("file")
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_read_upload_file"))
		return
	}

	result,err := ioutil.ReadAll(fd)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_org_read_upload_file"))
		return
	}

	// 读取上传过来的文件信息
	// 转换成二进制数据流
	file, err := xlsx.OpenBinary(result)
	sheet, ok := file.Sheet["机构信息"]
	if !ok {
		logs.Error("没有找到'机构信息'这个sheet页")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_org_sheet"))
		return
	}
	var data []models.SysOrgInfo
	for index, val := range sheet.Rows {
		if index > 0 {
			var one models.SysOrgInfo
			one.Code_number = val.Cells[0].Value
			one.Org_unit_desc = val.Cells[1].Value
			one.Domain_id = val.Cells[3].Value
			one.Org_unit_id = utils.JoinCode(one.Domain_id,one.Code_number)
			one.Up_org_id = utils.JoinCode(one.Domain_id,val.Cells[2].Value)
			one.Create_user = jclaim.User_id

			if !hrpc.DomainAuth(ctx.Request, one.Domain_id, "w") {
				hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
				return
			}
			data = append(data,one)
		}
	}
	err = this.models.Upload(data)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,err.Error())
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter,i18n.Success(ctx.Request))
}

func init() {
	hcache.RegisterStaticFile("AsofdateOrgPage", "./views/hauth/org_page.tpl")
}
