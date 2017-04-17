package service

import (
	"github.com/astaxie/beego"
	"github.com/hzwy23/asofdate/hauth/controllers"
)

func registerRouter() {

	beego.Get("/HomePage", controllers.HomePage)

	beego.Post("/login", controllers.LoginSystem)

	beego.Any("/logout", controllers.LogoutSystem)

	beego.Post("/v1/auth/theme/update", controllers.ThemeCtl.Post)
	beego.Put("/v1/auth/resource/config/theme", controllers.ThemeCtl.Put)
	beego.Get("/v1/auth/resource/queryTheme", controllers.ThemeCtl.QueryTheme)


	beego.Get("/", controllers.IndexPage)

	beego.Get("/v1/auth/index/entry", controllers.SubSystemEntry)
	beego.Get("/v1/auth/main/menu", controllers.HomePageMenus)
	beego.Post("/v1/auth/passwd/update", controllers.PasswdController.PostModifyPasswd)

	//domain_info
	beego.Get("/v1/auth/domain/share/page", controllers.DomainShareCtl.Page)
	beego.Get("/v1/auth/domain/get", controllers.DomainCtl.Get)
	beego.Post("/v1/auth/domain/post", controllers.DomainCtl.Post)
	beego.Post("/v1/auth/domain/delete", controllers.DomainCtl.Delete)
	beego.Put("/v1/auth/domain/update", controllers.DomainCtl.Put)

	//domain_share_info
	beego.Get("/v1/auth/domain/share/get", controllers.DomainShareCtl.Get)
	beego.Post("/v1/auth/domain/share/post", controllers.DomainShareCtl.Post)
	beego.Put("/v1/auth/domain/share/put", controllers.DomainShareCtl.Put)
	beego.Post("/v1/auth/domain/share/delete", controllers.DomainShareCtl.Delete)

	beego.Get("/v1/auth/domain/owner", controllers.DomainShareCtl.GetAccessDomain)
	beego.Get("/v1/auth/domain/self/owner", controllers.DomainShareCtl.GetDomainOwner)
	beego.Get("/v1/auth/domain/row/details", controllers.DomainCtl.GetDetails)
	beego.Get("/v1/auth/domain/share/unauth", controllers.DomainShareCtl.UnAuth)

	//handle_logs
	beego.Get("/v1/auth/handle/logs/search", controllers.HandleLogsCtl.SerachLogs)
	beego.Get("/v1/auth/handle/logs", controllers.HandleLogsCtl.GetHandleLogs)
	beego.Get("/v1/auth/handle/logs/download", controllers.HandleLogsCtl.Download)

	//org_info
	beego.Get("/v1/auth/resource/org/get", controllers.OrgCtl.Get)
	beego.Post("/v1/auth/resource/org/insert", controllers.OrgCtl.Post)
	beego.Put("/v1/auth/resource/org/update", controllers.OrgCtl.Update)
	beego.Post("/v1/auth/resource/org/delete", controllers.OrgCtl.Delete)
	beego.Get("/v1/auth/resource/org/download", controllers.OrgCtl.Download)
	beego.Post("/v1/auth/resource/org/upload", controllers.OrgCtl.Upload)
	beego.Get("/v1/auth/relation/domain/org", controllers.OrgCtl.GetSubOrgInfo)
	beego.Get("/v1/auth/domain/id", controllers.DomainCtl.GetId)

	//resource_info
	beego.Post("/v1/auth/resource/delete", controllers.ResourceCtl.Delete)
	beego.Post("/v1/auth/resource/post", controllers.ResourceCtl.Post)
	beego.Put("/v1/auth/resource/update", controllers.ResourceCtl.Update)
	beego.Get("/v1/auth/resource/get", controllers.ResourceCtl.Get)
	beego.Get("/v1/auth/resource/query", controllers.ResourceCtl.Query)

	//role_resource_info
	beego.Get("/v1/auth/role/resource/get", controllers.RoleAndResourceCtl.GetResource)
	beego.Post("/v1/auth/role/resource/rights", controllers.RoleAndResourceCtl.HandleResource)

	//role_info
	beego.Get("/v1/auth/role/get", controllers.RoleCtl.Get)
	beego.Post("/v1/auth/role/post", controllers.RoleCtl.Post)
	beego.Put("/v1/auth/role/update", controllers.RoleCtl.Update)
	beego.Post("/v1/auth/role/delete", controllers.RoleCtl.Delete)

	// role and resource relation
	beego.Get("/v1/auth/role/resource/details", controllers.RoleAndResourceCtl.ResourcePage)

	//sys_batch_info
	beego.Get("/v1/auth/user/roles/get", controllers.UserRolesCtl.GetRolesByUserId)
	beego.Get("/v1/auth/user/search", controllers.UserCtl.Search)
	beego.Get("/v1/auth/user/roles/other", controllers.UserRolesCtl.GetOtherRoles)
	beego.Post("/v1/auth/user/roles/auth", controllers.UserRolesCtl.Auth)
	beego.Post("/v1/auth/user/roles/revoke", controllers.UserRolesCtl.Revoke)

	//user_info
	beego.Get("/v1/auth/user/get", controllers.UserCtl.Get)
	beego.Post("/v1/auth/user/post", controllers.UserCtl.Post)
	beego.Put("/v1/auth/user/put", controllers.UserCtl.Put)
	beego.Put("/v1/auth/user/modify/passwd", controllers.UserCtl.ModifyPasswd)
	beego.Put("/v1/auth/user/modify/status", controllers.UserCtl.ModifyStatus)
	beego.Post("/v1/auth/user/delete", controllers.UserCtl.Delete)
	beego.Get("/v1/auth/user/query", controllers.UserCtl.GetUserDetails)

	// help
	beego.Get("/v1/help/system/help", controllers.HelpCtl.Page)
	///////////////////////////////////////////////////////////////////////////

	beego.Get("/v1/auth/HandleLogsPage", controllers.HandleLogsCtl.Page)
	beego.Get("/v1/auth/domain/page", controllers.DomainCtl.Page)
	beego.Get("/v1/auth/batch/page", controllers.UserRolesCtl.Page)
	beego.Get("/v1/auth/resource/org/page", controllers.OrgCtl.Page)
	beego.Get("/v1/auth/resource/page", controllers.ResourceCtl.Page)
	beego.Get("/v1/auth/user/page", controllers.UserCtl.Page)
	beego.Get("/v1/auth/role/page", controllers.RoleCtl.Page)
	beego.Get("/v1/auth/swagger/page", controllers.SwaggerCtl.Page)

}
