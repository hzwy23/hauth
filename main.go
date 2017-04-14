package main

import (
	_ "github.com/hzwy23/asofdate/apps"
	"github.com/astaxie/beego"
	"github.com/hzwy23/asofdate/hauth/service"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.SetStaticPath("/swagger","swagger")
	}
	service.StartHauth()
	beego.Run()
}
