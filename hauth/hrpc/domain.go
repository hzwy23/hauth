package hrpc

import (
	"net/http"
	"github.com/hzwy23/utils/jwt"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/dbobj"
)


// check the user wheather handle the domain
// return value :
// -1   : have no right to handle the domain
// 1    : can read the domain info
// 2    : can read and wirte the domain info
func GetAuthLevel(user_id string, domain_id string) int {
	var cnt = -1
	err := dbobj.QueryRow(sys_rdbms_hrpc_001, domain_id, user_id).Scan(&cnt)
	if err != nil {
		logs.Error(err)
		return -1
	}
	return cnt
}

// 根据用户账号,获取用户所在的域
func GetDomainId(user_id string) (string, error) {
	domain_id := ""
	err := dbobj.QueryRow(sys_rdbms_hrpc_003, user_id).Scan(&domain_id)
	return domain_id, err
}


// 返回值是-1 表示没有读写权限
// 返回值是1 表示有读取权限，没有写入权限
// 返回值是2 表示有读写权限
func checkDomainAuthLevel(req *http.Request, domain_id string) int {
	level := -1
	cookie, _ := req.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		logs.Error(err)
		return level
	}

	// if the user is not admin, and user_id is not owner this domain_id
	// check share info. or not
	if jclaim.User_id != "admin" && jclaim.Domain_id != domain_id {
		level = GetAuthLevel(jclaim.User_id, domain_id)
		return level
	} else {
		return 2
	}
}
