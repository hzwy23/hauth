package controllers

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/logs"

	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/i18n"
)

var indexModels = new(models.LoginModels)

// swagger:operation GET /HomePage StaticFiles IndexPage
//
// 返回用户登录后的主菜单页面
//
// 用户登录成功后,将会根据用户主题,返回用户的主菜单页面.
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
func HomePage(ctx *context.Context) {
	defer hret.HttpPanic(func() {
		ctx.Redirect(302, "/")
	})

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		ctx.Redirect(302, "/")
		return
	}

	url := indexModels.GetDefaultPage(jclaim.User_id)

	h, err := template.ParseFiles(url)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_get_login_page"), err)
		return
	}
	h.Execute(ctx.ResponseWriter, jclaim.User_id)
}

// swagger:operation POST /login LoginSystem LoginSystem
//
// 系统登录处理服务
//
// 客户端发起登录请求到这个API,系统对用户和密码进行校验,成功返回true,如果用户和密码对应不上,返回false
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: username
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: password
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func LoginSystem(ctx *context.Context) {

	ctx.Request.ParseForm()

	userId := ctx.Request.FormValue("username")

	userPasswd := ctx.Request.FormValue("password")

	psd, err := utils.Encrypt(userPasswd)
	if err != nil {
		logs.Error("decrypt passwd failed.", psd)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 400, i18n.Get(ctx.Request,"error_system"))
		return
	}

	domainId, err := indexModels.GetDefaultDomainId(userId)
	if err != nil {
		logs.Error(userId, " 用户没有指定的域", err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, i18n.Get(ctx.Request,"error_user_no_domain"))
		return
	}

	orgid, err := indexModels.GetDefaultOrgId(userId)
	if err != nil {
		logs.Error(userId, " 用户没有指定机构", err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 402, i18n.Get(ctx.Request,"error_user_no_org"))
		return
	}

	if ok, code, cnt, rmsg := hrpc.CheckPasswd(userId, psd); ok {
		token := hjwt.GenToken(userId, domainId, orgid, 86400)
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: 86400}
		http.SetCookie(ctx.ResponseWriter, &cookie)
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
	} else {
		emsg := hret.NewHttpErrMsg(code, i18n.Get(ctx.Request,rmsg), cnt)
		hret.WriteHttpErrMsg(ctx.ResponseWriter, emsg)
	}
}

//
// swagger:operation POST /logout LoginSystem LoginSystem
//
// 安全退出系统
//
// API处理用户退出系统请求,退出系统后,系统将修改客户端的cookie信息,使其连接过时.
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: username
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: password
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: all domain information
func LogoutSystem(ctx *context.Context) {
	cookie := http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(ctx.ResponseWriter, &cookie)
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Get(ctx.Request,"logout"))
}
