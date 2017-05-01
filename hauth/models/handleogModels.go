package models

import (
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/utils/logs"
	"github.com/hzwy23/utils/validator"
)

type HandleLogMode struct {
}

type handleLogs struct {
	Uuid        string `json:"uuid"`
	User_id     string `json:"user_id"`
	Handle_time string `json:"handle_time"`
	Client_ip   string `json:"client_ip"`
	Status_code string `json:"status_code"`
	Method      string `json:"method"`
	Url         string `json:"url"`
	Data        string `json:"data"`
}

func (this HandleLogMode) Download(domain_id string) ([]handleLogs, error) {
	var rst []handleLogs
	rows, err := dbobj.Query(sys_rdbms_012, domain_id)
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return rst, nil
}

func (this HandleLogMode) getTotal(domain_id string) (total int64, err error) {
	err = dbobj.QueryRow(sys_rdbms_030, domain_id).Scan(&total)
	return
}

func (this HandleLogMode) Get(domain_id, offset, limit string) ([]handleLogs, int64, error) {
	var rst []handleLogs
	rows, err := dbobj.Query(sys_rdbms_029, domain_id, offset, limit)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	err = dbobj.Scan(rows, &rst)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	total, err := this.getTotal(domain_id)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return rst, total, nil
}

func (this HandleLogMode) Search(domain_id, userid, start, end string) ([]handleLogs, error) {
	var rst []handleLogs
	if userid != "" && validator.IsDate(start) && validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_031, domain_id, userid, start, end)
		defer rows.Close()
		if err != nil {
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if userid != "" && validator.IsDate(start) {

		rows, err := dbobj.Query(sys_rdbms_032, domain_id, userid, start)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if userid != "" && validator.IsDate(end) {

		rows, err := dbobj.Query(sys_rdbms_031, domain_id, userid, start, end)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if validator.IsDate(start) && validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_033, domain_id, start, end)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if validator.IsDate(start) {
		rows, err := dbobj.Query(sys_rdbms_035, domain_id, start, end)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if validator.IsDate(end) {
		rows, err := dbobj.Query(sys_rdbms_039, domain_id, start, end)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else if userid != "" {
		rows, err := dbobj.Query(sys_rdbms_040, domain_id, userid)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	} else {
		rows, err := dbobj.Query(sys_rdbms_042, domain_id)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		err = dbobj.Scan(rows, &rst)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
	}
	return rst, nil
}
