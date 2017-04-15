package controllers

import (
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/hauth/hcache"
)

var homePageMenusModel = new(models.HomePageMenusModel)

// swagger:operation GET /v1/auth/main/menu HomePageMenus HomePageMenus
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
// - name: TypeId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: Id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func HomePageMenus(ctx *context.Context) {
	defer hret.HttpPanic()
	typeId := ctx.Request.FormValue("TypeId")
	Id := ctx.Request.FormValue("Id")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect())
		return
	}

	ojs, err := homePageMenusModel.Get(Id, typeId, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "please contact sysadmin.query main menu failed.")
		return
	}
	ctx.ResponseWriter.Write(ojs)
}

// swagger:operation GET /v1/auth/index/entry HomePageMenus HomePageMenus
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
// - name: TypeId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: Id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func SubSystemEntry(ctx *context.Context) {
	defer hret.HttpPanic()
	ctx.Request.ParseForm()
	id := ctx.Request.FormValue("Id")
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	url := homePageMenusModel.GetUrl(jclaim.User_id, id)
	if !hcache.FileIsExist(id) {
		hcache.RegisterStaticFile(id, url)
	}
	tpl,err := hcache.GetStaticFile(id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,404,"应用不存在,或者没有注册到平台上.")
		return
	}
	ctx.ResponseWriter.Write(tpl)
}
