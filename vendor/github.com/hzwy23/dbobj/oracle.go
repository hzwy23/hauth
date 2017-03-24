package dbobj

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hzwy23/dbobj/utils/config"

	"github.com/hzwy23/dbobj/utils"

	//	_ "github.com/mattn/go-oci8"
)

type dbobj struct {
	db *sql.DB
}

func NewOraDbObj() *dbobj {

	var err error

	o := new(dbobj)

	nlsLang := os.Getenv("NLS_LANG")
	if !strings.HasSuffix(nlsLang, "UTF8") {
		os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	}

	HOME := os.Getenv("HBIGDATA_HOME")
	if HOME == ""{
		HOME = "./"
	}

	filedir := path.Join(HOME, "conf", "system.properties")
	red, err := config.GetResource(filedir)
	if err != nil {
		fmt.Errorf("cant not read ./conf/system.properties.please check this file.")
	}

	tns := red.Conf["DB.tns"]
	usr := red.Conf["DB.user"]
	pad := red.Conf["DB.passwd"]

	if len(pad) == 24 {
		pad, err = utils.Decrypt(pad)
		if err != nil {
			fmt.Errorf("Decrypt mysql passwd failed.")
			return nil
		}
	}

	tnsname := usr + "/" + pad + "@" + tns

	o.db, err = sql.Open("oci8", tnsname)

	if err != nil {
		fmt.Errorf("open oracle database failed.%v", err)
		return nil
	}
	if len(pad) != 24 {
		psd, err := utils.Encrypt(pad)
		if err != nil {
			fmt.Errorf("decrypt passwd failed.%v", psd)
		}
		psd = "\"" + psd + "\""
		red.Set("DB.passwd", psd)
	}
	o.db.SetMaxOpenConns(0)
	o.db.SetConnMaxLifetime(0)
	fmt.Println("create Oracle obj success.")
	return o
}

func (this *dbobj) GetErrorCode(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[:n])
	} else {
		fmt.Println("this error information is not Oracle return info")
		return ""
	}
}

func (this *dbobj) GetErrorMsg(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[n+1:])
	} else {
		fmt.Println("this error information is not Oracle return info")
		return ""
	}
}

func (this *dbobj) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewOraDbObj()
		register("oracle", nd)
		this = nd
	}
	return this.db.Query(sql, args...)
}

func (this *dbobj) Exec(sql string, args ...interface{}) error {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewOraDbObj()
		register("oracle", nd)
		this = nd
	}
	_, err := this.db.Exec(sql, args...)
	return err
}

func (this *dbobj) Begin() (*sql.Tx, error) {
	return this.db.Begin()
}

func (this *dbobj) Prepare(query string) (*sql.Stmt, error) {
	return this.db.Prepare(query)
}

func (this *dbobj) QueryRow(sql string, args ...interface{}) *sql.Row {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewOraDbObj()
		register("oracle", nd)
		this = nd
	}
	return this.db.QueryRow(sql, args...)
}

func init() {
	register("oracle", NewOraDbObj())
}
