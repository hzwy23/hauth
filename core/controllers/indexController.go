package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/hauth/core/groupcache"
	"github.com/hzwy23/hauth/utils/hret"
	"github.com/hzwy23/hauth/utils/i18n"
)

// swagger:operation GET / StaticFiles IndexPage
//
// 系统首页页面
//
// API将会返回系统首页页面给客户端
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: all domain information
func IndexPage(ctx *context.Context) {
	rst, err := groupcache.GetStaticFile("AsofdateIndexPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}
	ctx.ResponseWriter.Write(rst)
}

func init() {
	groupcache.RegisterStaticFile("AsofdateIndexPage", "./views/login.tpl")
}
