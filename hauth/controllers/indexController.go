package controllers

import (
	"html/template"

	"github.com/astaxie/beego/context"
)

func IndexPage(ctx *context.Context) {
	huang, _ := template.ParseFiles("./views/login.tpl")
	huang.Execute(ctx.ResponseWriter, nil)
}
