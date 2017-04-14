package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/utils/hret"
)

func IndexPage(ctx *context.Context) {
	rst, err := hcache.GetStaticFile("AsofdateIndexPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}
	ctx.ResponseWriter.Write(rst)
}

func init() {
	hcache.RegisterStaticFile("AsofdateIndexPage", "./views/login.tpl")
}
