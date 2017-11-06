package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/hauth/core/groupcache"
	"github.com/hzwy23/hauth/core/hrpc"
	"github.com/hzwy23/hauth/core/models"
	"github.com/hzwy23/hauth/utils/hret"
	"github.com/hzwy23/hauth/utils/i18n"
	"github.com/hzwy23/hauth/utils/jwt"
	"github.com/hzwy23/hauth/utils/logs"
)

type domainController struct {
	models *models.DomainMmodel
}

var DomainCtl = &domainController{models: &models.DomainMmodel{}}

// swagger:operation GET /v1/auth/domain/page StaticFiles AuthorityController
//
// If the request is successful, will be return domain information page,
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
// responses:
//   '200':
//     description: success
//   '403':
//     description: disconnect or not access.
//   '404':
//     description: page not found
func (this *domainController) Page(ctx *context.Context) {
	defer hret.HttpPanic()

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := groupcache.GetStaticFile("DomainPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
		return
	}

	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/domain/get domainController getDomainInfo
//
// If the request is successful, will return domain information that the user be able to access.
//
// The system will check user permissions.
// So,you must first login system,and then you can send the request.
//
// You must pass in two arguments, first is offset ,second is limit.
//
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
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: get domain information failed.
func (this *domainController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()

	// 权限控制
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := this.models.Get()
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domain_query"))
		return
	}

	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/domain/post domainController postDomainInfo
//
// Add domain information
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
// - name: domainId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: domainDesc
//   in: query
//   description: domain name
//   required: true
//   type: string
//   format:
// - name: domainStatus
//   in: query
//   description: domain status, 0 is enable, 1 is disable
//   required: true
//   type: integer
//   format: int32
// responses:
//   '200':
//     description: all domain information
func (this *domainController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()

	// Check user permissions
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form

	// get user connection information from cookie
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// submit new domain info to user model
	// If success, will return nil, or not.
	msg, err := this.models.Post(form, jclaim.UserId, jclaim.DomainId)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}

	// return success
	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
}

// swagger:operation POST /v1/auth/domain/delete domainController deleteDomainInfo
//
// Delete domain information.
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
// - name: _method
//   in: query
//   description: DELETE
//   required: true
//   type: string
//   format:
// - name: JSON
//   in: query
//   description: domain info, for example is ,[{Project_id\:value}]
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this *domainController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	ijs := []byte(ctx.Request.FormValue("JSON"))
	var js []models.DomainMmodel
	err := json.Unmarshal(ijs, &js)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "as_of_date_domain_delete"))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "as_of_date_disconnect"))
		return
	}

	// 授权校验
	for _, val := range js {
		if val.Project_id == jclaim.DomainId {
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_forbid_delete_your_domain"))
			return
		}

		if !hrpc.DomainAuth(ctx.Request, val.Project_id, "w") {
			hret.Error(ctx.ResponseWriter, 403, i18n.WriteDomain(ctx.Request, val.Project_id))
			return
		}
	}

	err = this.models.Delete(js)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 421, err.Error())
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
}

// swagger:operation PUT /v1/auth/domain/update domainController putDomainInfo
//
// update domain describe and domain status
//
// update domain info , you neet input three arguments, domainId,domainDesc,domainStatus. column domain_id can't update.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domainId
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: domainDesc
//   in: query
//   description: domain name
//   required: true
//   type: string
//   format:
// - name: domainStatus
//   in: query
//   description: domain status, 0 is enable, 1 is disable
//   required: true
//   type: integer
//   format: int32
// responses:
//   '200':
//     description: success
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: Post domain information failed.
func (this *domainController) Put(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	form := ctx.Request.Form

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, form.Get("domainId"), "w") {
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "as_of_date_domain_permission_denied_modify"))
		return
	}

	msg, err := this.models.Update(form, jclaim.UserId)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}

	hret.Success(ctx.ResponseWriter, i18n.Get(ctx.Request, "success"))
}

// swagger:operation GET /v1/auth/domain/row/details domainController getDomainDetailsInfo
//
// 返回指定域的详细信息.
//
// 根据客户端请求参数domain_id,返回这个域的详细信息.
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
//   '419':
//     description: get domain detailes failed.
func (this *domainController) GetDetails(ctx *context.Context) {
	ctx.Request.ParseForm()
	var domain_id = ctx.Request.FormValue("domain_id")

	rst, err := this.models.GetRow(domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "as_of_date_domain_details"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/domain/id domainController getDomainId
//
// 返回用户所属域ID号
//
// 查询用户登录信息, 根据用户信息,返回用户所属域.
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
func (this *domainController) GetId(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request), err)
		return
	}

	hret.Json(ctx.ResponseWriter, jclaim.DomainId)
}

func init() {
	groupcache.RegisterStaticFile("DomainPage", "./views/hauth/domain_info.tpl")
}
