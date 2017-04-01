// Copyright 2017 The hzwy23 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// authorization management
// Now this package did not use for the time being
// all authorization handle is relation to resources. so you can find authorization in roleUserController
package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
)

// Controller
// for authorization
// this class provide 4 Method
type AuthorityController struct {
	models *models.AuthorityModel
}

var AuthroityCtl = &AuthorityController{
	models: new(models.AuthorityModel),
}

// Get authorization page
func (this *AuthorityController) GetBatchPage(ctx *context.Context) {
	if !models.BasicAuth(ctx) {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 403, "权限不足")
		return
	}

	rst, err := hcache.GetStaticFile("AuthorityPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, "页面不存在")
		return
	}

	ctx.ResponseWriter.Write(rst)
}

func init() {
	hcache.Register("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
}
