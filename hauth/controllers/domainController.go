package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type domainController struct {
	models *models.ProjectMgr
}

var DomainCtl = &domainController{models: &models.ProjectMgr{}}

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

	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("DomainPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.Get(ctx.Request,"as_of_date_page_not_exist"))
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
// - name: offset
//   in: query
//   description: start row number
//   required: true
//   type: integer
//   format: int32
// - name: limit
//   in: query
//   description: maximum number of results to return
//   required: true
//   type: integer
//   format: int32
// responses:
//   '200':
//     description: success
//   '403':
//     description: Insufficient permissions
//   '421':
//     description: get domain information failed.
func (this *domainController) Get(ctx *context.Context) {

	ctx.Request.ParseForm()

	if !hrpc.BasicAuth(ctx) {
		return
	}

	offset := ctx.Request.FormValue("offset")
	limit := ctx.Request.FormValue("limit")

	rst, total, err := this.models.GetAll(offset, limit)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_query"))
		return
	}

	hret.WriteBootstrapTableJson(ctx.ResponseWriter, total, rst)
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	// Get the form data
	domainId := ctx.Request.FormValue("domainId")
	domainDesc := ctx.Request.FormValue("domainDesc")
	domainStatus := ctx.Request.FormValue("domainStatus")

	// validator domain id format
	if !govalidator.IsWord(domainId) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_id_check"))
		return
	}

	// validator domain describe format. It does not allow null values
	if govalidator.IsEmpty(domainDesc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_isempty"))
		return
	}

	// validator domain status format
	// It must be in the 0 and 1
	if !govalidator.IsIn(domainStatus, "0", "1") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_status_check"))
		return
	}

	// get user connection information from cookie
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// submit new domain info to user model
	// If success, will return nil, or not.
	err = this.models.Post(domainId, domainDesc, domainStatus, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_add_failed"), err)
		return
	}

	// return success
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	ijs := []byte(ctx.Request.FormValue("JSON"))
	var js []models.ProjectMgr
	err := json.Unmarshal(ijs, &js)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_delete"))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_disconnect"))
		return
	}

	err = this.models.Delete(js, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, err.Error())
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
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

	if !hrpc.BasicAuth(ctx) {
		return
	}

	domainId := ctx.Request.FormValue("domainId")
	domainDesc := ctx.Request.FormValue("domainDesc")
	domainStatus := ctx.Request.FormValue("domainStatus")

	// 校验域名称,不能为空
	if govalidator.IsEmpty(domainDesc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_isempty"))
		return
	}

	// 校验域状态编码,必须是0或者1
	if !govalidator.IsIn(domainStatus, "0", "1") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_status_check"))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, domainId, "w") {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	err = this.models.Update(domainDesc, domainStatus, jclaim.User_id, domainId)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_update"))
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"success"))
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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_details"))
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
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
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	hret.WriteJson(ctx.ResponseWriter, jclaim.Domain_id)
}

func init() {
	hcache.RegisterStaticFile("DomainPage", "./views/hauth/domain_info.tpl")
}
