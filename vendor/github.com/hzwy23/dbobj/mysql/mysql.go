package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/hzwy23/dbobj/utils"

	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hzwy23/dbobj/dbhandle"
)

type mysql struct {
	db *sql.DB
}

func NewMySQL() dbhandle.DbObj {

	var err error

	o := new(mysql)

	HOME := os.Getenv("HBIGDATA_HOME")

	filedir := filepath.Join(HOME, "conf", "asofdate.conf")

	red, err := utils.GetResource(filedir)
	if err != nil {
		fmt.Errorf("cant not read ./conf/asofdate.conf.please check this file.")
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

	o.db, err = sql.Open("mysql", usr+":"+pad+"@"+tns)

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
	rows, err := this.db.Query(sql, args...)
	if err != nil {
		if this.db.Ping() != nil {
			// if dbobj connection is broken,
			// reconnect database.
			fmt.Errorf("%s", "Connection is broken")
			if val, ok := NewMySQL().(*mysql); ok {
				this.db = val.db
			}
			return this.db.Query(sql, args...)
		}
	}
	return rows, err
}

func (this *mysql) Exec(sql string, args ...interface{}) (sql.Result, error) {
	result, err := this.db.Exec(sql, args...)
	if err != nil {
		if this.db.Ping() != nil {
			// if dbobj connection is broken,
			// reconnect database.
			fmt.Errorf("%s", "Connection is broken")
			if val, ok := NewMySQL().(*mysql); ok {
				this.db = val.db
			}
			return this.db.Exec(sql, args...)
		}
	}
	return result, err
}

func (this *mysql) Begin() (*sql.Tx, error) {
	tx, err := this.db.Begin()
	if err != nil {
		// if dbobj connection is broken,
		// reconnect database.
		if this.db.Ping() != nil {
			fmt.Errorf("%s", "Connection is broken")
			if val, ok := NewMySQL().(*mysql); ok {
				this.db = val.db
			}
			return this.db.Begin()
		}
	}
	return tx, err
}

func (this *mysql) Prepare(sql string) (*sql.Stmt, error) {
	stmt, err := this.db.Prepare(sql)
	if err != nil {
		// if dbobj connection is broken,
		// reconnect database.
		if this.db.Ping() != nil {
			fmt.Errorf("%s", "Connection is broken")
			if val, ok := NewMySQL().(*mysql); ok {
				this.db = val.db
			}
			return this.db.Prepare(sql)
		}
	}
	return stmt, err
}

func (this *mysql) QueryRow(sql string, args ...interface{}) *sql.Row {
	if this.db.Ping() != nil {
		fmt.Errorf("%s", "Connection is broken")
		if val, ok := NewMySQL().(*mysql); ok {
			this.db = val.db
		}
	}
	return this.db.QueryRow(sql, args...)
}

func init() {
	dbhandle.Register("mysql", NewMySQL)
}
