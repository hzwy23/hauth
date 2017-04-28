package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hzwy23/asofdate/apps"
	"github.com/hzwy23/asofdate/hauth/service"
)

func main() {
	service.StartHauth()
	beego.Run()
}
