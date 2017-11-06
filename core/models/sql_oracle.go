package models

import (
	"github.com/hzwy23/dbobj"
)

func init() {
	defdb := dbobj.GetDefaultName()
	if "oracle" == defdb {
		sys_rdbms_001 = `select f.authorization_level from sys_user_info t inner join sys_org_info i on t.org_unit_id = i.org_unit_id inner join sys_domain_share_info f on f.domain_id = :1 and i.domain_id = f.target_domain_id where t.user_id = :2`
		sys_rdbms_002 = `select t.domain_id from sys_org_info t where t.org_unit_id  = :1`
		sys_rdbms_003 = `select i.domain_id from sys_user_info t inner join sys_org_info i on t.org_unit_id = i.org_unit_id where t.user_id = :1`
		sys_rdbms_004 = `select domain_id from sys_role_info where role_id = :1`
		sys_rdbms_005 = `update sys_resource_info set res_name = :1 where res_id = :2`
		sys_rdbms_006 = `select count(*) from sys_theme_value where theme_id = :1 and res_id = :2`
		sys_rdbms_007 = `delete from sys_user_info where user_id = :1 and org_unit_id = :2`
		sys_rdbms_008 = `insert into sys_theme_value(uuid,theme_id,res_id,res_url,res_type,res_bg_color,res_class,group_id,res_img,sort_id) value(uuid(),:1,:2,:3,:4,:5,:6,:7,:8,:9)`
		sys_rdbms_009 = `update sys_theme_value set res_url = :1, res_bg_color = :2, res_class = :3, res_img = :4, group_id = :5, sort_id = :6, res_type = :7 where theme_id = :8 and res_id = :9`
		sys_rdbms_010 = `select user_id,user_passwd,status_id,continue_error_cnt from sys_sec_user where user_id = :1`
		sys_rdbms_011 = `select distinct t2.res_url from sys_user_theme t1 inner join sys_theme_value t2 on t1.theme_id = t2.theme_id inner join sys_resource_info t3 on t2.res_id = t3.res_id where t1.user_id = :1 and t2.res_id = :2 and t3.res_type = '0'`
		sys_rdbms_012 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 order by handle_time desc`
		sys_rdbms_013 = `select res_type from sys_resource_info where res_id = :1`
		sys_rdbms_014 = `update sys_sec_user set user_passwd = :1 where user_id = :2 and user_passwd = :3`
		sys_rdbms_015 = `update sys_sec_user set user_passwd = :1 where user_id = :2`
		sys_rdbms_016 = `update sys_sec_user set status_id = :1 ,continue_error_cnt = '0' where user_id = :2`
		sys_rdbms_017 = `select t.user_id,t.user_name,a.status_desc,t.user_create_date, t.user_owner,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc, di.domain_id,di.domain_name,t.user_maintance_date,t.user_maintance_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id inner join sys_domain_info di on i.domain_id = di.domain_id where di.domain_id = :1`
		sys_rdbms_018 = `insert into sys_user_info (user_id,user_name,user_create_date,user_owner,user_email,user_phone,org_unit_id,user_maintance_date,user_maintance_user) values(:1,:2,now(),:3,:4,:5,:6,now(),:7)`
		sys_rdbms_019 = `insert into sys_sec_user(user_id,user_passwd,status_id) values(:1,:2,:3)`
		sys_rdbms_020 = `update sys_sec_user set user_passwd = :1 where user_id = :2`
		sys_rdbms_021 = `update sys_user_info t set t.user_name = :1, t.user_phone = :2, t.user_email = :3 ,t.user_maintance_date = now(), t.user_maintance_user = :4,t.org_unit_id = :5 where t.user_id = :6`
		sys_rdbms_022 = `select count(*) from sys_role_user_relation r inner join sys_role_resource_relat e on r.role_id = e.role_id inner join sys_theme_value v on e.res_id = v.res_id inner join sys_user_theme m on v.theme_id = m.theme_id and r.user_id = m.user_id where r.user_id = :1 and v.res_url = :2`
		sys_rdbms_023 = `select t.user_id,t.user_name,a.status_desc,t.user_create_date, t.user_owner,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc,di.domain_id,di.domain_name,t.user_maintance_date,t.user_maintance_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id inner join sys_domain_info di on i.domain_id = di.domain_id where t.user_id = :1`
		sys_rdbms_024 = `update sys_user_theme set theme_id = :1 where user_id = :2`
		sys_rdbms_025 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.domain_create_date  as maintance_date, t.domain_owner as user_id,t.domain_maintance_date,t.domain_maintance_user,t.domain_status_id from sys_domain_info t inner join sys_domain_status_attr s on t.domain_status_id = s.domain_status_id`
		sys_rdbms_026 = `insert into sys_role_info(role_id,role_name,role_owner,role_create_date,role_status_id,domain_id,role_maintance_date,role_maintance_user,code_number) values(:1,:2,:3,now(),:4,:5,now(),:6,:7)`
		sys_rdbms_027 = `delete from sys_role_info where role_id = :1 and domain_id = :2`
		sys_rdbms_028 = `select t.code_number,t.role_name,t.role_owner,t.role_create_date,a.role_status_desc,a.role_status_id,t.domain_id,o.domain_name,t.role_maintance_date,t.role_maintance_user,t.role_id from sys_role_info t inner join sys_role_status_attr a on t.role_status_id = a.role_status_id inner join sys_domain_info o on t.domain_id = o.domain_id where t.domain_id = :1`
		sys_rdbms_029 = `select uuid, user_id, handle_time, client_ip, status_code, method, url, data from (select b.*,rownum rn from (select a.*,rownum as rk  from ( select uuid, user_id, handle_time, client_ip, status_code, method, url, data from sys_handle_logs t where t.domain_id = :1 order by handle_time desc ) a ) b where b.rk > :2 ) c where c.rn < :3`
		sys_rdbms_030 = `select count(*) from sys_handle_logs t where t.domain_id = :1`
		sys_rdbms_031 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and user_id = :2 and handle_time >= str_to_date(:3,'%Y-%m-%d') and handle_time < str_to_date(:4,'%Y-%m-%d') order by handle_time desc`
		sys_rdbms_032 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and user_id = :2 and handle_time >= str_to_date(:3,'%Y-%m-%d') order by handle_time desc`
		sys_rdbms_033 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and handle_time >= str_to_date(:2,'%Y-%m-%d') and handle_time < str_to_date(:3,'%Y-%m-%d') order by handle_time desc`
		sys_rdbms_034 = `select domain_id from sys_domain_share_info t where t.target_domain_id = :1`
		sys_rdbms_035 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and handle_time >= str_to_date(:2,'%Y-%m-%d') order by handle_time desc`
		sys_rdbms_036 = `insert into sys_domain_info(domain_id,domain_name,domain_status_id,domain_create_date,domain_owner,domain_maintance_date,domain_maintance_user) values(:1,:2,:3,now(),:4,now(),:5)`
		sys_rdbms_037 = `delete from sys_domain_info where domain_id = :1`
		sys_rdbms_038 = `update sys_domain_info set domain_name = :1, domain_status_id = :2, domain_maintance_date = now(), domain_maintance_user = :3 where domain_id = :4`
		sys_rdbms_039 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and handle_time < str_to_date(:2,'%Y-%m-%d') order by handle_time desc`
		sys_rdbms_040 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 and user_id = :2 order by handle_time desc`
		sys_rdbms_041 = `select org_unit_id,org_unit_desc,up_org_id,t.domain_id,create_date,maintance_date,create_user,maintance_user,code_number from sys_org_info t where t.domain_id = :1`
		sys_rdbms_042 = `select uuid,user_id,handle_time,client_ip,status_code,method,url,data from sys_handle_logs t where t.domain_id = :1 order by user_id,handle_time desc`
		sys_rdbms_043 = `insert into sys_org_info(code_number,org_unit_desc,up_org_id,domain_id,create_date,maintance_date,create_user,maintance_user,org_unit_id) values(:1,:2,:3,:4,now(),now(),:5,:6,:7)`
		sys_rdbms_044 = `delete from sys_org_info where org_unit_id = :1 and domain_id = :2`
		sys_rdbms_045 = `insert into sys_user_theme(user_id,theme_id) values(:1,:2)`
		sys_rdbms_046 = `select t.role_id,t.role_name,t.code_number from sys_role_info t where ( t.role_owner = :1 or exists ( select 1 from sys_role_user_relation r where r.user_id = :1 and t.role_id = r.role_id ))`
		sys_rdbms_047 = `select t.role_id,t.role_name,t.code_number from sys_role_info t where ( t.role_owner = :1 or exists ( select 1 from sys_role_user_relation r where r.user_id = :2 and t.role_id = r.role_id )) and not exists (select 1 from sys_role_user_relation n where n.user_id = :3 and t.role_id = n.role_id )`
		sys_rdbms_048 = `insert into sys_role_user_relation(uuid,role_id,user_id,maintance_date,maintance_user) values(uuid(),:1,:2,now(),:3)`
		sys_rdbms_050 = `update sys_role_info t set t.role_name = :1 ,t.role_status_id = :2, role_maintance_date = now(), role_maintance_user = :3 where t.role_id = :4`
		sys_rdbms_069 = `update sys_org_info set org_unit_desc = :1 ,up_org_id = :2, maintance_date = now(),maintance_user=:3 where org_unit_id = :4`
		sys_rdbms_070 = `select t.theme_id,i.theme_desc, res_id,res_url,res_type,res_bg_color,res_class,group_id,res_img,sort_id from sys_theme_value t left join sys_theme_info i on t.theme_id = i.theme_id where t.theme_id = :1 and t.res_id = :2`
		sys_rdbms_071 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc,t.sys_flag from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type`
		sys_rdbms_072 = `insert into sys_resource_info(res_id,res_name,res_attr,res_up_id,res_type) values(:1,:2,:3,:4,:5)`
		sys_rdbms_073 = `insert into sys_theme_value(uuid,theme_id,res_id,res_url,res_type,res_bg_color,res_class,group_id,res_img,sort_id) values(uuid(),:1,:2,:3,:4,:5,:6,:7,:8,:9)`
		sys_rdbms_074 = `insert into sys_role_resource_relat(uuid,role_id,res_id) values(uuid(),:1,:2)`
		sys_rdbms_075 = `delete from sys_role_resource_relat where res_id = :1`
		sys_rdbms_076 = `delete from sys_theme_value where res_id = :1`
		sys_rdbms_077 = `delete from sys_resource_info where res_id = :1`
		sys_rdbms_078 = `select t1.res_url from sys_index_page t1 inner join sys_user_theme t2 on t1.theme_id = t2.theme_id where t2.user_id = :1`
		sys_rdbms_079 = `select distinct domain_id from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id where user_id = :1`
		sys_rdbms_080 = `select o.org_unit_id from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id where user_id = :1`
		sys_rdbms_083 = `select t.uuid,t.target_domain_id,i.domain_name,t.authorization_level,t.create_user,t.create_date,t.modify_user,t.modify_date from sys_domain_share_info t inner join sys_domain_info i on t.target_domain_id = i.domain_id where t.domain_id = :1`
		sys_rdbms_084 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.domain_create_date  as maintance_date, t.domain_owner as user_id,t.domain_maintance_date,t.domain_maintance_user from sys_domain_info t inner join sys_domain_status_attr s  on t.domain_status_id = s.domain_status_id where t.domain_id = :1`
		sys_rdbms_085 = `select t.domain_id as project_id, t.domain_name as project_name from sys_domain_info t where not exists ( select 1 from sys_domain_share_info i where t.domain_id = i.target_domain_id and i.domain_id = :1 )`
		sys_rdbms_086 = `insert into sys_domain_share_info(uuid,domain_id,target_domain_id,authorization_level,create_user,create_date,modify_date,modify_user) values(uuid(),:1,:2,:3,:4,now(),now(),:5)`
		sys_rdbms_087 = `delete from sys_domain_share_info where uuid = :1 and domain_id = :2`
		sys_rdbms_088 = `update sys_domain_share_info set authorization_level = :1,modify_user = :2 , modify_date = now() where uuid = :3`
		sys_rdbms_089 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type where res_id = :1`
		sys_rdbms_093 = `delete from sys_role_resource_relat where role_id = :1 and res_id = :2`
		sys_rdbms_094 = `select r.user_id, t.role_id, t.code_number,t.role_name,t.role_status_id from sys_role_info t inner join sys_role_user_relation r on t.role_id = r.role_id where r.user_id = :1 and t.role_status_id = '0'`
		sys_rdbms_095 = `select '', t.role_id,t.code_number,t.role_name from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id inner join sys_role_info t on o.domain_id = t.domain_id where i.user_id = :1 and t.role_status_id = '0' and  not exists ( select 1 from sys_role_user_relation r where i.user_id = r.user_id and r.role_id = t.role_id )`
		sys_rdbms_096 = `insert into sys_role_user_relation(uuid,role_id,user_id,maintance_date,maintance_user) values(:1,:2,:3,sysdate,:4)`
		sys_rdbms_097 = `delete from sys_role_user_relation where user_id = :1 and role_id = :2`
		sys_rdbms_098 = `update sys_sec_user set continue_error_cnt = :1 where user_id = :2`
		sys_rdbms_099 = `update sys_sec_user set status_id = 1 where user_id = :1`
		sys_rdbms_100 = `select role_id,res_id from sys_role_resource_relat where role_id = :1`
		sys_rdbms_101 = `select t.theme_id,i.theme_desc,res_id,res_url,res_type,res_bg_color,res_class,group_id,res_img,sort_id from sys_theme_value t inner join sys_theme_info i on t.theme_id = i.theme_id where t.theme_id = :1 order by group_id,sort_id asc`
		sys_rdbms_102 = `select t.user_id,t.theme_id,i.theme_desc from sys_user_theme t inner join sys_theme_info i on t.theme_id = i.theme_id where t.user_id = :1`
	}
}
