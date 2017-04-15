package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
)

const (
	error_resource_query       = "查询菜单资源信息失败，请重试"
	error_resource_query_theme = "查询主题信息失败，请重试"
	error_resource_exec        = "执行SQL，提交新增用户信息失败，请联系管理员"
	error_resource_type        = "资源类型不符合要求"
	error_resource_update      = "更新菜单资源名称失败"
)

type resourceController struct {
	models *models.ResourceModel
}

var ResourceCtl = &resourceController{
	new(models.ResourceModel),
}

// 菜单资源子页面路由
// swagger:operation GET /v1/auth/resource/page StaticFiles domainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (resourceController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("AsofdateResourcePage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}

	ctx.ResponseWriter.Write(rst)
}

// swagger:operation GET /v1/auth/resource/get resourceController getdomainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (this resourceController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}
	rst, err := this.models.Get()
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_query, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// swagger:operation GET /v1/auth/resource/query resourceController getdomainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (this resourceController) Query(ctx *context.Context) {
	ctx.Request.ParseForm()
	res_id := ctx.Request.FormValue("res_id")
	rst, err := this.models.Query(res_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_query, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}


// swagger:operation POST /v1/auth/resource/post resourceController getdomainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (this resourceController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	theme_id := ctx.Request.FormValue("theme_id")
	res_type := ctx.Request.FormValue("res_type")
	res_id := ctx.Request.FormValue("res_id")
	res_name := ctx.Request.FormValue("res_name")
	res_up_id := ctx.Request.FormValue("res_up_id")
	res_url := ctx.Request.FormValue("res_url")
	res_class := ctx.Request.FormValue("res_class")
	res_img := ctx.Request.FormValue("res_img")
	res_bg_color := ctx.Request.FormValue("res_bg_color")
	group_id := ctx.Request.FormValue("group_id")
	sort_id := ctx.Request.FormValue("sort_id")
	res_attr := "1"
	if res_type == "0" || res_type == "4" {
		res_attr = "0"
	}
	if res_type == "0" {
		res_up_id = "-1"
	}

	if !govalidator.IsWord(res_id) {
		logs.Error("资源编码必须由1,30位字母或数字组成")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "资源编码必须由1,30位字母或数字组成")
		return
	}

	if govalidator.IsEmpty(res_name) {
		logs.Error("菜单名称不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单名称不能为空")
		return
	}

	if govalidator.IsEmpty(res_type) {
		logs.Error("菜单类别不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单类别不能为空")
		return
	}

	switch res_type {
	case "0":
		// 首页主菜单信息
		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单路由地址不能为空")
			return
		}

		if govalidator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单样式类型不能为空")
			return
		}

		if !govalidator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单图标不能为空")
			return
		}

		if !govalidator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单分组信息必须是数字")
			return
		}

		if !govalidator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单排序号必须是数字")
			return
		}
		err := this.models.Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_exec, err)
			return
		}
	case "1":
		// 子系统菜单信息
		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单路由地址不能为空")
			return
		}

		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		if govalidator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单样式类型不能为空")
			return
		}

		if !govalidator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单图标不能为空")
			return
		}

		if !govalidator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单分组信息必须是数字")
			return
		}

		if !govalidator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单排序号必须是数字")
			return
		}

		err := this.models.Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_exec, err)
			return
		}

	case "2":
		// 功能按钮信息
		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单路由地址不能为空")
			return
		}
		sort_id = "0"
		res_img = ""
		group_id = ""
		res_class = ""
		res_bg_color = ""
		err := this.models.Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_exec, err)
			return
		}

	case "4":
		// 虚拟节点信息
		// 功能按钮信息
		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		err := this.models.PostButton(res_id, res_name, res_attr, res_up_id, res_type)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_exec, err)
			return
		}
	default:
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, error_resource_type)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")

}


// swagger:operation DELETE /v1/auth/resource/delete resourceController getdomainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (this resourceController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	res_id := ctx.Request.FormValue("res_id")

	msg, err := this.models.Delete(res_id)

	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	}

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "remove resource successfully.")
}

// swagger:operation PUT /v1/auth/resource/update resourceController getdomainShareControll
//
// Returns all domain information
//
// get special domain share information
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
func (this resourceController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	res_id := ctx.Request.FormValue("res_id")
	res_name := ctx.Request.FormValue("res_name")

	if govalidator.IsEmpty(res_name) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "资源描述不能为空.")
		return
	}

	err := this.models.Update(res_id, res_name)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, error_resource_update, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

func init() {
	hcache.RegisterStaticFile("AsofdateResourcePage", "./views/hauth/res_info_page.tpl")
}
