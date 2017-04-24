package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func StartHauth() {
	// 开启消息，
	// 将80端口的请求，重定向到443上
	go RedictToHtpps()

	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		go WriteHandleLogs(ctx)
	}, false)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		CheckConnection(ctx.ResponseWriter, ctx.Request)
	}, false)

	registerRouter()
}
