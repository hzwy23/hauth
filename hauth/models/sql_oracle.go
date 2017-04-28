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
		sys_rdbms_006 = `select count(*) as cnt from sys_theme_value where theme_id = :1 and res_id = :2`
		sys_rdbms_010 = `select user_id,user_passwd,status_id,continue_error_cnt from sys_sec_user where user_id = :1`
		sys_rdbms_011 = `select distinct t2.res_url from sys_user_theme t1 inner join sys_theme_value t2 on t1.theme_id = t2.theme_id where t1.user_id = :1 and t2.res_id = :2 and t2.res_type = '0'`
		sys_rdbms_013 = `select res_type from sys_theme_value where theme_id = '1001' and res_id = :1`
		sys_rdbms_017 = `select t.user_id,t.user_name,a.status_desc,t.user_create_date, t.user_owner,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc, di.domain_id,di.domain_name,t.user_maintance_date,t.user_maintance_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id inner join sys_domain_info di on i.domain_id = di.domain_id where di.domain_id = :1`
		sys_rdbms_022 = `select count(*) from sys_user_info t inner join sys_role_user_relation r on t.user_id = r.user_id inner join sys_role_resource_relat e on r.role_id = e.role_id inner join sys_theme_value v on e.res_id = v.res_id inner join sys_user_theme m on v.theme_id = m.theme_id and t.user_id = m.user_id where t.user_id = :1 and v.res_url = :2`
		sys_rdbms_023 = `select t.user_id,t.user_name,a.status_desc,t.user_create_date, t.user_owner,t.user_email,t.user_phone,i.org_unit_id,i.org_unit_desc,di.domain_id,di.domain_name,t.user_maintance_date,t.user_maintance_user,u.status_id from sys_user_info t inner join sys_sec_user u on t.user_id = u.user_id inner join sys_user_status_attr a on u.status_id = a.status_id inner join sys_org_info i on i.org_unit_id = t.org_unit_id inner join sys_domain_info di on i.domain_id = di.domain_id where t.user_id = :1`
		sys_rdbms_028 = `select  t.code_number,t.role_name,t.role_owner,t.role_create_date,a.role_status_desc,a.role_status_id,t.domain_id,o.domain_name,t.role_maintance_date,t.role_maintance_user,t.role_id from sys_role_info t inner join sys_role_status_attr a on t.role_status_id = a.role_status_id inner join sys_domain_info o on t.domain_id = o.domain_id where t.domain_id = :1`
		sys_rdbms_034 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.domain_create_date  as maintance_date, t.domain_owner as user_id,t.domain_maintance_date,t.domain_maintance_user from sys_domain_info t inner join sys_domain_status_attr s  on t.domain_status_id = s.domain_status_id where exists ( select 1 from sys_domain_share_info i where i.target_domain_id = :1  and t.domain_id = i.domain_id ) or t.domain_id = :2`
		sys_rdbms_041 = `select org_unit_id,org_unit_desc,up_org_id,t.org_status_id,r.org_status_desc,t.domain_id,create_date,maintance_date,create_user,maintance_user,code_number from sys_org_info t inner join sys_org_status_attr r on t.org_status_id = r.org_status_id where t.domain_id = :1`
		sys_rdbms_046 = `select t.role_id,t.role_name,t.code_number from sys_role_info t where ( t.role_owner = :1 or exists ( select 1 from sys_role_user_relation r where r.user_id = :2 and t.role_id = r.role_id ))`
		sys_rdbms_047 = `select t.role_id,t.role_name,t.code_number from sys_role_info t where ( t.role_owner = :1 or exists ( select 1 from sys_role_user_relation r where r.user_id = :2 and t.role_id = r.role_id )) and not exists (select 1 from sys_role_user_relation n where n.user_id = :3 and t.role_id = n.role_id )`
		sys_rdbms_070 = `select t.theme_id,i.theme_desc, res_id,res_url,res_type,res_bg_color,res_class,group_id,res_img,sort_id from sys_theme_value t left join sys_theme_info i on t.theme_id = i.theme_id where t.theme_id = :1 and t.res_id = :2`
		sys_rdbms_071 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc,t.sys_flag from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type`
		sys_rdbms_078 = `select t1.res_url from sys_index_page t1 inner join sys_user_theme t2 on t1.theme_id = t2.theme_id where t2.user_id = :1`
		sys_rdbms_079 = `select distinct domain_id from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id where user_id = :1`
		sys_rdbms_080 = `select o.org_unit_id from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id where user_id = :1`
		sys_rdbms_083 = `select t.uuid,t.target_domain_id,i.domain_name,t.authorization_level,t.create_user,t.create_date,t.modify_user,t.modify_date from sys_domain_share_info t inner join sys_domain_info i on t.target_domain_id = i.domain_id where t.domain_id = :1`
		sys_rdbms_084 = `select t.domain_id as project_id, t.domain_name as project_name, s.domain_status_name  as status_name, t.domain_create_date  as maintance_date, t.domain_owner as user_id,t.domain_maintance_date,t.domain_maintance_user from sys_domain_info t inner join sys_domain_status_attr s  on t.domain_status_id = s.domain_status_id where t.domain_id = :1`
		sys_rdbms_085 = `select t.domain_id as project_id, t.domain_name as project_name from sys_domain_info t where not exists ( select 1 from sys_domain_share_info i where t.domain_id = i.target_domain_id and i.domain_id = :1 )`
		sys_rdbms_089 = `select t.res_id,t.res_name,t.res_attr, a.res_attr_desc,t.res_up_id,t.res_type,r.res_type_desc from sys_resource_info t inner join sys_resource_info_attr a on t.res_attr = a.res_attr inner join sys_resource_type_attr r on t.res_type = r.res_type where res_id = :1`
		sys_rdbms_094 = `select r.user_id, t.role_id,t.code_number,t.role_name from sys_role_info t inner join sys_role_user_relation r on t.role_id = r.role_id where r.user_id = :1`
		sys_rdbms_095 = `select '',t.role_id,t.code_number,t.role_name from sys_user_info i inner join sys_org_info o on i.org_unit_id = o.org_unit_id inner join sys_role_info t on o.domain_id = t.domain_id where i.user_id = :1 and  not exists ( select 1 from sys_role_user_relation r where i.user_id = r.user_id and r.role_id = t.role_id )`
	}
}
