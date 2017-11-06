package service

import (
	"sync"

	"github.com/hzwy23/hauth/utils/logs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type RegisterFunc func()

// key 应用名称
// value 注册路由方法
var regApp = make(map[string]RegisterFunc)
var regLock = new(sync.RWMutex)

func AppRegister(name string, registerFunc RegisterFunc) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, ok := regApp[name]; ok {
		panic("应用已经被注册，无法再次注册")
	} else {
		regApp[name] = registerFunc
	}
}

func Bootstrap() {
	// 开启消息，
	// 将80端口的请求，重定向到443上
	go RedictToHtpps()

	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		go WriteHandleLogs(ctx)
	}, false)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		CheckConnection(ctx.ResponseWriter, ctx.Request)
	}, false)

	// 注册路由信息
	registerRouter()

	for key, fc := range regApp {
		logs.Info("register App, name is:", key)
		fc()
	}
	// 启动beego服务
	beego.Run()
}
