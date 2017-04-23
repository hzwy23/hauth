package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/utils/hret"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/utils/token/hjwt"
)

type themeController struct {
	muser *models.UserThemeModel
	mres  *models.ThemeResourceModel
}

var ThemeCtl = &themeController{
	new(models.UserThemeModel),
	new(models.ThemeResourceModel),
}

// swagger:operation POST /v1/auth/theme/update themeController themeController
//
// 更新用户主题信息
//
// 更新用户主题信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: theme_id
//   in: query
//   description: theme code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this *themeController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	theme_id := ctx.Request.FormValue("theme_id")

	// get user connection info from cookes.
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// 提交更新数据库请求.
	// 更新当前连接用户的主题信息
	err = this.muser.Put(jclaim.User_id, theme_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/resource/config/theme themeController themeController
//
// 更新某个资源的主题信息
//
// 更新菜单资源的主题信息,如果这个资源没有主题信息,则新增,否则更新.
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
func (this themeController) Put(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
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
	res_open_type := ctx.Request.FormValue("res_openType")

	if govalidator.IsNull(res_sort_id) {
		res_sort_id = "0"
	}

	flag, res_type := this.mres.CheckThemeExists(theme_id, res_id)
	if govalidator.IsIn(res_type, "0", "1", "2") {
		if flag == 0 {
			// 没有这个主题的配置信息,新增主题信息
			msg, err := this.mres.Post(theme_id, res_id, res_url, res_class, res_img, res_by_color, res_group_id, res_sort_id, res_open_type)
			if err != nil {
				hret.Error(ctx.ResponseWriter, 421, msg, err)
				return
			}
			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
			return
		} else if flag > 0 {
			// 更新主题信息
			err := this.mres.Update(res_url, res_by_color, res_class, res_img, res_group_id, res_sort_id, theme_id, res_id, res_open_type)
			if err != nil {
				logs.Error(err)
				hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"), err)
				return
			}
			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
			return
		} else {
			hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_update"))
			return
		}
	} else {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_theme_virtual_forbid"))
		return
	}
}

// swagger:operation GET /v1/auth/resource/queryTheme themeController themeController
//
// 查询主题信息
//
// 查询某个菜单资源,某个主题的详细信息
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
func (this themeController) QueryTheme(ctx *context.Context) {
	ctx.Request.ParseForm()
	res_id := ctx.Request.FormValue("res_id")
	theme_id := ctx.Request.FormValue("theme_id")
	rst, err := this.mres.GetDetails(res_id, theme_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_resource_query_theme"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}
