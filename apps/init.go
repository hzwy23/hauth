package apps

import (
	"github.com/hzwy23/asofdate/apps/mas/ftp"
	"github.com/hzwy23/asofdate/apps/mas/ca"
	"github.com/hzwy23/asofdate/apps/mas/common"
)

// 在init中注册自己应用的路由
func init(){
	// 系统开启路由服务之前,会先执行这个init中的函数
	// 系统默认采用beego的路由
	//ca.Register()
	//ftp.Register()
	//common.Register()
}
