package controllers

import (
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

const (
	error_user_role_query   = "获取用户拥有的角色信息失败"
	error_user_role_no_user = "请选择需要查询的用户账号信息"
	error_user_role_un_auth = "查询未授权角色信息失败"
)

type userRolesController struct {
	models *models.UserRolesModel
}

var UserRolesController = &userRolesController{
	models: new(models.UserRolesModel),
}

// 通过user_id用户账号，来查询这个用户拥有的角色信息
// @(http request params)  user_id
func (this userRolesController) GetRolesByUserId(ctx *context.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	rst, err := this.models.GetRolesByUser(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_role_query, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 通过user_id账号，查询这个用户能够访问，但是又没有获取到的角色信息
// @(http request param) user_id
func (this userRolesController) GetOtherRoles(ctx *context.Context) {
	ctx.Request.ParseForm()
	user_id := ctx.Request.FormValue("user_id")

	if user_id == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_role_no_user)
		return
	}

	rst, err := this.models.GetOtherRoles(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_role_un_auth, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 给指定的用户授予角色
// @(http request param) JSON
// 这个函数接收一个指定的json字符串。
func (this userRolesController) Auth(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	ijs := ctx.Request.FormValue("JSON")
	logs.Error(ijs)

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		return
	}

	msg, err := this.models.Auth(jclaim.Domain_id, jclaim.User_id, ijs)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

// 删除用户拥有的角色信息
// @(http request param) user_id role_id
func (this userRolesController) Revoke(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	user_id := ctx.Request.FormValue("user_id")
	role_id := ctx.Request.FormValue("role_id")

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		return
	}

	msg, err := this.models.Revoke(user_id, role_id, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	} else {
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
		return
	}
}
