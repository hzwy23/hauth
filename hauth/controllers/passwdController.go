package controllers

import (
	"net/http"

	"strings"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type passwdController struct {
	p *models.PasswdModels
}

var PasswdController = &passwdController{
	p: &models.PasswdModels{},
}

func (this passwdController) PostModifyPasswd(ctx *context.Context) {

	ctx.Request.ParseForm()

	oriPasswd := ctx.Request.FormValue("orapasswd")
	newPasswd := ctx.Request.FormValue("newpasswd")
	surePasswd := ctx.Request.FormValue("surepasswd")

	if oriPasswd == newPasswd {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "新密码与原密码一致")
		return
	}

	if newPasswd != surePasswd {
		logs.Error("new passwd confirm failed. please check your new password and confirm password")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "new passwd confirm failed. please check your new password and confirm password")
		return
	}

	if len(strings.TrimSpace(newPasswd)) != len(newPasswd) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "新密码不能以空格开头或结尾")
		return
	}

	if len(strings.TrimSpace(newPasswd)) < 6 || len(strings.TrimSpace(newPasswd)) > 30 {
		logs.Error("新密码长度不能小于6位,且不能大于30位")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "新密码长度不能小于6位,且不能大于30位")
		return
	}

	oriEn, err := utils.Encrypt(oriPasswd)
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "new passwd confirm failed. please check your new password and confirm password")
		return
	}

	newPd, err := utils.Encrypt(newPasswd)
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "new passwd confirm failed. please check your new password and confirm password")
		return
	}
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	err = this.p.UpdateMyPasswd(newPd, jclaim.User_id, oriEn)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 312, err.Error())
		return
	}
	http.SetCookie(ctx.ResponseWriter, &http.Cookie{Name: "Authorization", Value: "", Path: "/", MaxAge: -1})
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "modify user password successfully.")
}

//func (this passwdController) AdminModifyPasswd(ctx *context.Context) {
//	ctx.Request.ParseForm()
//
//	userid := ctx.Request.FormValue("userid")
//	newPasswd := ctx.Request.FormValue("newpasswd")
//	surePasswd := ctx.Request.FormValue("surepasswd")
//
//	if newPasswd != surePasswd {
//		logs.Error("new passwd confirm failed. please check your new password and confirm password")
//		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 309, "new passwd confirm failed. please check your new password and confirm password")
//		return
//	}
//
//	newPd, err := utils.Encrypt(newPasswd)
//	if err != nil {
//		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 311, "new passwd confirm failed. please check your new password and confirm password")
//		return
//	}
//
//	err = this.p.UpdateUserPasswd(newPd, userid)
//	if err != nil {
//		logs.Error(dbobj.GetErrorMsg(err))
//		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 312, "new passwd confirm failed. please check your new password and confirm password")
//		return
//	}
//	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "modify user password successfully.")
//}
