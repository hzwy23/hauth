package service

import "github.com/hzwy23/dbobj"

func init() {
	defdb := dbobj.GetDefaultName()
	if "oracle" == defdb {
		hauth_service_001 = `insert into sys_handle_logs(uuid,user_id,handle_time,client_ip,status_code,method,url,domain_id,data) values(sys_guid(),:1,sysdate,:2,:3,:4,:5,:6,:7)`
	}
}
