package dbobj

import (
	"database/sql"
	"dbobj"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	type RTrst struct {
		Va1 string
		Va2 string
		Va3 string
		Va4 string
		Va5 string
		Va6 string
		Va7 string
		Va8 string
		Va9 string
		//	Va12 string
		//	Va13 string
		//	Va14 string
		//	Va15 string
	}

	type Trst struct {
		Va1 sql.NullString
		Va2 sql.NullString
		Va3 sql.NullString
		Va4 sql.NullString
		Va5 sql.NullString
		Va6 sql.NullString
		Va7 sql.NullString
		Va8 sql.NullString
		Va9 sql.NullString
		//	Va12 string
		//	Va13 string
		//	Va14 string
		//	Va15 string
	}
	rows, err := dbobj.Default.Query("select * from FTP_BUSIZ_METHOD_RELATION t")
	defer rows.Close()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	var one Trst
	var tmp RTrst
	var rst []RTrst
	for rows.Next() {
		err := rows.Scan(&one.Va1,
			&one.Va2,
			&one.Va3,
			&one.Va4,
			&one.Va5,
			&one.Va6,
			&one.Va7,
			&one.Va8,
			&one.Va9)
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}
		tmp.Va1 = one.Va1.String
		tmp.Va2 = one.Va2.String
		tmp.Va3 = one.Va3.String
		tmp.Va4 = one.Va4.String
		tmp.Va5 = one.Va5.String
		tmp.Va6 = one.Va6.String
		tmp.Va7 = one.Va7.String
		tmp.Va8 = one.Va8.String
		tmp.Va9 = one.Va9.String

		rst = append(rst, tmp)
	}
	for _, val := range rst {
		fmt.Println(val)
	}
}

func TestScanDB(t *testing.T) {
	type Trst struct {
		Va1  string
		Va2  string
		Va3  string
		Va4  string
		Va5  string
		Va6  string
		Va7  string
		Va8  string
		Va9  string
		Va10 string
		Va11 string
		//	Va12 string
		//	Va13 string
		//	Va14 string
		//	Va15 string
	}
	rows, err := dbobj.Default.Query("select * from FTP_BUSIZ_METHOD_RELATION t")
	defer rows.Close()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	var one []Trst
	dbobj.Scan(rows, &one)
	for _, val := range one {
		fmt.Println(val)
	}
}
