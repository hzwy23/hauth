package models

import (
	"github.com/hzwy23/dbobj"
)

type ThemeModel struct {
}

func (ThemeModel) Post(user_id, theme_id string) error {
	return dbobj.Exec(sys_rdbms_024, theme_id, user_id)
}
