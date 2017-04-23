package hrpc

import (
	"database/sql"
	"github.com/hzwy23/dbobj"
)

type mSysUserSec struct {
	User_id                 string        `json:"user_id"`
	User_passwd             string        `json:"user_passwd"`
	User_status             sql.NullInt64 `json:"user_status"`
	User_continue_error_cnt sql.NullInt64
}

// check user's passwd is right.
func CheckPasswd(user_id, user_passwd string) (bool, int, int64, string) {
	var sec mSysUserSec
	err := dbobj.QueryRow(sys_rdbms_hrpc_005, user_id).Scan(&sec.User_id, &sec.User_passwd, &sec.User_status, &sec.User_continue_error_cnt)
	if err != nil {
		return false, 402, 0, "error_querydb"
	}

	if sec.User_status.Int64 != 0 {
		return false, 406, sec.User_status.Int64, "error_maxerror"
	}

	if sec.User_continue_error_cnt.Int64 > 6 {
		forbidUsers(user_id)
		return false, 403, sec.User_continue_error_cnt.Int64, "error_maxerror"
	}

	if sec.User_id == user_id && sec.User_passwd == user_passwd {
		updateContinueErrorCnt(0, user_id)
		return true, 200, 0, ""
	} else {
		updateContinueErrorCnt(sec.User_continue_error_cnt.Int64+1, user_id)
		return false, 405, sec.User_continue_error_cnt.Int64 + 1, "error_password"
	}
}

func updateContinueErrorCnt(cnt int64, user_id string) {
	dbobj.Exec(sys_rdbms_hrpc_007, cnt, user_id)
}

func forbidUsers(user_id string) {
	dbobj.Exec(sys_rdbms_hrpc_008, user_id)
}
