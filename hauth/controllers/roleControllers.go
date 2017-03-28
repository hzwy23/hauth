package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"text/template"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type RoleController struct {
	models    models.RoleModel
	resModels models.ResourceModel
	roleResModels models.RoleAndResourceModel
}

var RoleCtl = &RoleController{
	models.RoleModel{},
	models.ResourceModel{},
	models.RoleAndResourceModel{},
}

func (RoleController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	hz, _ := template.ParseFiles("./views/hauth/role_info_page.tpl")
	hz.Execute(ctx.ResponseWriter, nil)
}

func (this RoleController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

	rst, err := this.models.Get(domain_id)

	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "get role info failed.", err)
		return
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this RoleController) Post(ctx *context.Context) {

	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

	if domainid != jclaim.Domain_id && jclaim.User_id != "admin" {
		level := models.CheckDomainRights(jclaim.User_id, domainid)
		if level != 2 {
			logs.Error("没有权限在这个域中新增角色信息")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "没有权限在这个域中新增角色信息")
			return
		}
	}

	//校验
	if !utils.ValidWord(roleid, 1, 30) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "please input role id number.")
		return
	}
	//
	if !utils.ValidHanAndWord(rolename, 1, 30) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "角色名称必须是汉字,字母,或者下划线的组合,并且长度不能小于30")
		return
	}

	if strings.TrimSpace(domainid) == "" {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请选择域信息")
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

func (this RoleController) Delete(ctx *context.Context) {

	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

func (this RoleController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

	did, err := models.CheckDomainByRoleId(Role_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "查询角色信息上边")
		return
	}

	if jclaim.Domain_id != did && jclaim.User_id != "admin" {
		level := models.CheckDomainRights(jclaim.User_id, did)
		if level != 2 {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "您没有权限编辑这个域中的角色信息")
			return
		}
	}

	err = this.models.Update(Role_name, Role_status, Role_id, jclaim.User_id,did)
	if err != nil {
		logs.Error(err.Error())
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "update role info failed.", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "update role info successfully.")
}
