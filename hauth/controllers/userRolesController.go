package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/utils/i18n"
)


type userRolesController struct {
	models *models.UserRolesModel
}

var UserRolesCtl = &userRolesController{
	models: new(models.UserRolesModel),
}


// swagger:operation GET /v1/auth/batch/page StaticFiles domainShareControll
//
// If the request is successful,
// will return authorization page information to the client
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// If the user is authorized to visit, the return authorization information page
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: request success.
//   '404':
//     description: page not found.
func (this *userRolesController) Page(ctx *context.Context) {
	if !hrpc.BasicAuth(ctx) {
		return
	}

	// According to the key get the value from the groupCache system
	rst, err := hcache.GetStaticFile("AuthorityPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.Get(ctx.Request,"as_of_date_page_not_exist"))
		return
	}

	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/user/roles/get userRolesController userRolesController
//
// 通过user_id用户账号，来查询这个用户拥有的角色信息
//
// 查询角色信息
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
func (this userRolesController) GetRolesByUserId(ctx *context.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	rst, err := this.models.GetRolesByUser(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_role_query"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/user/roles/other userRolesController userRolesController
//
// 通过user_id账号，查询这个用户能够访问，但是又没有获取到的角色信息
//
// 查询用户没有获取的角色
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
func (this userRolesController) GetOtherRoles(ctx *context.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	if user_id == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_role_no_user"))
		return
	}

	rst, err := this.models.GetOtherRoles(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_role_un_auth"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/user/roles/auth userRolesController userRolesController
//
// 给指定的用户授予角色
//
// 给指定的用户授予角色
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
func (this userRolesController) Auth(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	ijs := ctx.Request.FormValue("JSON")
	logs.Error(ijs)

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,403,i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Auth(jclaim.Domain_id, jclaim.User_id, ijs)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation POST /v1/auth/user/roles/revoke userRolesController userRolesController
//
// Delete user has been granted the roles
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// If the user is authorized to visit, the system will delete the roles that client request specified.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: user_id
//   in: query
//   description: Removed the role of the user
//   required: true
//   type: string
//   format:
// - name: role_id
//   in: query
//   description: The role of ready to delete
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userRolesController) Revoke(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	user_id := ctx.Request.FormValue("user_id")
	role_id := ctx.Request.FormValue("role_id")

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,403,i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Revoke(user_id, role_id, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	} else {
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
		return
	}
}

func init() {
	// Registered in the static page to the groupCache system
	hcache.RegisterStaticFile("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
}
