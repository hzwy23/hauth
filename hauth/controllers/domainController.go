package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

type DomainController struct {
	models *models.ProjectMgr
}

var DomainCtl = &DomainController{models: &models.ProjectMgr{}}

// 获取domain_info配置页面
func (DomainController) GetDomainInfoPage(ctx *context.Context) {
	defer hret.HttpPanic()

	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	file, _ := ioutil.ReadFile("./views/hauth/domain_info.tpl")
	ctx.ResponseWriter.Write(file)
}

// 查询域信息
func (this DomainController) GetDomainInfo(ctx *context.Context) {

	ctx.Request.ParseForm()

	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	offset := ctx.Request.FormValue("offset")
	limit := ctx.Request.FormValue("limit")

	rst, total, err := this.models.GetAll(offset, limit)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 312, "查询数据库失败")
	}

	hret.WriteBootstrapTableJson(ctx.ResponseWriter, total, rst)
}

// 新增域信息
func (this DomainController) PostDomainInfo(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "您没有权限新增域信息页面，请联系系统管理员")
		return
	}

	domainId := ctx.Request.FormValue("domainId")
	domainDesc := ctx.Request.FormValue("domainDesc")
	domainStatus := ctx.Request.FormValue("domainStatus")
	//校验
	if !utils.ValidAlnumAndSymbol(domainId, 1, 30) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "域名编码格式错误,应为字母或数字组合，不为空")
		return
	}

	//
	if !utils.ValidBool(domainStatus) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "域状态不能为空")
		return
	}

	if strings.TrimSpace(domainDesc) == "" {
		logs.Error("域名信息为空")
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "所属域描述信息为空，请填写域描述信息")
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	err = this.models.Post(domainId, domainDesc, domainStatus, jclaim.User_id, jclaim.Domain_id)

	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, http.StatusExpectationFailed, "add domain info failed.", err)
		return
	}
}

// 删除域信息
func (this DomainController) DeleteDomainInfo(ctx *context.Context) {
	ctx.Request.ParseForm()
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	ijs := []byte(ctx.Request.FormValue("JSON"))
	var js []models.ProjectMgr
	err := json.Unmarshal(ijs, &js)
	if err != nil {
		logs.Error(err)
		ctx.ResponseWriter.WriteHeader(http.StatusExpectationFailed)
		ctx.ResponseWriter.Write([]byte("域编码格式错误,无法删除" + string(ijs)))
		return
	}

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 310, "No Auth")
		return
	}

	err = this.models.Delete(js, jclaim.User_id, jclaim.Domain_id)
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 311, err.Error())
	} else {
		hret.WriteHttpOkMsgs(ctx.ResponseWriter, "删除域信息成功")
	}
}

// 更新域信息
func (this DomainController) UpdateDomainInfo(ctx *context.Context) {
	ctx.Request.ParseForm()

	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	domainId := ctx.Request.FormValue("domainId")
	domainDesc := ctx.Request.FormValue("domainDesc")
	domainStatus := ctx.Request.FormValue("domainStatus")

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}

	if jclaim.User_id != "admin" && domainId != jclaim.Domain_id {
		level := models.CheckDomainRights(jclaim.User_id, domainId)
		if level != 2 {
			logs.Error(err)
			hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "您没有权限编辑这个域")
			return
		}
	}

	err = this.models.Update(domainDesc, domainStatus, jclaim.User_id, domainId)
	if err != nil {
		logs.Error(err)
		ctx.ResponseWriter.WriteHeader(http.StatusExpectationFailed)
		ctx.ResponseWriter.Write([]byte("更新域信息失败" + domainId))
		return
	}
}

func (this DomainController) GetOwner(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}
	rst, err := this.models.GetOwner(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 412, "查询数据库失败")
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 获取用户自身能够访问到的域信息
func (this DomainController) GetDomainOwner(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}
	rst, err := this.models.Get(jclaim.Domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 421, "查询数据库失败")
	}

	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 获取指定域详细信息
func (this DomainController) GetDetails(ctx *context.Context) {
	ctx.Request.ParseForm()
	var domain_id = ctx.Request.FormValue("domain_id")

	rst, err := this.models.GetRow(domain_id)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 419, "查询域详细信息失败")
		return
	}
	hret.WriteJson(ctx.ResponseWriter, rst)
}

// 获取用户自己所属域的编码
func (this DomainController) GetDomainId(ctx *context.Context) {
	ctx.Request.ParseForm()

	cookie, _ := ctx.Request.Cookie("Authorization")
	jclaim, err := hjwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "No Auth")
		return
	}
	var domain_id = jclaim.Domain_id

	hret.WriteJson(ctx.ResponseWriter, domain_id)
}
