package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/i18n"
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("AsofdateResourcePage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}
	rst, err := this.models.Get()
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_resource_query"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_resource_query"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_res_id"))
		return
	}

	if govalidator.IsEmpty(res_name) {
		logs.Error("菜单名称不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_desc_empty"))
		return
	}

	if govalidator.IsEmpty(res_type) {
		logs.Error("菜单类别不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_type"))
		return
	}

	switch res_type {
	case "0":
		// 首页主菜单信息
		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_route_uri"))
			return
		}

		if govalidator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_class_style"))
			return
		}

		if !govalidator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_icon"))
			return
		}

		if !govalidator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_group"))
			return
		}

		if !govalidator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_sort"))
			return
		}
		err := this.models.Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_resource_exec"), err)
			return
		}
	case "1":
		// 子系统菜单信息
		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_route_uri"))
			return
		}

		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_up_id"))
			return
		}

		if govalidator.IsEmpty(res_class) {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_class_style"))
			return
		}

		if !govalidator.IsURI(res_img) {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_icon"))
			return
		}

		if !govalidator.IsNumeric(group_id) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_group"))
			return
		}

		if !govalidator.IsNumeric(sort_id) {
			logs.Error("菜单排序号必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_sort"))
			return
		}

		err := this.models.Post(res_id, res_name, res_attr, res_up_id, res_type, theme_id, res_url, res_bg_color, res_class, group_id, res_img, sort_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_exec"), err)
			return
		}

	case "2":
		// 功能按钮信息
		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_up_id"))
			return
		}

		if !govalidator.IsURI(res_url) {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_route_uri"))
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
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_resource_exec"), err)
			return
		}

	case "4":
		// 虚拟节点信息
		// 功能按钮信息
		if !govalidator.IsWord(res_up_id) {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_up_id"))
			return
		}

		err := this.models.PostButton(res_id, res_name, res_attr, res_up_id, res_type)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_resource_exec"), err)
			return
		}
	default:
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_type"))
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")

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

	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	res_id := ctx.Request.FormValue("res_id")
	res_name := ctx.Request.FormValue("res_name")

	if govalidator.IsEmpty(res_name) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_desc_empty"))
		return
	}

	err := this.models.Update(res_id, res_name)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_resource_update"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

func init() {
	hcache.RegisterStaticFile("AsofdateResourcePage", "./views/hauth/res_info_page.tpl")
}
