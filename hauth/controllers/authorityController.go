// Copyright 2017 The hzwy23 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// authorization management
// Now this package did not use for the time being
// all authorization handle is relation to resources. so you can find authorization in roleUserController
package controllers

import (
	"html/template"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/hauth/models"
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
func (AuthorityController) GetBatchPage(ctx *context.Context) {
	hz, _ := template.ParseFiles("./views/hauth/sys_batch_page.tpl")
	hz.Execute(ctx.ResponseWriter, nil)
}
