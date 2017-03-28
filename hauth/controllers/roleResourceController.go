package controllers

import (
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/logs"
	"html/template"
)

type roleAndResourceController struct {
	model *models.RoleModel
	resRoleModel *models.RoleAndResourceModel
	resModel *models.ResourceModel
}

var RoleAndResourceCtl = &roleAndResourceController{
	new(models.RoleModel),
	new(models.RoleAndResourceModel),
	new(models.ResourceModel),
}


func (this roleAndResourceController) ResourcePage(ctx *context.Context) {
	defer hret.HttpPanic()
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	var role_id = ctx.Request.FormValue("role_id")
	rst, err := this.model.GetRow(role_id)
	if err != nil || len(rst) == 0 {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询角色资源信息失败")
		return
	}
	file, _ := template.ParseFiles("./views/hauth/res_role_rel_page.tpl")

	file.Execute(ctx.ResponseWriter, rst[0])
}


// 查询指定角色的资源信息
func (this roleAndResourceController) GetResource(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	role_id := ctx.Request.FormValue("role_id")
	type_id := ctx.Request.FormValue("type_id")

	if type_id == "0" {
		// 查询角色已经获取到的资源信息
		rst, err := this.resRoleModel.GetByRoleId(role_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询角色对应的资源信息失败")
			return
		}
		hret.WriteJson(ctx.ResponseWriter, rst)
	} else if type_id == "1" {
		// 查询角色没有获取到的资源信息
		rst, err := this.resModel.UnGetted(role_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询角色对应的资源信息失败")
			return
		}
		hret.WriteJson(ctx.ResponseWriter, rst)
	}
}
// 授权与撤权操作
func (this roleAndResourceController) HandleResource(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}
	res_id := ctx.Request.FormValue("res_id")
	role_id := ctx.Request.FormValue("role_id")
	type_id := ctx.Request.FormValue("type_id")
	logs.Debug(res_id, type_id, role_id)

	// 撤销权限操作
	if type_id == "0" {
		err := this.resRoleModel.Revoke(role_id, res_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "删除角色对应的资源信息失败")
			return
		} else {
			hret.WriteHttpOkMsgs(ctx.ResponseWriter, "撤销资源权限成功")
			return
		}
	} else {
		//授权操作
		err := this.resRoleModel.Auth(role_id, res_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "删除角色对应的资源信息失败")
			return
		} else {
			hret.WriteHttpOkMsgs(ctx.ResponseWriter, "撤销资源权限成功")
			return
		}
	}
}
