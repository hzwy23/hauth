package service

import (
	"net/http"

	"github.com/hzwy23/utils/jwt"
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

func CheckConnection(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Authorization")
	if err != nil || !jwt.CheckToken(cookie.Value) {
		w.Write([]byte(redirect))
	}
}
