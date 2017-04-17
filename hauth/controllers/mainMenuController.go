package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/hauth/hcache"
)

var homePageMenusModel = new(models.HomePageMenusModel)

// swagger:operation GET /v1/auth/index/entry StaticFiles SubSystemEntry
//
// According to the ID number, return subsystem information page
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: Id
//   in: query
//   description: subsystem id number.
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//   '403':
//     description: disconnect, please login.
//   '404':
//     description: page not found
func SubSystemEntry(ctx *context.Context) {
	defer hret.HttpPanic()

	ctx.Request.ParseForm()
	id := ctx.Request.FormValue("Id")

	// get user connection information from cookie.
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// get url of the id number.
	url := homePageMenusModel.GetUrl(jclaim.User_id, id)

	if !hcache.FileIsExist(id) {
		hcache.RegisterStaticFile(id, url)
	}

	tpl,err := hcache.GetStaticFile(id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,404,i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(tpl)
}

// swagger:operation GET /v1/auth/main/menu HomePageMenus HomePageMenus
//
// If the request is successful, will return the user to be able to access the menu information
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
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
//   description: The menu type, 1 means home page ,2 means subsystem page
//   required: true
//   type: string
//   format:
// - name: Id
//   in: query
//   description: This up menu id , the response will return the lower menu information of the up menu id
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//   '403':
//     description: disconnect
//   '421':
//     description: get menu information failed.
func HomePageMenus(ctx *context.Context) {
	defer hret.HttpPanic()
	typeId := ctx.Request.FormValue("TypeId")
	Id := ctx.Request.FormValue("Id")

	// get user connection information from cookie
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	ojs, err := homePageMenusModel.Get(Id, typeId, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_query_menu"))
		return
	}
	ctx.ResponseWriter.Write(ojs)
}