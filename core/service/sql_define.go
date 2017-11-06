package service

var hauth_service_001 = `insert into sys_handle_logs(uuid,user_id,handle_time,client_ip,status_code,method,url,domain_id,data) values(uuid(),?,now(),?,?,?,?,?,left(?,2999))`
