package hrpc

var (
	sys_rdbms_hrpc_001 = `select f.authorization_level from sys_user_info t inner join sys_org_info i on t.org_unit_id = i.org_unit_id inner join sys_domain_share_info f on f.domain_id = ? and i.domain_id = f.target_domain_id where t.user_id = ?`
	sys_rdbms_hrpc_002 = `select t.domain_id from sys_org_info t where t.org_unit_id  = ?`
	sys_rdbms_hrpc_003 = `select i.domain_id from sys_user_info t inner join sys_org_info i on t.org_unit_id = i.org_unit_id where t.user_id = ?`
	sys_rdbms_hrpc_004 = `select domain_id from sys_role_info where role_id = ?`
	sys_rdbms_hrpc_005 = `select user_id,user_passwd,status_id,continue_error_cnt from sys_sec_user where user_id = ?`
	sys_rdbms_hrpc_006 = `select count(*) from sys_role_user_relation r inner join sys_role_resource_relat e on r.role_id = e.role_id inner join sys_theme_value v on e.res_id = v.res_id inner join sys_user_theme m on v.theme_id = m.theme_id and r.user_id = m.user_id where r.user_id = ? and v.res_url = ?`
	sys_rdbms_hrpc_007 = `update sys_sec_user set continue_error_cnt = ? where user_id = ?`
	sys_rdbms_hrpc_008 = `update sys_sec_user set status_id = 1 where user_id = ?`
)
