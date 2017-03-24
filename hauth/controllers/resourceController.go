package controllers

import (
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils"
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

type ResourceController struct {
	models *models.ResourceModel
}

var ResourceCtl = &ResourceController{
	new(models.ResourceModel),
}

// 菜单资源子页面路由
func (ResourceController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	hz, _ := ioutil.ReadFile("./views/hauth/res_info_page.tpl")
	ctx.ResponseWriter.Write(hz)
}

func (this ResourceController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

func (this ResourceController) Query(ctx *context.Context) {
	ctx.Request.ParseForm()
	res_id := ctx.Request.FormValue("res_id")
	rst, err := this.models.Query(res_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_query, err)
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this ResourceController) QueryTheme(ctx *context.Context) {
	ctx.Request.ParseForm()
	res_id := ctx.Request.FormValue("res_id")
	theme_id := ctx.Request.FormValue("theme_id")
	logs.Debug(res_id, theme_id)
	rst, err := this.models.QueryTheme(res_id, theme_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_resource_query_theme, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this ResourceController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

	if !utils.ValidAlphaNumber(res_id, 1, 30) {
		logs.Error("资源编码必须由1,30位字母或数字组成")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "资源编码必须由1,30位字母或数字组成")
		return
	}

	if strings.TrimSpace(res_name) == "" {
		logs.Error("菜单名称不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单名称不能为空")
		return
	}

	if strings.TrimSpace(res_type) == "" {
		logs.Error("菜单类别不能为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单类别不能为空")
		return
	}

	switch res_type {
	case "0":
		// 首页主菜单信息
		if strings.TrimSpace(res_url) == "" {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单路由地址不能为空")
			return
		}

		if strings.TrimSpace(res_up_id) == "" {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		if strings.TrimSpace(res_class) == "" {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单样式类型不能为空")
			return
		}

		if strings.TrimSpace(res_img) == "" {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单图标不能为空")
			return
		}

		if !utils.ValidNumber(group_id, 1, 2) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单分组信息必须是数字")
			return
		}

		if !utils.ValidNumber(sort_id, 1, 2) {
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
		if strings.TrimSpace(res_url) == "" {
			logs.Error("菜单路由地址不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单路由地址不能为空")
			return
		}

		if strings.TrimSpace(res_up_id) == "" {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		if strings.TrimSpace(res_class) == "" {
			logs.Error("菜单样式类型不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单样式类型不能为空")
			return
		}

		if strings.TrimSpace(res_img) == "" {
			logs.Error("菜单图标不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单图标不能为空")
			return
		}

		if !utils.ValidNumber(group_id, 1, 2) {
			logs.Error("菜单分组信息必须是数字")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单分组信息必须是数字")
			return
		}

		if !utils.ValidNumber(sort_id, 1, 2) {
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
		if strings.TrimSpace(res_up_id) == "" {
			logs.Error("菜单上级编码不能为空")
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 333, "菜单上级编码不能为空")
			return
		}

		if strings.TrimSpace(res_url) == "" {
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
		if strings.TrimSpace(res_up_id) == "" {
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

func (this ResourceController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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

func (this ResourceController) Update(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	res_id := ctx.Request.FormValue("res_id")
	res_name := ctx.Request.FormValue("res_name")
	err := this.models.Update(res_id, res_name)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, error_resource_update, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

func (this ResourceController) ConfigTheme(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	res_id := ctx.Request.FormValue("res_id")

	theme_id := ctx.Request.FormValue("theme_id")
	res_url := ctx.Request.FormValue("res_url")
	res_class := ctx.Request.FormValue("res_class")
	res_img := ctx.Request.FormValue("res_img")
	res_by_color := ctx.Request.FormValue("res_by_color")
	res_group_id := ctx.Request.FormValue("res_group_id")
	res_sort_id := ctx.Request.FormValue("res_sort_id")

	flag := this.models.CheckThemeExists(theme_id, res_id)
	if flag <= 0 {
		// 没有这个主题的配置信息,新增主题信息
		msg, err := this.models.AddThemeInfo(theme_id, res_id, res_url, res_class, res_img, res_by_color, res_group_id, res_sort_id)
		if err != nil {
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
			return
		}
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
	} else {
		// 新增主题配置信息
		err := this.models.UpdateTheme(res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, theme_id, res_id)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "修改主题配置信息失败", err)
			return
		}
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
	}
}
