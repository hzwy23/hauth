package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/hzwy23/dbobj/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hzwy23/dbobj/dbhandle"
	"path/filepath"
)

type mysql struct {
	db *sql.DB
}

func NewMySQL() dbhandle.DbObj {

	var err error

	o := new(mysql)

	HOME := os.Getenv("HBIGDATA_HOME")

	filedir := filepath.Join(HOME, "conf", "system.properties")

	red, err := utils.GetResource(filedir)
	if err != nil {
		fmt.Errorf("cant not read ./conf/system.properties.please check this file.")
		return nil
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

	o.db, err = sql.Open("mysql", usr + ":" + pad + "@" + tns)

	if err != nil {
		fmt.Errorf("open oracle database failed.", err)
		return nil
	}
	if len(pad) != 24 {
		psd, err := utils.Encrypt(pad)
		if err != nil {
			fmt.Errorf("decrypt passwd failed.", psd)
			return nil
		}
		psd = "\"" + psd + "\""
		red.Set("DB.passwd", psd)
	}
	fmt.Println("create mysql dbhandle success.")
	return o
}

func (this *mysql) GetErrorCode(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[:n])
	} else {
		fmt.Errorf("this error information is not mysql return info")
		return ""
	}
}

func (this *mysql) GetErrorMsg(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[n+1:])
	} else {
		fmt.Errorf("this error information is not mysql return info")
		return ""
	}
}

func (this *mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		//nd := NewMySQL()
		//register("mysql", nd)
		//this = nd
	}
	return this.db.Query(sql, args...)
}

func (this *mysql) Exec(sql string, args ...interface{}) error {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		//nd := NewMySQL()
		//register("mysql", nd)
		//this = nd
	}
	_, err := this.db.Exec(sql, args...)
	return err
}

func (this *mysql) Begin() (*sql.Tx, error) {
	return this.db.Begin()
}

func (this *mysql) Prepare(query string) (*sql.Stmt, error) {
	return this.db.Prepare(query)
}

func (this *mysql) QueryRow(sql string, args ...interface{}) *sql.Row {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		//nd := NewMySQL()
		//register("mysql", nd)
		//this = nd
	}
	return this.db.QueryRow(sql, args...)
}

func init() {
	dbhandle.Register("mysql", NewMySQL)
}
