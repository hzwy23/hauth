package controllers

import (


	"github.com/astaxie/beego/context"

	"github.com/hzwy23/asofdate/hauth/hcache"


	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"

	"github.com/tealeg/xlsx"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/i18n"
	"path/filepath"
	"os"
)

type handleLogsController struct {
	model models.HandleLogMode
}


var HandleLogsCtl = &handleLogsController{}

// Page return views/hauth/handle_logs_page.tpl content
// swagger:operation GET /v1/auth/HandleLogsPage StaticFiles handleLogsController
//
// Returns all domain information
//
// get special domain share information
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: all domain information
func (this *handleLogsController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("AsofdateHandleLogPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/handle/logs/download handleLogsController handleLogsController
//
// Returns all domain information
//
// get special domain share information
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
//     description: all domain information
func (this handleLogsController) Download(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect())
		return
	}
	rst,err := this.model.Download(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"查询日志信息失败.")
		return
	}

	file,err := xlsx.OpenFile(filepath.Join(os.Getenv("HBIGDATA_HOME"),"upload","template","hauthHandleLogsTemplate.xlsx"))
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,422,"创建excel失败.",err)
		return
	}
	sheet, ok := file.Sheet["handle_logs"]
	if !ok {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,422,"获取sheet页失败,没有找打sheet名称为 handle_logs.",)
		return
	}

	for _, v := range rst {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.Value = v.User_id
		cell1.SetStyle(sheet.Rows[1].Cells[0].GetStyle())

		cell2 := row.AddCell()
		cell2.Value = v.Handle_time
		cell2.SetStyle(sheet.Rows[1].Cells[1].GetStyle())


		cell3 := row.AddCell()
		cell3.Value = v.Client_ip
		cell3.SetStyle(sheet.Rows[1].Cells[2].GetStyle())

		cell4 := row.AddCell()
		cell4.Value = v.Method
		cell4.SetStyle(sheet.Rows[1].Cells[3].GetStyle())

		cell5 := row.AddCell()
		cell5.Value = v.Url
		cell5.SetStyle(sheet.Rows[1].Cells[4].GetStyle())

		cell6 := row.AddCell()
		cell6.Value = v.Status_code
		cell6.SetStyle(sheet.Rows[1].Cells[5].GetStyle())

		cell7 := row.AddCell()
		cell7.Value = v.Data
		cell7.SetStyle(sheet.Rows[1].Cells[6].GetStyle())
	}

	if len(sheet.Rows) >= 3 {
		sheet.Rows = append(sheet.Rows[0:1], sheet.Rows[2:]...)
	}

	file.Write(ctx.ResponseWriter)
}

// swagger:operation GET /v1/auth/handle/logs handleLogsController handleLogsController
//
// Returns all domain information
//
// get special domain share information
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: offset
//   in: query
//   description: domain code number
//   required: true
//   type: integer
//   format:
// - name: limit
//   in: query
//   description: domain code number
//   required: true
//   type: integer
//   format:
// responses:
//   '200':
//     description: all domain information
func (this handleLogsController) GetHandleLogs(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	offset := ctx.Request.FormValue("offset")
	limit := ctx.Request.FormValue("limit")
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	rst,total,err := this.model.Get(jclaim.Domain_id,offset,limit)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"查询日志信息失败.")
		return
	}
	hret.WriteBootstrapTableJson(ctx.ResponseWriter,total, rst)
}


// swagger:operation GET /v1/auth/handle/logs/search handleLogsController handleLogsController
//
// Returns all domain information
//
// get special domain share information
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: UserId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: StartDate
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: EndDate
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func (this handleLogsController) SerachLogs(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	userid := ctx.Request.FormValue("UserId")
	start := ctx.Request.FormValue("StartDate")
	end := ctx.Request.FormValue("EndDate")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, i18n.Disconnect())
		return
	}

	rst,err := this.model.Search(jclaim.Domain_id,userid,start,end)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"搜索日志失败.")
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func init() {
	hcache.RegisterStaticFile("AsofdateHandleLogPage", "./views/hauth/handle_logs_page.tpl")
}
