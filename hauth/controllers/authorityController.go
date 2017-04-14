// Copyright 2017 The hzwy23 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// authorization management
// Now this package did not use for the time being
// all authorization handle is relation to resources.
// so you can find authorization in roleUserController
package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/hcache"
	"github.com/hzwy23/asofdate/hauth/models"
	"github.com/hzwy23/asofdate/utils/hret"
	"github.com/hzwy23/asofdate/utils/i18n"
	"github.com/hzwy23/asofdate/hauth/hrpc"
)

// Controller
// for authorization
type AuthorityController struct {
	models *models.AuthorityModel
}

var AuthroityCtl = &AuthorityController{
	models: new(models.AuthorityModel),
}

// Page return views/hauth/sys_batch_page.tpl content
// first get static file content from groupcache, if didn't find
// return 404
// swagger:route GET /v1/auth/batch/page StaticFiles AuthorityController
//
// Lists Page filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//     - text/plain
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Schemes: https
//
//     Responses:
//       200:
//       404:
func (this *AuthorityController) Page(ctx *context.Context) {
	if !hrpc.BasicAuth(ctx) {
		return
	}

	rst, err := hcache.GetStaticFile("AuthorityPage")
	if err != nil {
		hret.WriteHttpErrMsgs(ctx.ResponseWriter, 404, i18n.Get("as_of_date_page_not_exist"))
		return
	}

	ctx.ResponseWriter.Write(rst)
}

func init() {
	hcache.RegisterStaticFile("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
}
