// Copyright 2016 huangzhanwei. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dbobj

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"sync"

	"github.com/hzwy23/dbobj/utils/config"
)

// Database handle function list
// Every database drive must implements this interface
//
type db interface {
	// Query database
	Query(sql string, args ...interface{}) (*sql.Rows, error)

	// Query one row
	QueryRow(sql string, args ...interface{}) *sql.Row

	// Execute
	Exec(sql string, args ...interface{}) error

	// Begin transaction
	Begin() (*sql.Tx, error)

	// Prepare
	Prepare(query string) (*sql.Stmt, error)

	// Get Error Code
	GetErrorCode(err error) string

	// Get Message info
	GetErrorMsg(err error) string
}

var (
	dbLock    = new(sync.RWMutex)
	dbPool    = make(map[string]db)
	defaultdb = initDefaultDB()
)

func initDefaultDB() string {

	HOME := os.Getenv("HBIGDATA_HOME")
	if HOME == "" {
		HOME = "./"
	}

	filedir := path.Join(HOME, "conf", "system.properties")

	red, err := config.GetConfig(filedir)
	if err != nil {

		fmt.Errorf("cant not read ./conf/system.properties.please check this file.")

	}

	confdb, err := red.Get("DB.type")
	if err != nil {

		fmt.Errorf("Get default failed. set default db mysql")
		return "mysql"

	}

	return confdb
}

// Function: register database instance
// Time: 2016-06-15
// Author: huangzhanwei
// this function service for database driver
func register(dsn string, d db) {
	dbLock.Lock()
	defer dbLock.Unlock()
	if d == nil {
		fmt.Errorf("sql: Register driver is nil")
	}
	if _, dup := dbPool[dsn]; dup {
		fmt.Println("reregister diver. dsn is :", dsn)
	}
	dbPool[dsn] = d
	if dsn == defaultdb {
		fmt.Println("default db is ", dsn)
	} else {
		fmt.Println(dsn, "register success.")
	}
}

//get other db obj
func GetdbObj(name string) (db, error) {
	if val, ok := dbPool[name]; ok {
		return val, nil
	} else {
		return nil, errors.New(name + " was not register.")
	}
}

// get all dbpool
// return map contains all db obj
func GetdbPool() map[string]db {
	return dbPool
}

// get default db name
// return db name
func GetDefaultName() string {
	return defaultdb
}

func Begin() (*sql.Tx, error) {
	if val, ok := dbPool[defaultdb]; ok {
		return val.Begin()
	}
	return nil, errors.New("can not found dbname, please register first.")
}

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if val, ok := dbPool[defaultdb]; ok {
		return val.Query(sql, args...)
	}
	return nil, errors.New("can not found dbname,please register first")
}

func QueryRow(sql string, args ...interface{}) *sql.Row {
	if val, ok := dbPool[defaultdb]; ok {
		return val.QueryRow(sql, args...)
	}
	return nil
}

func Exec(sql string, args ...interface{}) error {
	if val, ok := dbPool[defaultdb]; ok {
		return val.Exec(sql, args...)
	}
	return errors.New("can not found dbname,please register first")
}

func Prepare(query string) (*sql.Stmt, error) {
	if val, ok := dbPool[defaultdb]; ok {
		return val.Prepare(query)
	}
	return nil, errors.New("can not found dbname,please register first")
}

func GetErrorCode(err error) string {
	if val, ok := dbPool[defaultdb]; ok {
		return val.GetErrorCode(err)
	}
	return ""
}

func GetErrorMsg(err error) string {
	if val, ok := dbPool[defaultdb]; ok {
		return val.GetErrorMsg(err)
	}
	return ""
}

// Function: scan db query result
// Time: 2016-06-10
// Author: huangzhanwei
// Notice: second argument of type must be valid pointer.
func Scan(rows *sql.Rows, rst interface{}) error {
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Errorf("%v", err)
		return err
	}
	size := len(cols)
	values := make([]sql.NullString, size)
	scanArgs := make([]interface{}, size)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	obj := reflect.ValueOf(rst)
	if obj.Kind() != reflect.Ptr || obj.IsNil() {
		fmt.Errorf("second must be valid pointer")
		return errors.New("second argument of type must be valid pointer")
	}
	switch obj.Elem().Kind() {
	case reflect.Slice:
		obj = obj.Elem()
		var i = 0
		for rows.Next() {
			if i >= obj.Cap() {
				newcap := obj.Cap() + obj.Cap()/2
				if newcap < 4 {
					newcap = 4
				}
				newv := reflect.MakeSlice(obj.Type(), obj.Len(), newcap)
				reflect.Copy(newv, obj)
				obj.Set(newv)
			}
			if i >= obj.Len() {
				obj.SetLen(i + 1)
				if max := obj.Index(0).NumField(); max < size {
					fmt.Errorf("slice colunms less then dest ", max, size)
					return errors.New("slice colunms less then dest. numFiled is :" + strconv.Itoa(max) + ". need numFiled " + strconv.Itoa(size))
				}
			}
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Errorf("%v", err)
				return err
			}
			for index, vals := range values {
				obj.Index(i).Field(index).SetString(vals.String)
			}
			i++
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		for rows.Next() {
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Errorf("%v", err)
				return err
			}
			for _, vals := range values {
				ret, err := strconv.Atoi(vals.String)
				if err != nil {
					fmt.Errorf("%v", err)
					return err
				}
				obj.Elem().SetInt(int64(ret))
			}
		}
		return nil
	case reflect.String:
		for rows.Next() {
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Errorf("%v", err)
				return err
			}
			for _, vals := range values {
				obj.Elem().SetString(vals.String)
			}
		}
		return nil
	case reflect.Float32, reflect.Float64:
		for rows.Next() {
			err := rows.Scan(scanArgs...)
			if err != nil {
				fmt.Errorf("%v", err)
				return err
			}
			for _, vals := range values {
				ret, err := strconv.ParseFloat(vals.String, 64)
				if err != nil {
					fmt.Errorf("%v", err)
					return err
				}
				obj.Elem().SetFloat(ret)
			}
		}
		return nil
	default:
		return errors.New("second argument of type must be slice")
	}
	return nil
}

// Count total rows by sqlText
//
func Count(sql string, args ...interface{}) int64 {
	var total int64
	err := QueryRow(sql, args...).Scan(&total)
	if err != nil {
		fmt.Errorf("%v", err)
		return 0
	}
	return total
}
