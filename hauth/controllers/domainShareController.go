// Copyright 2017 The hzwy23 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// this package provide method which handle domain share action
package controllers

import (
	"html/template"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type DomainShareController struct {
	models *models.DomainShareModel
}

var DomainShareCtl = DomainShareController{
	models: new(models.DomainShareModel),
}

// swagger:operation GET /v1/auth/domain/share/page StaticFiles DomainShareController
//
// 返回共享域管理页面
//
// 用户需要首先登录系统, 系统会对用户权限进行校验,校验成功,将会返回共享域管理页面.
//
// 用户需要在request中传入domain_id, 系统将会返回这个域的详细信息
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
//   '403':
//     description: Insufficient permissions
//   '419':
//     description: get domain information failed
//   '404':
//     description: page not found
func (DomainShareController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()

	// Check the user permissions
	if !hrpc.BasicAuth(ctx) {
		return
	}

	// check the domain details
	// config this domain to others
	var domain_id = ctx.Request.FormValue("domain_id")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 420, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	// get the domain details info
	rst, err := DomainCtl.models.GetRow(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_get_info_failed"), err)
		return
	}

	hz, _ := template.ParseFiles("./views/hauth/domain_share_info.tpl")
	hz.Execute(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/domain/share/get domainShareController getdomainShareControll
//
// 返回某个域的共享对象
//
// 客户端在请求中传入域,系统将会查询这个域的共享独享列表.
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
//   '403':
//     description: Insufficient permissions
//   '419':
//     description: get domain share information failed.
func (this DomainShareController) Get(ctx *context.Context) {
	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")
	// if the request argument domain_id is empty,
	// so set domain_id yourself.
	if strings.TrimSpace(domain_id) == "" {
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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	// get domain_id share info
	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_get_info_failed"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/domain/share/unauth domainShareController getdomainShareControll
//
// 返回某个指定的域没有共享的对象
//
// 根据传入的domain_id, 查询这个域还没有共享的域信息.
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
func (this DomainShareController) UnAuth(ctx *context.Context) {
	ctx.Request.ParseForm()
	domain_id := ctx.Request.FormValue("domain_id")
	if strings.TrimSpace(domain_id) == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_isempty"))
		return
	}
	rst, err := this.models.UnAuth(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_check_unshare"))
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/domain/share/post domainShareController postomainShareControll
//
// 新增共享域信息
//
// 首先,系统会校验用户的权限信息,如果用户被授权,则进行字段校验,如果新增的域字段信息格式正确,将会写入数据库.
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
// - name: target_domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: auth_level
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this DomainShareController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")
	target_domain_id := ctx.Request.FormValue("target_domain_id")
	auth_level := ctx.Request.FormValue("auth_level")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 420, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	if !govalidator.IsWord(target_domain_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_target"))
		return
	}

	if !govalidator.IsIn(auth_level, "1", "2") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_mode"))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Disconnect(ctx.Request))
		return
	}

	err = this.models.Post(domain_id, target_domain_id, auth_level, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_share_failed"))
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
}

// swagger:operation POST /v1/auth/domain/share/delete domainShareController postomainShareControll
//
// 删除某个域指定域的共享对象
//
// 在用户请求时,需要传入用户账号
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: _method
//   in: body
//   description: DELETE
//   required: true
//   type: string
//   format:
// - name: JSON
//   in: query
//   description: json格式信息,例如[{Uuid\:value}]
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this DomainShareController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	js := ctx.Request.FormValue("JSON")
	domain_id := ctx.Request.FormValue("domain_id")

	if !govalidator.IsWord(domain_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_id_check"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 420, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	// delete share domain info
	err := this.models.Delete(js, domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_share_delete"), err)
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
}

// swagger:operation PUT /v1/auth/domain/share/put domainShareController postomainShareControll
//
// 更新指定域的共享对象
//
// 在请求更新时,需要传入域id,域的共享对象id,共享模式
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
// - name: target_domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: auth_level
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this DomainShareController) Put(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	uuid := ctx.Request.FormValue("uuid")
	level := ctx.Request.FormValue("auth_level")
	domain_id := ctx.Request.FormValue("domain_id")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 420, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	if !govalidator.IsWord(domain_id) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_target"))
		return
	}

	if !govalidator.IsIn(level, "1", "2") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_mode"))
		return
	}

	// get user session from cookies
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Disconnect(ctx.Request))
		return
	}

	err = this.models.Update(uuid, jclaim.User_id, level)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_share_update"))
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
}

// swagger:operation GET /v1/auth/domain/owner domainShareController postomainShareControll
//
// 获取用户能够访问到的域信息.
//
// 当一个域A被共享给域B,则B将可以访问到A中的信息,这个API,将会放回某个指定域所有能够被访问到的对象.
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
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: get domain that user is able to access failed.
func (this *DomainShareController) GetAccessDomain(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	rst, err := this.models.GetList(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_getowner"))
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/domain/self/owner domainShareController postomainShareControll
//
// 获取用户能够访问到的域信息.
//
// 当一个域A被共享给域B,则B将可以访问到A中的信息,这个API,将会放回某个指定域所有能够被访问到的对象.
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
func (this *DomainShareController) GetDomainOwner(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	rst, err := this.models.GetOwner(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domains_of_user"))
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}


