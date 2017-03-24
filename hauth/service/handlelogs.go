package service

import (
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/utils/logs"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
	"github.com/hzwy23/dbobj"
)

func WriteHandleLogs(ctx *context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("写入日志信息失败")
		}
	}()

	url := ctx.Request.URL.Path

	if strings.HasPrefix(url, "/") {

		sql := `insert into sys_handle_logs(uuid,user_id,handle_time,client_ip,status_code,method,url,domain_id,data) values(uuid(),?,now(),?,?,?,?,?,left(?,2999))`
		status := ctx.ResponseWriter.Status
		if status == 0 {
			status = 200
		}

		dt := ctx.Request.Form.Encode()
		cookie, _ := ctx.Request.Cookie("Authorization")
		jclaim, err := hjwt.ParseJwt(cookie.Value)
		if err != nil {
			dbobj.Exec(sql, "dev", ctx.Input.IP(), status, ctx.Request.Method, url, "dev", dt)
		} else {
			dbobj.Exec(sql, jclaim.User_id, ctx.Input.IP(), status, ctx.Request.Method, url, jclaim.Domain_id, dt)
		}
	}
}
