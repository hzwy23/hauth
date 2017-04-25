// Copyright 2017 The hzwy23 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// this package provide method which handle domain share action
package controllers

import (
	"encoding/json"
	"html/template"

	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logs"
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
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	// check the domain details
	// config this domain to others
	var domain_id = ctx.Request.FormValue("domain_id")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.Error(ctx.ResponseWriter, 403, i18n.ReadDomain(ctx.Request, domain_id))
		return
	}

	// get the domain details info
	rst, err := DomainCtl.models.GetRow(domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "as_of_date_domain_get_info_failed"), err)
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
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")
	// if the request argument domain_id is empty,
	// so set domain_id yourself.
	if govalidator.IsEmpty(domain_id) {
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := jwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.Error(ctx.ResponseWriter, 403, i18n.ReadDomain(ctx.Request, domain_id))
		return
	}

	// get domain_id share info
	rst, err := this.models.Get(domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domain_get_info_failed"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
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

	if govalidator.IsEmpty(domain_id) {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domain_isempty"))
		return
	}

	rst, err := this.models.UnAuth(domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "as_of_date_check_unshare"))
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
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

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form

	domain_id := form.Get("domain_id")

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.Error(ctx.ResponseWriter, 420, i18n.Get(ctx.Request, "as_of_date_domain_permission_denied_modify"))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Post(form, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg))
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
// - name: domain_id
//   in: query
//   description: domain code number
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

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")
	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.Error(ctx.ResponseWriter, 420, i18n.Get(ctx.Request, "as_of_date_domain_permission_denied_modify"))
		return
	}

	var rst []models.DomainShareData
	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_unmarsh_json"))
		return
	}

	// delete share domain info
	msg, err := this.models.Delete(rst, domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
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

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form

	domain_id := form.Get("domain_id")
	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.Error(ctx.ResponseWriter, 420, i18n.Get(ctx.Request, "as_of_date_domain_permission_denied_modify"))
		return
	}

	// get user session from cookies
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Update(form, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
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
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	rst, err := this.models.GetList(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domain_getowner"))
		return
	}

	hret.Json(ctx.ResponseWriter, rst)
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
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	rst, err := this.models.GetOwner(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domains_of_user"))
		return
	}

	hret.Json(ctx.ResponseWriter, rst)
}
