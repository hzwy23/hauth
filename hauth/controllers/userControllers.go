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
)

const (
	error_user_no_auth        = "没有权限访问这个域中的信息"
	error_user_query          = "查询域中用户信息失败，请选择域后重新查询"
	error_user_read_only      = "你对这个域拥有只读权限，无法在这个域中新增用户信息"
	error_user_id_check       = "用户账号，必须有1-30位字母，数字组成"
	error_user_name_check     = "用户名必须由1-30位汉字，字母，数字组成"
	error_user_passwd_check   = "密码长度不能为空格，切长度不能小于6位"
	error_user_passwd_encrypt = "对密码进行加密失败"
	error_user_email_check    = "邮箱格式不正确"
	error_user_phone_check    = "手机号格式不正确"
	error_user_post           = "新增用户信息失败"
	error_user_modify_passwd  = "您没有权限修改这个域中的用户信息"
)

type userController struct {
	models *models.UserModel
}

var UserCtl = &userController{
	new(models.UserModel),
}

// 获取用户管理子页面
func (userController) Page(ctx *context.Context) {
	defer hret.HttpPanic()

	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	rst, err := hcache.GetStaticFile("AsofdasteUserPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}

	ctx.ResponseWriter.Write(rst)

}

// 获取指定域中用户信息
// @(http request param) domain_id
func (this userController) Get(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	domain_id := ctx.Request.FormValue("domain_id")

	// get user connection info from cookes.
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, "No Auth")
		return
	}

	// if the domain_id argument is empty
	// so query default domain info
	if strings.TrimSpace(domain_id) == "" {
		domain_id = jclaim.Domain_id
	}

	if !models.CheckDomain(ctx,domain_id,"r"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, error_user_no_auth)
		return
	}

	// query domain info.
	rst, err := this.models.GetDefault(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 410, error_user_query, err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func (this userController) Post(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	userId := ctx.Request.FormValue("userId")
	userDesc := ctx.Request.FormValue("userDesc")
	domain_id := ctx.Request.FormValue("domainId")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, "No Auth")
		return
	}

	if !models.CheckDomain(ctx,domain_id,"w"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, error_user_no_auth)
		return
	}


	if !govalidator.IsWord(userId) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_id_check)
		return
	}
	//

	if govalidator.IsEmpty(userDesc) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_name_check)
		return
	}
	//
	password := ctx.Request.FormValue("userPasswd")
	if govalidator.IsEmpty(password) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_passwd_check)
		return
	}

	surepassword := ctx.Request.FormValue("userPasswdConfirm")
	if govalidator.IsEmpty(surepassword) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "请确认密码.")
		return
	}

	if password != surepassword{
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,419,"两次输入密码不一致.")
		return
	}

	if len(strings.TrimSpace(password)) < 6 {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"密码长度不能小于6位.")
		return
	}

	userPasswd, err := utils.Encrypt(ctx.Request.FormValue("userPasswd"))
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_passwd_encrypt, err)
		return
	}

	userStatus := ctx.Request.FormValue("userStatus")
	userEmail := ctx.Request.FormValue("userEmail")
	userPhone := ctx.Request.FormValue("userPhone")
	userOrgUnitId := ctx.Request.FormValue("userOrgUnitId")

	//
	if !govalidator.IsEmail(userEmail) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_email_check)
		return
	}

	if !govalidator.IsWord(userOrgUnitId){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择机构信息.")
		return
	}

	//
	if !govalidator.IsNumeric(userPhone) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_phone_check)
		return
	}

	err = this.models.Post(userId, userPasswd, userDesc, userStatus, jclaim.User_id, userEmail, userPhone, userOrgUnitId, domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, error_user_post, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

func (this userController) Delete(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	ijs := []byte(ctx.Request.FormValue("JSON"))

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	msg, err := this.models.Delete(ijs, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

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
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, "No Auth")
			return
		}
		domain_id = jclaim.Domain_id
	}
	logs.Debug(org_id, status_id)
	rst, err := this.models.Search(org_id, status_id, domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询用户信息失败", err)
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 修改用户信息
func (this userController) Put(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
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
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	did, err := models.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	if !models.CheckDomain(ctx,did,"w"){
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, error_user_modify_passwd)
		return
	}

	if !govalidator.IsWord(user_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"用户账号不正确.")
		return
	}

	if govalidator.IsEmpty(user_name){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"用户名不能为空")
		return
	}

	if !govalidator.IsEmail(email){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"邮箱格式不正确.")
		return
	}

	if !govalidator.IsWord(org_id){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择机构号")
		return
	}

	if !govalidator.IsNumeric(phone){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请填写手机号")
		return
	}

	msg, err := this.models.Put(user_name, org_id, phone, email, jclaim.User_id, user_id, did)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

// 导出用户信息
func (this userController) Download(ctx *context.Context) {

}

// 修改用户密码
func (this userController) ModifyPasswd(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	user_id := ctx.Request.FormValue("userid")

	user_password := ctx.Request.FormValue("newpasswd")
	confirm_password := ctx.Request.FormValue("surepasswd")
	if user_password != confirm_password {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "两次输入密码不一致，请重新输入")
		return
	}

	if len(strings.TrimSpace(confirm_password)) < 6 || len(strings.TrimSpace(confirm_password)) > 30 {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "密码长度必须是大于6位，且小于30位")
		return
	}

	did, err := models.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "修改用户密码失败", err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	if did != jclaim.Domain_id && "admin" != jclaim.User_id {
		level := models.CheckDomainRights(jclaim.User_id, did)
		if level != 2 {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, error_user_modify_passwd)
			return
		}
	}

	encry_passwd, err := utils.Encrypt(user_password)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "用户密码加密失败")
		return
	}

	msg, err := this.models.ModifyPasswd(encry_passwd, user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")

}

// 修改用户锁状态
func (this userController) ModifyStatus(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	user_id := ctx.Request.FormValue("userId")
	status_id := ctx.Request.FormValue("userStatus")

	did, err := models.CheckDomainByUserId(user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "修改用户锁状态失败", err)
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	if jclaim.User_id == user_id {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "禁止修改自身的用户状态")
		return
	}

	if !models.CheckDomain(ctx,did,"w"){
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, error_user_modify_passwd)
		return
	}

	if !govalidator.IsIn(status_id,"0","1"){
		hret.WriteHttpErrMsgs(ctx.ResponseWriter,421,"请选择用户状态.")
		return
	}

	msg, err := this.models.ModifyStatus(status_id, user_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, msg, err)
		return
	}
	hret.WriteHttpOkMsgs(ctx.ResponseWriter, "success")
}

// 查询用户自身信息
func (this userController) GetUserDetails(ctx *context.Context) {
	ctx.Request.ParseForm()
	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 401, "No Auth")
		return
	}
	rst, err := this.models.GetOwnerDetails(jclaim.User_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询用户信息失败")
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

func init() {
	hcache.Register("AsofdasteUserPage", "./views/hauth/UserInfoPage.tpl")
}
