package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type themeController struct {
	models *models.ThemeModel
}

var ThemeCtl = &themeController{
	new(models.ThemeModel),
}

// 修改用户主题信息
func (this *themeController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	theme_id := ctx.Request.FormValue("theme_id")

	// get user connection info from cookes.
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, "No Auth")
		return
	}
	err = this.models.Post(jclaim.User_id, theme_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "修改主题信息失败", err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "修改主题信息成功")
}
