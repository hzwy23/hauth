package service

import (
	"github.com/astaxie/beego/context"
	"github.com/hzwy23/asofdate/utils/token/hjwt"
)

const redirect = `
<script type="text/javascript">
    $.Hconfirm({
		cancelBtn:false,
        header:"连接已断开",
        body:"<span style='font-weight:600;font-size:16px;padding-left:60px;height:90px;line-height:90px;'>用户连接已断开，请重新登录</span>",
        callback:function () {
            window.location.href="/"
        }
    })
</script>
`

func CheckJWT(ctx *context.Context) {
	cookie, err := ctx.Request.Cookie("Authorization")
	if err != nil || !hjwt.CheckToken(cookie.Value) {
		ctx.WriteString(redirect)
	}
}
