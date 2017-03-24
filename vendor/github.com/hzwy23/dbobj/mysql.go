package dbobj

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hzwy23/dbobj/utils/config"

	"github.com/hzwy23/dbobj/utils"

	_ "github.com/go-sql-driver/mysql"
)

type mysqldb struct {
	db *sql.DB
}

func NewMysqlDbObj() *mysqldb {

	var err error

	o := new(mysqldb)

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

	tnsname := usr + ":" + pad + "@" + tns

	o.db, err = sql.Open("mysql", tnsname)

	if err != nil {
		fmt.Errorf("open oracle database failed.", err)
		return nil
	}
	if len(pad) != 24 {
		psd, err := utils.Encrypt(pad)
		if err != nil {
			fmt.Errorf("decrypt passwd failed.", psd)
		}
		psd = "\"" + psd + "\""
		red.Set("DB.passwd", psd)
	}

	fmt.Println("create mysql obj success.")
	return o
}

func (this *mysqldb) GetErrorCode(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[:n])
	} else {
		fmt.Errorf("this error information is not mysql return info")
		return ""
	}
}

func (this *mysqldb) GetErrorMsg(err error) string {
	ret := err.Error()
	if n := strings.Index(ret, ":"); n > 0 {
		return strings.TrimSpace(ret[n+1:])
	} else {
		fmt.Errorf("this error information is not mysql return info")
		return ""
	}
}

func (this *mysqldb) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewMysqlDbObj()
		register("mysql", nd)
		this = nd
	}
	return this.db.Query(sql, args...)
}

func (this *mysqldb) Exec(sql string, args ...interface{}) error {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewMysqlDbObj()
		register("mysql", nd)
		this = nd
	}
	_, err := this.db.Exec(sql, args...)
	return err
}

func (this *mysqldb) Begin() (*sql.Tx, error) {
	return this.db.Begin()
}

func (this *mysqldb) Prepare(query string) (*sql.Stmt, error) {
	return this.db.Prepare(query)
}

func (this *mysqldb) QueryRow(sql string, args ...interface{}) *sql.Row {
	if this.db.Ping() != nil {
		fmt.Errorf("数据库连接已断开")
		nd := NewMysqlDbObj()
		register("mysql", nd)
		this = nd
	}
	return this.db.QueryRow(sql, args...)
}

func init() {
	register("mysql", NewMysqlDbObj())
}
