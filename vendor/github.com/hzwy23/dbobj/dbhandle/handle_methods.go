package dbhandle

import (
	"database/sql"
	"fmt"
	"sync"
)

type instance func() DbObj

var (
	dbLock  = new(sync.RWMutex)
	Adapter = make(map[string]instance)
)

// Database handle function list
// Every database drive must implements this interface
//
type DbObj interface {
	// Query database
	Query(sql string, args ...interface{}) (*sql.Rows, error)

	// Query one row
	QueryRow(sql string, args ...interface{}) *sql.Row

	// Execute
	Exec(sql string, args ...interface{}) (sql.Result, error)

	// Begin transaction
	Begin() (*sql.Tx, error)

	// Prepare
	Prepare(query string) (*sql.Stmt, error)

	// GetDetails Error Code
	GetErrorCode(err error) string

	// GetDetails Message info
	GetErrorMsg(err error) string
}

// Function: register database instance
// Time: 2016-06-15
// Author: huangzhanwei
// this function service for database driver
func Register(dsn string, f instance) {
	dbLock.Lock()
	defer dbLock.Unlock()
	if f == nil {
		fmt.Errorf("sql: Register driver is nil")
	}
	if _, dup := Adapter[dsn]; dup {
		fmt.Println("reregister diver. dsn is :", dsn)
	}
	Adapter[dsn] = f
}
