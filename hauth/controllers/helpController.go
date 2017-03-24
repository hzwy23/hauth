package controllers

import (
	"github.com/astaxie/beego/context"
	"io/ioutil"
)

type helpController struct{

}

var HelpCtl = &helpController{}

func (this helpController)Page(ctx *context.Context){
	hz,_:=ioutil.ReadFile("./views/help/auth_help.tpl")
	ctx.ResponseWriter.Write(hz)
}