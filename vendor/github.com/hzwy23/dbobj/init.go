package dbobj

import (
	_ "github.com/hzwy23/dbobj/mysql"
	//_ "github.com/hzwy23/dbobj/oracle"
	"os"
	"path/filepath"
	"github.com/hzwy23/dbobj/utils"
)

func init(){
	HOME:=os.Getenv("HBIGDATA_HOME")
	filedir := filepath.Join(HOME,"conf/system.properties")
	conf,err := utils.GetConfig(filedir)
	if err != nil {
		panic("init database failed.")
	}
	Default,err = conf.Get("DB.type")
	if err != nil {
		panic("get default database type failed.")
	}
	InitDB(Default)
}