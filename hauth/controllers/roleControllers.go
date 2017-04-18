package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/i18n"
)

type roleController struct {
	models        models.RoleModel
}

var RoleCtl = &roleController{
	models.RoleModel{},
}

// swagger:operation GET /v1/auth/role/page StaticFiles roleController
//
// 角色管理页面
//
// 如果用户被授权访问角色管理页面,则系统返回角色管理页面内容,否则返回404错误
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
func (roleController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}
	rst, err := hcache.GetStaticFile("AsofdateRolePage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/role/get roleController roleController
//
// 查询角色信息
//
// 查询指定域中的角色信息
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
func (this roleController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	if domain_id == "" {
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := hjwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,403,i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	rst, err := this.models.Get(domain_id)

	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_role_query"), err)
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/role/post roleController roleController
//
// 新增角色信息
//
// 在某个指定的域中,新增角色信息
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
func (this roleController) Post(ctx *context.Context) {

	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	//取数据
	roleid := ctx.Request.FormValue("role_id")
	rolename := ctx.Request.FormValue("role_name")
	domainid := ctx.Request.FormValue("domain_id")
	rolestatus := ctx.Request.FormValue("role_status")
	id := domainid + "_join_" + roleid
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if !hrpc.DomainAuth(ctx.Request,domainid,"w"){
		logs.Error("没有权限在这个域中新增角色信息")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied"))
		return
	}

	//校验
	if !govalidator.IsWord(roleid) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_role_id_format"))
		return
	}
	//
	if govalidator.IsEmpty(rolename) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_role_desc_empty"))
		return
	}

	if !govalidator.IsWord(domainid) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"as_of_date_domain_id_check"))
		return
	}

	if !govalidator.IsIn(rolestatus,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_role_status"))
		return
	}

	err = this.models.Post(id, rolename, jclaim.User_id, rolestatus, domainid, roleid)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_role_add_failed"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation POST /v1/auth/role/delete roleController roleController
//
// 删除角色信息
//
// 删除某个指定域中的角色信息
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
func (this roleController) Delete(ctx *context.Context) {

	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	mjson := []byte(ctx.Request.FormValue("JSON"))
	var allrole []models.RoleInfo
	err := json.Unmarshal(mjson, &allrole)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_role_json_failed"), err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	err = this.models.Delete(allrole, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 418, i18n.Get(ctx.Request,"error_role_delete_failed"))
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter,i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/role/put roleController roleController
//
// 更新角色信息
//
// 更新某个域中的角色信息,角色编码不能更新
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
func (this roleController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	Role_id := ctx.Request.FormValue("Role_id")
	Role_name := ctx.Request.FormValue("Role_name")
	Role_status := ctx.Request.FormValue("Role_status")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	did, err := hrpc.CheckDomainByRoleId(Role_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_role_get_domain"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request,did,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"as_of_date_domain_permission_denied_modify"))
		return
	}

	if !govalidator.IsWord(Role_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_role_id_format"))
		return
	}

	if govalidator.IsEmpty(Role_name){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_role_desc_empty"))
		return
	}

	if !govalidator.IsIn(Role_status,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_role_status"))
		return
	}

	err = this.models.Update(Role_name, Role_status, Role_id, jclaim.User_id, did)
	if err != nil {
		logs.Error(err.Error())
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_role_update_failed"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

func init() {
	hcache.RegisterStaticFile("AsofdateRolePage", "./views/hauth/role_info_page.tpl")
}
