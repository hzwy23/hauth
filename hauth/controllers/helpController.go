package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
)

type helpController struct {
}

var HelpCtl = &helpController{}

func (this helpController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	rst, err := hcache.GetStaticFile("AsofdateHelpPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}
	ctx.ResponseWriter.Write(rst)
}

func init() {
	hcache.Register("AsofdateHelpPage", "./views/help/auth_help.tpl")
}
