package hrpc

// hrpc package
// this package provide permissions related function
import (
	"net/http"

	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/utils/validator"
)

// 校验用户是否有权限访问当前API
func BasicAuth(r *http.Request) bool {
	cookie, _ := r.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		return false
	}
	if jclaim.User_id == "admin" {
		return true
	}
	cnt := 0
	err = dbobj.QueryRow(sys_rdbms_hrpc_006, jclaim.User_id, r.URL.Path).Scan(&cnt)
	if err != nil {
		logs.Error(err)
		return false
	}
	if cnt == 0 {
		logs.Error("insufficient privileges", "user id is :", jclaim.User_id, "api is :", r.URL.Path)
		return false
	}
	return true
}

// 检查用户对指定的域的权限
// 第一个参数中,http.Request,包含了用户的连接信息,cookie中.
// 第二个参数中,domain_id,是用户想要访问的域
// 第三个参数是访问模式,r 表示 只读, w 表示 读写.
// 如果返回true,表示用户有权限
// 返回false,表示用户没有权限
func DomainAuth(req *http.Request, domain_id string, pattern string) bool {
	if validator.IsEmpty(domain_id) {
		return false
	}

	level := checkDomainAuthLevel(req, domain_id)
	switch pattern {
	case "r":
		if level != -1 {
			return true
		} else {
			return false
		}
	case "w":
		if level != 2 {
			return false
		} else {
			return true
		}
	default:
		return false
	}
}

func IsRoot(domainId string) bool {
	if domainId == "vertex_root" {
		return true
	}
	return false
}
