package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
)

type roleController struct {
	models        models.RoleModel
}

var RoleCtl = &roleController{
	models.RoleModel{},
}

func (roleController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}
	rst, err := hcache.GetStaticFile("AsofdateRolePage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}
	ctx.ResponseWriter.Write(rst)
}

// Page return views/hauth/sys_batch_page.tpl content
// swagger:route GET /v1/auth/resource/role/get roleController roleController
//
// Lists Page filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//     - application/x-protobuf
//
//     Produces:
//     - application/json
//     - application/octet-stream
//     - application/x-protobuf
//
//     Schemes: https
//
//     Security:
//       api_key:
//       oauth: read, write
//
//     Responses:
//       default:
//       200:
//       404:
//       422:
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
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.CheckDomain(ctx,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"您没有权限访问这个域中的角色信息.")
		return
	}

	rst, err := this.models.Get(domain_id)

	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "get role info failed.", err)
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	if !hrpc.CheckDomain(ctx,domainid,"w"){
		logs.Error("没有权限在这个域中新增角色信息")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "没有权限在这个域中新增角色信息")
		return
	}

	//校验
	if !govalidator.IsWord(roleid) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请输入由1-30位字母,数字组的角色编码")
		return
	}
	//
	if govalidator.IsEmpty(rolename) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "角色名称必须是汉字,字母,或者下划线的组合,并且长度不能小于30")
		return
	}

	if !govalidator.IsWord(domainid) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择域信息")
		return
	}

	if !govalidator.IsIn(rolestatus,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择角色状态.")
		return
	}

	err = this.models.Post(id, rolename, jclaim.User_id, rolestatus, domainid, roleid)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "add new role info failed.", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "add new role info successfully.")
}

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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "json解析失败，请重新选择需要删除的角色信息", err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	err = this.models.Delete(allrole, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 418, "操作数据库失败。")
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "删除角色信息成功。")
}

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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	did, err := hrpc.CheckDomainByRoleId(Role_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "查询角色信息上边")
		return
	}

	if !hrpc.CheckDomain(ctx,did,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限编辑这个域中的角色信息")
		return
	}

	if !govalidator.IsWord(Role_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"角色编码不正确")
		return
	}

	if govalidator.IsEmpty(Role_name){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"角色名称不正确.")
		return
	}

	if !govalidator.IsIn(Role_status,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择角色状态")
		return
	}

	err = this.models.Update(Role_name, Role_status, Role_id, jclaim.User_id, did)
	if err != nil {
		logs.Error(err.Error())
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "update role info failed.", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "update role info successfully.")
}

func init() {
	hcache.RegisterStaticFile("AsofdateRolePage", "./views/hauth/role_info_page.tpl")
}
