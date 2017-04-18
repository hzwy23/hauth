package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/utils/hret"

	"strings"

	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/asaskevich/govalidator"
	"github.com/hzwy23/asofdate/hauth/hrpc"
	"github.com/hzwy23/asofdate/utils/i18n"
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
	defer hret.HttpPanic()

	ctx.Request.ParseForm()
	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("AsofdasteUserPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	// get user connection info from cookes.
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	// if the domain_id argument is empty
	// so query default domain info
	if strings.TrimSpace(domain_id) == "" {
		domain_id = jclaim.Domain_id
	}

	if !hrpc.DomainAuth(ctx.Request,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"error_user_no_auth"))
		return
	}

	// query domain info.
	rst, err := this.models.GetDefault(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 410, i18n.Get(ctx.Request,"error_user_query"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	userId := ctx.Request.FormValue("userId")
	userDesc := ctx.Request.FormValue("userDesc")
	domain_id := ctx.Request.FormValue("domainId")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if !hrpc.DomainAuth(ctx.Request,domain_id,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_user_no_auth"))
		return
	}


	if !govalidator.IsWord(userId) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_id_check"))
		return
	}
	//

	if govalidator.IsEmpty(userDesc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_name_check"))
		return
	}
	//
	password := ctx.Request.FormValue("userPasswd")
	if govalidator.IsEmpty(password) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_passwd_check"))
		return
	}

	surepassword := ctx.Request.FormValue("userPasswdConfirm")
	if govalidator.IsEmpty(surepassword) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_passwd_empty"))
		return
	}

	if password != surepassword{
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,419,i18n.Get(ctx.Request,"error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(password)) < 6 {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_passwd_short"))
		return
	}

	userPasswd, err := utils.Encrypt(ctx.Request.FormValue("userPasswd"))
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_passwd_encrypt"), err)
		return
	}

	userStatus := ctx.Request.FormValue("userStatus")
	userEmail := ctx.Request.FormValue("userEmail")
	userPhone := ctx.Request.FormValue("userPhone")
	userOrgUnitId := ctx.Request.FormValue("userOrgUnitId")

	//
	if !govalidator.IsEmail(userEmail) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_email_check"))
		return
	}

	if !govalidator.IsWord(userOrgUnitId){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_role_org"))
		return
	}

	//
	if !govalidator.IsMobilePhone(userPhone) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_phone_check"))
		return
	}

	err = this.models.Post(userId, userPasswd, userDesc, userStatus, jclaim.User_id, userEmail, userPhone, userOrgUnitId, domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_post"), err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	ijs := []byte(ctx.Request.FormValue("JSON"))

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	msg, err := this.models.Delete(ijs, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
		jclaim, err := hjwt.ParseJwt(cookie.Value)
		if err != nil {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
			return
		}
		domain_id = jclaim.Domain_id
	}
	logs.Debug(org_id, status_id)
	rst, err := this.models.Search(org_id, status_id, domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_query"), err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	user_name := ctx.Request.FormValue("userDesc")
	org_id := ctx.Request.FormValue("orgId")
	phone := ctx.Request.FormValue("userPhone")
	email := ctx.Request.FormValue("userEmail")
	user_id := ctx.Request.FormValue("userId")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	did, err := hrpc.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"error_user_get_domain"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request,did,"w"){
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"error_user_modify_passwd"))
		return
	}

	if !govalidator.IsWord(user_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_id_empty"))
		return
	}

	if govalidator.IsEmpty(user_name){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_desc_empty"))
		return
	}

	if !govalidator.IsEmail(email){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_email_format"))
		return
	}

	if !govalidator.IsWord(org_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_org_id_format"))
		return
	}

	if !govalidator.IsMobilePhone(phone){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_phone_format"))
		return
	}

	msg, err := this.models.Put(user_name, org_id, phone, email, jclaim.User_id, user_id, did)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	user_id := ctx.Request.FormValue("userid")

	user_password := ctx.Request.FormValue("newpasswd")
	confirm_password := ctx.Request.FormValue("surepasswd")
	if user_password != confirm_password {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_passwd_confirm_failed"))
		return
	}

	if len(strings.TrimSpace(confirm_password)) < 6 || len(strings.TrimSpace(confirm_password)) > 30 {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_passwd_short"))
		return
	}

	did, err := hrpc.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_passwd_modify"), err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if did != jclaim.Domain_id && "admin" != jclaim.User_id {
		level := hrpc.CheckDomainRights(jclaim.User_id, did)
		if level != 2 {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_user_modify_passwd"))
			return
		}
	}

	encry_passwd, err := utils.Encrypt(user_password)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_password_encrpty"))
		return
	}

	msg, err := this.models.ModifyPasswd(encry_passwd, user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))

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
	if !hrpc.BasicAuth(ctx) {
		return
	}

	user_id := ctx.Request.FormValue("userId")
	status_id := ctx.Request.FormValue("userStatus")

	did, err := hrpc.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, i18n.Get(ctx.Request,"error_user_modify_status"), err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
		return
	}

	if jclaim.User_id == user_id {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, i18n.Get(ctx.Request,"error_user_modify_yourself"))
		return
	}

	if !hrpc.DomainAuth(ctx.Request,did,"w"){
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, i18n.Get(ctx.Request,"error_user_modify_passwd"))
		return
	}

	if !govalidator.IsIn(status_id,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,i18n.Get(ctx.Request,"error_user_status_empty"))
		return
	}

	msg, err := this.models.ModifyStatus(status_id, user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, i18n.Success(ctx.Request))
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
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, i18n.Disconnect(ctx.Request))
		return
	}
	rst, err := this.models.GetOwnerDetails(jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, i18n.Get(ctx.Request,"error_user_query"))
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func init() {
	hcache.RegisterStaticFile("AsofdasteUserPage", "./views/hauth/UserInfoPage.tpl")
}
