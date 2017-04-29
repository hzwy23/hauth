package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/groupcache"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/logs"
)

type resourceController struct {
	models *models.ResourceModel
}

var ResourceCtl = &resourceController{
	new(models.ResourceModel),
}

// swagger:operation GET /v1/auth/resource/page StaticFiles domainShareControll
//
// 返回菜单资源管理页面
//
// 系统会对请求用户权限进行校验,校验通过,将会返回菜单管理配置页面.
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// responses:
//   '200':
//     description: all domain information
func (resourceController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := groupcache.GetStaticFile("AsofdateResourcePage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}

	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/resource/get resourceController getdomainShareControll
//
// 返回系统中所有的菜单资源信息
//
// 系统会对用户权限进行校验,校验通过,将会返回菜单资源信息
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
func (this resourceController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := this.models.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_resource_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/resource/query resourceController getdomainShareControll
//
// 查询指定菜单的详细信息
//
// 查询某个指定资源的详细信息
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: res_id
//   in: query
//   description: resource code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this resourceController) Query(ctx *context.Context) {
	ctx.Request.ParseForm()
	res_id := ctx.Request.FormValue("res_id")
	rst, err := this.models.Query(res_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_resource_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/resource/post resourceController getdomainShareControll
//
// 新增菜单信息
//
// 向系统中新增菜单资源信息
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// responses:
//   '200':
//     description: success
func (this resourceController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	form := ctx.Request.Form

	msg, err := this.models.Post(form)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))

}

// swagger:operation POST /v1/auth/resource/delete resourceController getdomainShareControll
//
// 删除菜单信息
//
// 删除系统中的菜单资源信息,系统会对用户的权限进行校验,只有校验通过,才能删除菜单资源信息.
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: res_id
//   in: query
//   description: resource code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this resourceController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	res_id := ctx.Request.FormValue("res_id")

	msg, err := this.models.Delete(res_id)

	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, msg, err)
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/resource/update resourceController getdomainShareControll
//
// 更新菜单信息
//
// API只支持修改菜单的名称
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: res_id
//   in: query
//   description: resource code number
//   required: true
//   type: string
//   format:
// - name: res_name
//   in: query
//   description: resource describe
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this resourceController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	res_id := ctx.Request.FormValue("res_id")
	res_name := ctx.Request.FormValue("res_name")

	msg, err := this.models.Update(res_id, res_name)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

func init() {
	groupcache.RegisterStaticFile("AsofdateResourcePage", "./views/hauth/res_info_page.tpl")
}
