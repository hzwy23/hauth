package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/utils"
	"github.com/hzwy23/utils/hret"

	"strings"

	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/utils/i18n"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logs"
)

type userController struct {
	models *models.UserModel
}

var UserCtl = &userController{
	new(models.UserModel),
}

// swagger:operation GET /v1/auth/user/page StaticFiles userController
//
// 获取用户管理子页面
//
// 返回用户管理页面
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (userController) Page(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	rst, err := hcache.GetStaticFile("AsofdasteUserPage")
	if err != nil {
		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
		return
	}

	ctx.ResponseWriter.Write(rst)

}

// swagger:operation GET /v1/auth/user/get userController userController
//
// 获取指定域中用户信息
//
// 返回指定域中的用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	// if the domain_id argument is empty
	// so query default domain info
	if govalidator.IsEmpty(domain_id) {
		// get user connection info from cookes.
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := jwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "r") {
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_no_auth"))
		return
	}

	// query domain info.
	rst, err := this.models.GetDefault(domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 410, i18n.Get(ctx.Request, "error_user_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation POST /v1/auth/user/post userController userController
//
// 新增用户信息
//
// 在某个域中新增用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form
	domain_id := form.Get("domainId")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_no_auth"))
		return
	}

	msg, err := this.models.Post(form, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation POST /v1/auth/user/delete userController userController
//
// 删除用户信息
//
// 删除某个域中的用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// responses:
//   '200':
//     description: success
func (this userController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	var rst []models.UserInfo
	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_json"))
		return
	}

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	for _, val := range rst {
		domain_id, err := utils.SplitDomain(val.Org_unit_id)
		if err != nil {
			logs.Error(err)
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_query_org"))
			return
		}

		if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
			hret.Error(ctx.ResponseWriter, 403, i18n.WriteDomain(ctx.Request, val.Domain_id))
			return
		}
		if val.User_id == "admin" {
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_forbid_delete_admin"))
			return
		}
		if val.User_id == jclaim.User_id {
			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_forbid_yourself"))
			return
		}
	}

	msg, err := this.models.Delete(rst)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation GET /v1/auth/user/search userController userController
//
// 搜索用户信息
//
// 客户端发起请求时,必须带如下几个参数;
//
// org_id 机构编码
// status_id 机构状态
// domain_id 所属域
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// - name: org_id
//   in: query
//   description: org code number
//   required: true
//   type: string
//   format:
// - name: status_id
//   in: query
//   description: status code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Search(ctx *context.Context) {
	ctx.Request.ParseForm()
	var org_id = ctx.Request.FormValue("org_id")
	var status_id = ctx.Request.FormValue("status_id")
	var domain_id = ctx.Request.FormValue("domain_id")
	if strings.TrimSpace(domain_id) == "" {
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := jwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}
	logs.Debug(org_id, status_id)
	rst, err := this.models.Search(org_id, status_id, domain_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_query"), err)
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

// swagger:operation PUT /v1/auth/user/put userController userController
//
// 修改用户信息
//
// 修改用户信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) Put(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form

	domain_id, err := hrpc.GetDomainId(form.Get("userId"))
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_get_domain"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_modify_passwd"))
		return
	}

	cok, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}
	msg, err := this.models.Put(form, jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation PUT /v1/auth/user/modify/passwd userController userController
//
// 修改用户密码
//
// 修改用户密码
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) ModifyPasswd(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}
	form := ctx.Request.Form

	user_id := ctx.Request.FormValue("userid")
	did, err := hrpc.GetDomainId(user_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_passwd_modify"), err)
		return
	}

	if !hrpc.DomainAuth(ctx.Request, did, "w") {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_modify_passwd"))
		return
	}

	msg, err := this.models.ModifyPasswd(form)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))

}

// swagger:operation PUT /v1/auth/user/modify/status userController userController
//
// 修改用户锁状态
//
// 修改用户锁状态
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) ModifyStatus(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx.Request) {
		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
		return
	}

	user_id := ctx.Request.FormValue("userId")
	status_id := ctx.Request.FormValue("userStatus")

	did, err := hrpc.GetDomainId(user_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_user_modify_status"), err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if jclaim.User_id == user_id {
		hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_modify_yourself"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request, did, "w") {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 401, i18n.Get(ctx.Request, "error_user_modify_passwd"))
		return
	}

	msg, err := this.models.ModifyStatus(status_id, user_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, msg), err)
		return
	}
	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
}

// swagger:operation GET /v1/auth/user/query userController userController
//
// 查询用户自身信息
//
// 查询用户自身信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (this userController) GetUserDetails(ctx *context.Context) {
	ctx.Request.ParseForm()
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 401, i18n.Disconnect(ctx.Request))
		return
	}
	rst, err := this.models.GetOwnerDetails(jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_query"))
		return
	}
	hret.Json(ctx.ResponseWriter, rst)
}

func init() {
	hcache.RegisterStaticFile("AsofdasteUserPage", "./views/hauth/UserInfoPage.tpl")
}
