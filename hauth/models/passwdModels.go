package models

import (
	"github.com/hzwy23/dbobj"
	"errors"
	"github.com/hzwy23/asofdate/hauth/hrpc"
)

type PasswdModels struct {
}

func (PasswdModels) UpdateMyPasswd(newPd, User_id, oriEn string) error {
	flag,_,_,_:= hrpc.CheckPasswd(User_id,oriEn)
	if !flag{
		return errors.New("原密码不正确，请确认后重新输入")
	}
	_,err := dbobj.Exec(sys_rdbms_014, newPd, User_id, oriEn)
	return err
}

func (PasswdModels) UpdateUserPasswd(newPd, userid string) error {
	_,err :=dbobj.Exec(sys_rdbms_015, newPd, userid)
	return err
}
