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
	"github.com/hzwy23/dbobj/dbhandle"
	"reflect"
	"strconv"
)

var (
	dbobj   dbhandle.DbObj
	Default = "mysql"
)

func InitDB(dbtyp string) error {
	if dbobj == nil {
		if val, ok := dbhandle.Adapter[dbtyp]; ok {
			dbobj = val()
		}
	}
	return nil
}

// get default DbObj name
// return DbObj name
func GetDefaultName() string {
	return Default
}

func Begin() (*sql.Tx, error) {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return nil, errors.New("can not connect database again.")
		}
		return dbobj.Begin()
	}
	return dbobj.Begin()
}

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return nil, errors.New("can not connect database again.")
		}
		return dbobj.Query(sql, args...)
	}
	return dbobj.Query(sql, args...)
}

func QueryRow(sql string, args ...interface{}) *sql.Row {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return nil
		}
		return dbobj.QueryRow(sql, args...)
	}
	return dbobj.QueryRow(sql, args...)
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return nil, errors.New("connect database failed.")
		}
		return dbobj.Exec(sql, args...)
	}
	return dbobj.Exec(sql, args...)
}

func Prepare(sql string) (*sql.Stmt, error) {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return nil, errors.New("can not connect database again.")
		}
		return dbobj.Prepare(sql)
	}
	return dbobj.Prepare(sql)
}

func GetErrorCode(errs error) string {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return err.Error()
		}
		return dbobj.GetErrorCode(errs)
	}
	return dbobj.GetErrorCode(errs)
}

func GetErrorMsg(errs error) string {
	if dbobj == nil {
		err := InitDB(Default)
		if err != nil {
			return err.Error()
		}
		return dbobj.GetErrorMsg(errs)
	}
	return dbobj.GetErrorMsg(errs)
}

// Function: scan DbObj query result
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
