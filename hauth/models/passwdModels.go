package models

import (
	"github.com/hzwy23/dbobj"
	"errors"
)

type PasswdModels struct {
}

func (PasswdModels) UpdateMyPasswd(newPd, User_id, oriEn string) error {
	flag,_,_,_:=CheckPasswd(User_id,oriEn)
	if !flag{
		return errors.New("原密码不正确，请确认后重新输入")
	}
	return dbobj.Exec(sys_rdbms_014, newPd, User_id, oriEn)
}

func (PasswdModels) UpdateUserPasswd(newPd, userid string) error {
	return dbobj.Exec(sys_rdbms_015, newPd, userid)
}
