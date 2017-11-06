-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: test
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.14-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_domain_info`
--

DROP TABLE IF EXISTS `sys_domain_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_info` (
  `domain_id` varchar(30) NOT NULL,
  `domain_name` varchar(300) NOT NULL,
  `domain_status_id` char(1) NOT NULL,
  `domain_create_date` datetime NOT NULL,
  `domain_maintance_date` datetime DEFAULT NULL,
  `domain_maintance_user` varchar(30) DEFAULT NULL,
  `domain_owner` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`domain_id`),
  KEY `fk_sys_idx_05` (`domain_status_id`),
  CONSTRAINT `fk_sys_idx_05` FOREIGN KEY (`domain_status_id`) REFERENCES `sys_domain_status_attr` (`domain_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='域管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_info`
--

LOCK TABLES `sys_domain_info` WRITE;
/*!40000 ALTER TABLE `sys_domain_info` DISABLE KEYS */;
INSERT INTO `sys_domain_info` VALUES ('demo','演示域','0','2017-04-12 22:01:44','2017-04-24 19:17:00','admin','admin'),('mas','管理会计','0','2017-03-01 10:58:18','2017-04-11 22:08:03','ftpadmin','admin'),('product','生产环境','1','2017-04-12 22:02:00','2017-05-16 11:22:51','admin','admin'),('vertex_root','系统顶级域空间','0','2016-12-26 16:43:19','2017-04-24 18:30:34','admin','admin');
/*!40000 ALTER TABLE `sys_domain_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_domain_share_info`
--

DROP TABLE IF EXISTS `sys_domain_share_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_share_info` (
  `uuid` varchar(66) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `target_domain_id` varchar(30) NOT NULL,
  `Authorization_level` char(1) NOT NULL,
  `create_user` varchar(30) DEFAULT NULL,
  `create_date` date DEFAULT NULL,
  `modify_date` date DEFAULT NULL,
  `modify_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_domain_share_info_01_idx` (`domain_id`),
  CONSTRAINT `fk_sys_domain_share_info_01` FOREIGN KEY (`domain_id`) REFERENCES `sys_domain_info` (`domain_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_share_info`
--

LOCK TABLES `sys_domain_share_info` WRITE;
/*!40000 ALTER TABLE `sys_domain_share_info` DISABLE KEYS */;
INSERT INTO `sys_domain_share_info` VALUES ('40ffac77-1a72-11e7-9d82-a0c58951c8d5','mas','devops_product','2','demo','2017-04-06','2017-04-24','admin'),('662dc075-1f88-11e7-9677-a0c58951c8d5','vertex_root','324354325','2','admin','2017-04-12','2017-04-12','admin'),('9d5d6dd4-1f88-11e7-9677-a0c58951c8d5','product','vertex_root','2','admin','2017-04-12','2017-04-12','admin'),('a47eeddc-1f88-11e7-9677-a0c58951c8d5','product','demo','1','admin','2017-04-12','2017-04-12','admin'),('a6d2f828-1f88-11e7-9677-a0c58951c8d5','product','devops_product','1','admin','2017-04-12','2017-04-12','admin');
/*!40000 ALTER TABLE `sys_domain_share_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_domain_status_attr`
--

DROP TABLE IF EXISTS `sys_domain_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_domain_status_attr` (
  `domain_status_id` char(1) NOT NULL,
  `domain_status_name` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`domain_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_domain_status_attr`
--

LOCK TABLES `sys_domain_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_domain_status_attr` DISABLE KEYS */;
INSERT INTO `sys_domain_status_attr` VALUES ('0','正常'),('1','锁定'),('2','失效');
/*!40000 ALTER TABLE `sys_domain_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_handle_logs`
--

DROP TABLE IF EXISTS `sys_handle_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_handle_logs` (
  `uuid` varchar(60) NOT NULL,
  `user_id` varchar(30) DEFAULT NULL,
  `handle_time` datetime DEFAULT NULL,
  `client_ip` varchar(30) DEFAULT NULL,
  `status_code` varchar(10) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  `data` varchar(3000) DEFAULT NULL,
  `domain_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_handle_logs`
--

LOCK TABLES `sys_handle_logs` WRITE;
/*!40000 ALTER TABLE `sys_handle_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_handle_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_index_page`
--

DROP TABLE IF EXISTS `sys_index_page`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_index_page` (
  `theme_id` varchar(30) NOT NULL,
  `res_url` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`theme_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_index_page`
--

LOCK TABLES `sys_index_page` WRITE;
/*!40000 ALTER TABLE `sys_index_page` DISABLE KEYS */;
INSERT INTO `sys_index_page` VALUES ('1001','./views/hauth/theme/default/index.tpl'),('1002','./views/hauth/theme/blue/index.tpl'),('1003','./views/hauth/theme/apple/index.tpl'),('1004','./views/hauth/theme/cyan/index.tpl');
/*!40000 ALTER TABLE `sys_index_page` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_org_info`
--

DROP TABLE IF EXISTS `sys_org_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_org_info` (
  `org_unit_id` varchar(66) NOT NULL,
  `org_unit_desc` varchar(300) NOT NULL,
  `up_org_id` varchar(66) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `create_date` date NOT NULL,
  `maintance_date` date NOT NULL,
  `create_user` varchar(30) NOT NULL,
  `maintance_user` varchar(30) NOT NULL,
  `code_number` varchar(66) NOT NULL,
  PRIMARY KEY (`org_unit_id`),
  KEY `pk_sys_org_info_03_idx` (`domain_id`),
  CONSTRAINT `fk_sys_org_info_01` FOREIGN KEY (`domain_id`) REFERENCES `sys_domain_info` (`domain_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_org_info`
--

LOCK TABLES `sys_org_info` WRITE;
/*!40000 ALTER TABLE `sys_org_info` DISABLE KEYS */;
INSERT INTO `sys_org_info` VALUES ('mas_join_234fda','攀枝花市分行','mas_join_5233454','mas','2017-03-14','2017-04-20','admin','admin','234fda'),('mas_join_34124','工商银行','root_vertex_system','mas','2017-03-01','2017-03-01','admin','admin','34124'),('mas_join_45246543','武汉市分行','mas_join_512345423','mas','2017-03-01','2017-04-24','admin','demo','45246543'),('mas_join_4542346','孝感市分行','mas_join_512345423','mas','2017-03-01','2017-04-21','admin','admin','4542346'),('mas_join_512345423','湖北省分行','mas_join_34124','mas','2017-03-01','2017-04-05','admin','demo','512345423'),('mas_join_5233454','四川省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','5233454'),('mas_join_aefd','欧洲分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','aefd'),('mas_join_fdafdg','贵州省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','fdafdg'),('mas_join_fdaga','重庆市分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','fdaga'),('mas_join_fdagqe','宁夏省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','fdagqe'),('mas_join_fdasfd','上海市分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','fdasfd'),('mas_join_fdsagd','泸州市分行','mas_join_5233454','mas','2017-03-14','2017-03-14','admin','admin','fdsagd'),('mas_join_feqhda','海南省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','feqhda'),('mas_join_ffadg','安徽省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','ffadg'),('mas_join_fgasdbc','台湾省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','fgasdbc'),('mas_join_fgasdf','成都市分行','mas_join_5233454','mas','2017-03-14','2017-03-14','admin','admin','fgasdf'),('mas_join_fgdasdf','南充市分行','mas_join_5233454','mas','2017-03-14','2017-03-14','admin','admin','fgdasdf'),('mas_join_fhadf','香港特别行政区分行','mas_join_34124','mas','2017-03-14','2017-04-24','admin','admin','fhadf'),('mas_join_gasdh3','雅安市分行','mas_join_5233454','mas','2017-03-14','2017-03-14','admin','admin','gasdh3'),('mas_join_reqggfdas','江西省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','reqggfdas'),('mas_join_rqreg','北京市分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','rqreg'),('mas_join_trwt','湖南省分行','mas_join_34124','mas','2017-03-14','2017-03-14','admin','admin','trwt'),('vertex_root_join_vertex_root','系统管理组','root_vertex_system','vertex_root','2016-01-01','2017-04-20','sys','admin','vertex_root');
/*!40000 ALTER TABLE `sys_org_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_info`
--

DROP TABLE IF EXISTS `sys_resource_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_info` (
  `res_id` varchar(30) NOT NULL,
  `res_name` varchar(300) DEFAULT NULL,
  `res_attr` char(1) DEFAULT NULL,
  `res_up_id` varchar(30) DEFAULT NULL,
  `res_type` char(1) DEFAULT NULL,
  `sys_flag` char(1) DEFAULT NULL,
  PRIMARY KEY (`res_id`),
  KEY `fk_sys_idx_13` (`res_type`),
  KEY `fk_sys_idx_14` (`res_attr`),
  CONSTRAINT `fk_sys_idx_13` FOREIGN KEY (`res_type`) REFERENCES `sys_resource_type_attr` (`res_type`),
  CONSTRAINT `fk_sys_idx_14` FOREIGN KEY (`res_attr`) REFERENCES `sys_resource_info_attr` (`res_attr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_info`
--

LOCK TABLES `sys_resource_info` WRITE;
/*!40000 ALTER TABLE `sys_resource_info` DISABLE KEYS */;
INSERT INTO `sys_resource_info` VALUES ('0100000000','系统管理','0','-1','0','0'),('0101000000','系统审计','0','0100000000','4','0'),('0101010000','操作查询','1','0101000000','1','0'),('0101010100','查看操作日志权限','1','0101010000','2',NULL),('0101010200','下载操作日志按钮','1','0101010000','2',NULL),('0101010300','搜索日志信息按钮','1','0101010000','2',NULL),('0103000000','资源管理','0','0100000000','4','0'),('0103010000','菜单','1','0103000000','1','0'),('0103010100','查询资源信息','1','0103010000','2',NULL),('0103010200','新增资源信息按钮','1','0103010000','2',NULL),('0103010300','编辑资源信息按钮','1','0103010000','2',NULL),('0103010400','删除资源信息按钮','1','0103010000','2',NULL),('01030104001','删除资源信息按钮','1','0101010000','2',NULL),('0103010500','配置主题信息按钮','1','0103010000','2',NULL),('0103020000','组织','1','0103000000','1','0'),('0103020100','查询组织架构信息','1','0103020000','2',NULL),('0103020200','新增组织架构信息按钮','1','0103020000','2',NULL),('0103020300','更新组织架构信息按钮','1','0103020000','2',NULL),('0103020400','删除组织架构信息按钮','1','0103020000','2',NULL),('0103020500','导出组织架构信息按钮','1','0103020000','2',NULL),('0103030100','查询共享域信息','1','0104010200','2',NULL),('0103030200','新增共享域信息按钮','1','0104010200','2',NULL),('0103030300','删除共享域信息按钮','1','0104010200','2',NULL),('0103030400','更新共享域信息按钮','1','0104010200','2',NULL),('0104010000','域定义','1','0103000000','1','0'),('0104010100','查询域信息','1','0104010000','2',NULL),('0104010200','共享域管理','1','0104010000','2',NULL),('0104010300','编辑域信息按钮','1','0104010000','2',NULL),('0104010400','删除域信息按钮','1','0104010000','2',NULL),('0104010500','新增域信息按钮','1','0104010000','2',NULL),('0105000000','用户与安全管理','0','0100000000','4','0'),('0105010000','用户','1','0105000000','1','0'),('0105010100','查询用户信息','1','0105010000','2',NULL),('0105010200','新增用户信息按钮','1','0105010000','2',NULL),('0105010300','编辑用户信息按钮','1','0105010000','2',NULL),('0105010400','删除用户信息按钮','1','0105010000','2',NULL),('0105010500','修改用户密码按钮','1','0105010000','2',NULL),('0105010600','修改用户状态按钮','1','0105010000','2',NULL),('0105020000','角色','1','0105000000','1','0'),('0105020100','查询角色信息','1','0105020000','2',NULL),('0105020200','新增角色信息按钮','1','0105020000','2',NULL),('0105020300','更新角色信息按钮','1','0105020000','2',NULL),('0105020400','删除角色信息按钮','1','0105020000','2',NULL),('0105020500','角色资源管理','1','0105020000','2',NULL),('0105020510','查询角色资源信息','1','0105020500','2',NULL),('0105020520','修改角色资源信息','1','0105020500','2',NULL),('0105040000','授权','1','0105000000','1','0'),('0105040100','授予权限按钮','1','0105040000','2',NULL),('0105040200','移除权限','1','0105040000','2',NULL),('0200000000','成本分摊','0','-1','0',NULL),('0201000000','维度信息管理','0','0200000000','4',NULL),('0201010000','责任中心','1','0201000000','1',NULL),('0201030000','成本类别','1','0201000000','1',NULL),('0201040000','动因信息','1','0201000000','1',NULL),('0201060000','成本池信息','1','0201000000','1',NULL),('0202000000','规则定义管理','0','0200000000','4',NULL),('0202010000','静态规则配置','1','0202000000','1',NULL),('0202020000','分摊规则','1','0202000000','1',NULL),('0202040000','规则组配置','1','0202000000','1',NULL),('0203000000','批次综合管理','0','0200000000','4',NULL),('0203010000','批次管理','1','0203000000','1',NULL),('0203020000','批次历史信息','1','0203000000','1',NULL),('0203040000','费用查询','1','0203000000','1',NULL),('0203050000','动因查询','1','0203000000','1',NULL),('0300000000','内部资金转移定价','0','-1','0',NULL),('0301000000','曲线与规则','0','0300000000','4',NULL),('0301010000','曲线定义','1','0301000000','1',NULL),('0301020000','曲线管理','1','0301000000','1',NULL),('0301050000','定价规则','1','0301000000','1',NULL),('0302000000','调节项管理','0','0300000000','4',NULL),('0302010000','内生性调节项','1','0302000000','1',NULL),('0302020000','政策性调节项','1','0302000000','1',NULL),('0302030000','过滤器配置管理','1','0302000000','1',NULL),('0303000000','批次管理','0','0300000000','4',NULL),('0303010000','单笔试算','1','0303000000','1',NULL),('0303020000','批次配置','1','0303000000','1',NULL),('0303030000','批次历史','1','0303000000','1',NULL),('0400000000','公共维度信息','0','-1','0',NULL),('0401000000','条线信息','1','0400000000','1',NULL),('0402000000','产品信息','1','0400000000','1',NULL),('0403000000','科目信息','1','0400000000','1',NULL),('0404000000','币种信息','1','0400000000','1',NULL),('0500000000','ETL调度','0','-1','0',NULL),('0501000000','调度参数配置','0','0500000000','4',NULL),('0501010000','任务参数定义','1','0501000000','1',NULL),('0501020000','调度核心参数管理','1','0501000000','1',NULL),('0502000000','任务与任务组配置','0','0500000000','4',NULL),('0502010000','任务定义','1','0502000000','1',NULL),('0502020000','任务组定义','1','0502000000','1',NULL),('0503000000','批次配置管理','0','0500000000','4',NULL),('0503010000','批次定义','1','0503000000','1',NULL),('0503020000','批次监控','1','0503000000','1',NULL),('1100000000','系统帮助','0','-1','0',NULL),('1101000000','系统管理帮助','0','1100000000','4',NULL),('1101010000','系统维护帮助信息','1','1101000000','1',NULL),('1101020000','API文档','1','1101000000','1',NULL),('1102000000','管理会计帮助文档','0','1100000000','4',NULL),('1103000000','公共信息帮助','0','1100000000','4',NULL);
/*!40000 ALTER TABLE `sys_resource_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_info_attr`
--

DROP TABLE IF EXISTS `sys_resource_info_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_info_attr` (
  `res_attr` char(1) NOT NULL,
  `res_attr_desc` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`res_attr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_info_attr`
--

LOCK TABLES `sys_resource_info_attr` WRITE;
/*!40000 ALTER TABLE `sys_resource_info_attr` DISABLE KEYS */;
INSERT INTO `sys_resource_info_attr` VALUES ('0','目录'),('1','叶子');
/*!40000 ALTER TABLE `sys_resource_info_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource_type_attr`
--

DROP TABLE IF EXISTS `sys_resource_type_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource_type_attr` (
  `res_type` char(1) NOT NULL,
  `res_type_desc` varchar(90) DEFAULT NULL,
  PRIMARY KEY (`res_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource_type_attr`
--

LOCK TABLES `sys_resource_type_attr` WRITE;
/*!40000 ALTER TABLE `sys_resource_type_attr` DISABLE KEYS */;
INSERT INTO `sys_resource_type_attr` VALUES ('0','首页菜单'),('1','子系统菜单'),('2','功能按钮'),('4','虚拟节点');
/*!40000 ALTER TABLE `sys_resource_type_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_info`
--

DROP TABLE IF EXISTS `sys_role_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_info` (
  `role_id` varchar(66) NOT NULL,
  `role_name` varchar(300) NOT NULL,
  `role_owner` varchar(30) NOT NULL,
  `role_create_date` datetime NOT NULL,
  `role_status_id` char(1) NOT NULL,
  `domain_id` varchar(30) NOT NULL,
  `role_maintance_date` datetime NOT NULL,
  `role_maintance_user` varchar(30) NOT NULL,
  `code_number` varchar(66) NOT NULL,
  PRIMARY KEY (`role_id`),
  KEY `fk_sys_idx_11` (`role_status_id`),
  CONSTRAINT `fk_sys_idx_11` FOREIGN KEY (`role_status_id`) REFERENCES `sys_role_status_attr` (`role_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_info`
--

LOCK TABLES `sys_role_info` WRITE;
/*!40000 ALTER TABLE `sys_role_info` DISABLE KEYS */;
INSERT INTO `sys_role_info` VALUES ('devops_product_join_43124','43243','ftpadmin','2017-04-13 00:30:15','0','devops_product','2017-04-13 00:30:15','ftpadmin','43124'),('devops_product_join_454235','543254','ftpadmin','2017-04-13 00:31:47','0','devops_product','2017-04-13 00:31:47','ftpadmin','454235'),('devops_product_join_ftpadmin','FTP管理员角色','admin','2017-03-21 09:43:36','0','devops_product','2017-03-21 09:43:36','admin','ftpadmin'),('mas_join_cademo','成本分摊演示角色','admin','2017-03-07 10:36:45','0','mas','2017-04-20 23:11:32','admin','cademo'),('mas_join_ftpdemo','内部资金转移定价演示角色','admin','2017-03-07 10:36:59','0','mas','2017-03-07 10:36:59','admin','ftpdemo'),('mas_join_masadmin','管理会计管理员','admin','2017-03-14 14:44:34','0','mas','2017-04-22 21:03:20','admin','masadmin'),('vertex_root_join_sysadmin','超级管理员','admin','2016-01-01 00:00:00','0','vertex_root','2016-12-16 00:00:00','admin','sysadmin');
/*!40000 ALTER TABLE `sys_role_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_resource_relat`
--

DROP TABLE IF EXISTS `sys_role_resource_relat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_resource_relat` (
  `uuid` varchar(60) NOT NULL DEFAULT 'uuid()',
  `role_id` varchar(66) DEFAULT NULL,
  `res_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_idx_06` (`res_id`),
  KEY `fk_sys_role_res_01_idx` (`role_id`),
  CONSTRAINT `fk_sys_idx_06` FOREIGN KEY (`res_id`) REFERENCES `sys_resource_info` (`res_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_role_res_01` FOREIGN KEY (`role_id`) REFERENCES `sys_role_info` (`role_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_resource_relat`
--

LOCK TABLES `sys_role_resource_relat` WRITE;
/*!40000 ALTER TABLE `sys_role_resource_relat` DISABLE KEYS */;
INSERT INTO `sys_role_resource_relat` VALUES ('00716df3-07ed-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010600'),('02d6cb28-16e1-11e7-95e0-a0c58951c8d5','mas_join_masadmin','0105040000'),('02d74d86-16e1-11e7-95e0-a0c58951c8d5','mas_join_masadmin','0105040100'),('02d7d7f5-16e1-11e7-95e0-a0c58951c8d5','mas_join_masadmin','0105040200'),('0574d053-07e7-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103020300'),('0a7043a9-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0501010000'),('0a706464-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0502010000'),('0a7078f1-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0502020000'),('0a708f98-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0503010000'),('0a70a2f6-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0501000000'),('0a70ba07-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0502000000'),('0a70d529-467a-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0503000000'),('0ba023b2-4667-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0500000000'),('0f65406b-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0201000000'),('0f655305-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0201040000'),('0f656609-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0203000000'),('0f657dda-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0201030000'),('0f65938e-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0203020000'),('0f65a7da-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0203010000'),('0f65d3c9-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0202000000'),('0f671952-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0202010000'),('0f672d27-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0202020000'),('0f6753eb-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0202040000'),('0f676552-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0203040000'),('0f678912-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0200000000'),('0f679a9f-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0201010000'),('0f67bbf4-02df-11e7-9b60-a0c58951c8d5','mas_join_cademo','0201060000'),('0f931a5a-07f2-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105040100'),('0fed7044-024a-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0301000000'),('15498bd1-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0302000000'),('15499deb-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0303000000'),('1549b2c0-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0301020000'),('1549c489-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0302030000'),('1549da33-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0303010000'),('1549ebe7-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0303020000'),('1549ff00-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0301000000'),('154a0c8d-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0302010000'),('154a1a9e-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0303030000'),('154a2a7c-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0300000000'),('154a62a2-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0302020000'),('154a7233-02df-11e7-9b60-a0c58951c8d5','mas_join_ftpdemo','0301050000'),('17994440-024c-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0303030000'),('1bdeaba6-07e9-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103010100'),('1bf28a08-07e7-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103020400'),('1c3118cc-07e2-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103030400'),('1c7f66c1-4672-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0502010000'),('2372c034-46b1-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0503020000'),('25167037-07f2-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105040200'),('32cfc9e5-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0401000000'),('32cfe510-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0402000000'),('32cff514-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0403000000'),('32d00969-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0404000000'),('32d0a0f2-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0400000000'),('33bb66bb-07e9-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103010200'),('3b92fdf5-4672-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0502020000'),('3d23d85e-07e7-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103020500'),('43ad40d2-07f1-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020510'),('4704352b-0acc-11e7-998e-a0c58951c8d5','mas_join_masadmin','1100000000'),('470450e2-0acc-11e7-998e-a0c58951c8d5','mas_join_masadmin','1101000000'),('4704667c-0acc-11e7-998e-a0c58951c8d5','mas_join_masadmin','1102000000'),('47047a55-0acc-11e7-998e-a0c58951c8d5','mas_join_masadmin','1103000000'),('47048c2b-0acc-11e7-998e-a0c58951c8d5','mas_join_masadmin','1101010000'),('48463b39-07e9-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103010300'),('48fb522e-04a4-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0301010000'),('53c399c4-024c-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0302030000'),('55a149ee-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020500'),('55a16810-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020300'),('55a17bc3-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020400'),('55a18b54-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010200'),('55a199c3-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103030300'),('55a1b0d0-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020520'),('55a1c1e1-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010000'),('55a1da99-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010400'),('55a1ecf2-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010600'),('55a3cd2a-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105000000'),('55a42994-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010500'),('55a48f77-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010000'),('55a4c0d9-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010000'),('55a4efa6-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105040000'),('55a51f7f-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0101010200'),('55a566b2-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020100'),('55a58c3f-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020400'),('55a5abc3-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0100000000'),('55a5c961-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103000000'),('55a5ddd9-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103030100'),('55a5f73b-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103030400'),('55a61bb2-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020500'),('55a640b7-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105040100'),('55a65ed0-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020200'),('55a67332-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010300'),('55a684f2-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010500'),('55a6cb2e-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010300'),('55a711cc-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010500'),('55a7297f-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020100'),('55a74032-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0101000000'),('55a757d0-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0101010000'),('55a76915-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020200'),('55a77b15-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020300'),('55a78c3f-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010400'),('55a8088c-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010200'),('55a87773-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010100'),('55a8a7c8-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010100'),('55a8bd08-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105040200'),('55a8eaf7-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103030200'),('55a900c4-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020510'),('55a912e6-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103020000'),('55a925c8-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105020000'),('55a938ea-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0103010300'),('55a94aa1-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0104010400'),('55a95d48-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010100'),('55a98588-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0105010200'),('55a9998c-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0101010100'),('55a9af08-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0101010300'),('5a587e71-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0400000000'),('5a588e25-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0401000000'),('5a589e29-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0402000000'),('5a5a35ba-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0403000000'),('5a5a4743-1f9d-11e7-9677-a0c58951c8d5','devops_product_join_43124','0404000000'),('5a7db1f7-07f1-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020520'),('5c60bc08-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0301050000'),('5cdef223-4671-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0501000000'),('60700eba-1fed-11e7-9677-a0c58951c8d5','vertex_root_join_sysadmin','1101000000'),('607033cf-1fed-11e7-9677-a0c58951c8d5','vertex_root_join_sysadmin','1100000000'),('6070454b-1fed-11e7-9677-a0c58951c8d5','vertex_root_join_sysadmin','1101010000'),('6402f992-4672-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0503010000'),('68ebf2c8-1fed-11e7-9677-a0c58951c8d5','vertex_root_join_sysadmin','1103000000'),('692c628f-1c0a-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0101010100'),('6a935ea9-1fed-11e7-9677-a0c58951c8d5','vertex_root_join_sysadmin','1102000000'),('6bb7e04d-07e9-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103010400'),('6c7f6d2a-250a-11e7-9c7e-a0c58951c8d5','vertex_root_join_sysadmin','01030104001'),('72939327-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0302000000'),('7c3618ec-4671-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0502000000'),('7d73294c-07ec-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010100'),('8009b52c-0ba4-11e7-9649-a0c58951c8d5','mas_join_cademo','0203050000'),('8024c16b-07d8-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0104010300'),('824c1f28-04a3-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0400000000'),('83794268-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0302010000'),('8857ba73-4671-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0503000000'),('8ca4f732-07e5-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0101010200'),('8dc4fada-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0201060000'),('8dc56ba3-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0203040000'),('8dc57fe7-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0203050000'),('8dc59452-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0202000000'),('8dc5a6f0-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0201010000'),('8dc5bba7-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0200000000'),('8dc5d11a-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0201030000'),('8dc5e7da-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0202040000'),('8dc5ffda-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0203020000'),('8dc6176b-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0201000000'),('8dc62d85-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0203000000'),('8dc63ec1-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0201040000'),('8dc653b0-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0202010000'),('8dc669ab-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0202020000'),('8dc68185-5bb3-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','0203010000'),('9466d2dc-07d5-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0104010200'),('970569ee-07d8-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0104010400'),('974d1286-07ec-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010200'),('9e79cb72-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0302020000'),('9f6f310f-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','0400000000'),('9f6f4846-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','0401000000'),('9f6f630f-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','0402000000'),('9f6fadc6-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','0403000000'),('9f6fc475-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','0404000000'),('a0e2a82e-20f8-11e7-966c-a0c58951c8d5','vertex_root_join_sysadmin','1101020000'),('a11cab89-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','1100000000'),('a11cc274-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','1101000000'),('a11cd974-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','1102000000'),('a11cee27-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','1103000000'),('a11cfdc5-1f8a-11e7-9677-a0c58951c8d5','devops_product_join_ftpadmin','1101010000'),('a2658092-07ed-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020100'),('a2a01355-07e5-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0101010300'),('ad3e53ed-07d8-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0104010500'),('ad96ffe8-0992-11e7-952f-a0c58951c8d5','mas_join_masadmin','0101010000'),('ad972957-0992-11e7-952f-a0c58951c8d5','mas_join_masadmin','0101010300'),('ad973d01-0992-11e7-952f-a0c58951c8d5','mas_join_masadmin','0101000000'),('ad974e5b-0992-11e7-952f-a0c58951c8d5','mas_join_masadmin','0101010200'),('af623c20-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0301020000'),('af6254c6-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0302010000'),('af6268c2-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0302030000'),('af627c0a-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0303010000'),('af62b80e-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0300000000'),('af62c935-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0302000000'),('af62da9f-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0303000000'),('af62e857-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0302020000'),('af62f630-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0303020000'),('af64a874-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0303030000'),('af64be06-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0301000000'),('af64d2b0-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0301010000'),('af64e4f9-1aca-11e7-9d82-a0c58951c8d5','devops_product_join_ftpadmin','0301050000'),('b096b467-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0303000000'),('b257854d-04a3-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0401000000'),('b5801636-07ec-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010300'),('b687b293-024a-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0301020000'),('b6ca0b31-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010300'),('b6ca200b-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010400'),('b6ca36e4-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010500'),('b6ca480f-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010600'),('b6ca5c0b-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010000'),('b6cab506-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103030200'),('b6cac00f-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103000000'),('b6cad202-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020000'),('b6cae5b5-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020400'),('b6caf864-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010200'),('b6cc6dcb-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010100'),('b6cc8746-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020400'),('b6cc9c46-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105000000'),('b6ccae31-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020200'),('b6ccbf4f-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010100'),('b6ccd5ad-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010200'),('b6ccf9f1-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010300'),('b6cd0a06-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010300'),('b6cd1c82-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020510'),('b6cd3017-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010400'),('b6cd66f5-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010000'),('b6cd7506-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010100'),('b6cd8439-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020500'),('b6cd9375-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020100'),('b6cda1f9-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020500'),('b6cdb0d7-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103030100'),('b6cdccfe-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010000'),('b6cddc28-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103010200'),('b6cdea17-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020100'),('b6cdfb93-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010500'),('b6ce08d7-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103030400'),('b6ce14f1-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0104010400'),('b6ce228f-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105010500'),('b6ce2ded-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020200'),('b6ce39b5-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020300'),('b6ce49b5-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0100000000'),('b6ce568f-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0105020000'),('b6ce7217-0881-11e7-952f-a0c58951c8d5','mas_join_masadmin','0103020300'),('b8df3b71-07e9-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103010500'),('ba1baad1-0249-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0300000000'),('bd267b0e-07ed-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020200'),('becdf6e3-0eb9-11e7-9612-a0c58951c8d5','vertex_root_join_sysadmin','0101010100'),('c1177dbf-07e1-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103030100'),('c3baf059-07ee-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020500'),('c8650311-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0303010000'),('c988dc67-07ec-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010400'),('ca968c8b-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0101010100'),('ca96ae0b-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0101010200'),('ca96c387-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0101010300'),('ca96d85d-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','01030104001'),('ca96ecc7-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0101010000'),('ca970110-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0101000000'),('ca9713fa-289c-11e7-9c7e-a0c58951c8d5','mas_join_cademo','0100000000'),('cb09f0fd-0eb9-11e7-9612-a0c58951c8d5','mas_join_masadmin','0101010100'),('cb4b16fb-04a3-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0402000000'),('d347b0d3-4671-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0501010000'),('d517d48d-07ed-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020300'),('d6746779-0ba4-11e7-9649-a0c58951c8d5','mas_join_ftpdemo','0301010000'),('d8fd37ed-07e1-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103030200'),('daae0b92-07e6-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103020100'),('dbaf4cc1-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010200'),('dbaf6401-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010000'),('dbaf77a3-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0101010300'),('dbaf8930-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020300'),('dbaf991b-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020500'),('dbafaae3-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010100'),('dbafbc30-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010200'),('dbafce38-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010500'),('dbafdeca-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020200'),('dbaff192-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020500'),('dbb01efd-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0101010000'),('dbb03370-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105040000'),('dbb0424a-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020400'),('dbb0533d-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010300'),('dbb063b8-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010100'),('dbb07456-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010300'),('dbb0868e-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020100'),('dbb098db-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010000'),('dbb0b6bd-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103030200'),('dbb0c8d6-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103030400'),('dbb0d7e7-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020000'),('dbb0e45f-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010100'),('dbb0f052-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010400'),('dbb0ff4a-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020400'),('dbb10c30-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103030300'),('dbb1182c-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020000'),('dbb14505-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010200'),('dbb265ac-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010300'),('dbb27678-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010500'),('dbb2a54e-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020300'),('dbb2bf78-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020520'),('dbb2dbb4-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103030100'),('dbb2e9c5-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0100000000'),('dbb2f83d-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105000000'),('dbb30885-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010000'),('dbb322ca-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020100'),('dbb33adf-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010400'),('dbb3539b-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105010600'),('dbb36bf8-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105040200'),('dbb38238-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0104010400'),('dbb399f4-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105040100'),('dbb3b16c-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0101000000'),('dbb3c901-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103000000'),('dbb3ddce-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0101010200'),('dbb3f538-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103010500'),('dbb40745-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0103020200'),('dbb41aa7-0dd7-11e7-9612-a0c58951c8d5','devops_product_join_ftpadmin','0105020510'),('e4e93b85-46b1-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0501020000'),('e61931f7-04a3-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0403000000'),('ea23a4e6-07ed-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105020400'),('ec5e6b47-07ec-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0105010500'),('ecfe2317-024b-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0303020000'),('ee768238-07e6-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103020200'),('f0766b0d-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0100000000'),('f07680fd-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0101000000'),('f076a4d5-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0103000000'),('f076b2d1-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0103010000'),('f076c09b-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0103020000'),('f076e3ca-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0104010000'),('f076efb4-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0105000000'),('f076fb82-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0105010000'),('f077074b-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0105020000'),('f0771e6b-c597-11e6-9b11-d4bed967cdf1','vertex_root_join_sysadmin','0101010000'),('f0771e6b-c597-11e6-9b11-d4bed967cdff','vertex_root_join_sysadmin','0105040000'),('f0cd283e-4666-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0500000000'),('f2e86103-07d2-11e7-95d9-a0c58951c8d5','vertex_root_join_sysadmin','0104010100'),('f44f6baa-46b0-11e7-9beb-a0c58951c8d5','vertex_root_join_sysadmin','0503020000'),('f6a653e9-04a3-11e7-9b60-a0c58951c8d5','vertex_root_join_sysadmin','0404000000'),('f82d2048-46b1-11e7-9beb-a0c58951c8d5','mas_join_masadmin','0501020000'),('fb9787a0-07e1-11e7-952f-a0c58951c8d5','vertex_root_join_sysadmin','0103030300');
/*!40000 ALTER TABLE `sys_role_resource_relat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_status_attr`
--

DROP TABLE IF EXISTS `sys_role_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_status_attr` (
  `role_status_id` char(1) NOT NULL,
  `role_status_desc` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`role_status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_status_attr`
--

LOCK TABLES `sys_role_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_role_status_attr` DISABLE KEYS */;
INSERT INTO `sys_role_status_attr` VALUES ('0','正常'),('1','锁定'),('2','失效');
/*!40000 ALTER TABLE `sys_role_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_user_relation`
--

DROP TABLE IF EXISTS `sys_role_user_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_user_relation` (
  `uuid` varchar(60) NOT NULL,
  `role_id` varchar(66) DEFAULT NULL,
  `user_id` varchar(30) DEFAULT NULL,
  `maintance_date` date DEFAULT NULL,
  `maintance_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`uuid`),
  KEY `fk_sys_idx_03` (`user_id`),
  KEY `fk_sys_role_user_01_idx` (`role_id`),
  CONSTRAINT `fk_sys_idx_03` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_role_user_01` FOREIGN KEY (`role_id`) REFERENCES `sys_role_info` (`role_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_user_relation`
--

LOCK TABLES `sys_role_user_relation` WRITE;
/*!40000 ALTER TABLE `sys_role_user_relation` DISABLE KEYS */;
INSERT INTO `sys_role_user_relation` VALUES ('19890228hzwy23','vertex_root_join_sysadmin','admin','2000-01-01','hzwy23'),('74f3ba5a-5bb4-11e7-9d30-a0c58951c8d5','vertex_root_join_sysadmin','431243','2017-06-28','admin'),('caadmin_join_mas_join_cademo','mas_join_cademo','caadmin','2017-06-28','admin'),('demo_join_mas_join_cademo','mas_join_cademo','demo','2017-06-28','admin');
/*!40000 ALTER TABLE `sys_role_user_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_sec_user`
--

DROP TABLE IF EXISTS `sys_sec_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_sec_user` (
  `user_id` varchar(30) NOT NULL,
  `user_passwd` varchar(30) DEFAULT NULL,
  `status_id` char(1) DEFAULT NULL,
  `continue_error_cnt` int(11) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_idx_02` (`status_id`),
  CONSTRAINT `fk_sys_idx_01` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_sys_idx_02` FOREIGN KEY (`status_id`) REFERENCES `sys_user_status_attr` (`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_sec_user`
--

LOCK TABLES `sys_sec_user` WRITE;
/*!40000 ALTER TABLE `sys_sec_user` DISABLE KEYS */;
INSERT INTO `sys_sec_user` VALUES ('431243','CguSVgQY2Df4LxG0UT/xwA==','0',NULL),('admin','rVbaiQ3XuCj8aCnhIL1KAA==','0',0),('caadmin','CguSVgQY2Df4LxG0UT/xwA==','0',0),('demo','CguSVgQY2Df4LxG0UT/xwA==','0',0);
/*!40000 ALTER TABLE `sys_sec_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_theme_info`
--

DROP TABLE IF EXISTS `sys_theme_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_theme_info` (
  `theme_id` varchar(30) NOT NULL,
  `theme_desc` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`theme_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_theme_info`
--

LOCK TABLES `sys_theme_info` WRITE;
/*!40000 ALTER TABLE `sys_theme_info` DISABLE KEYS */;
INSERT INTO `sys_theme_info` VALUES ('1001','绿色主题'),('1002','深蓝主题'),('1003','粉色主题'),('1004','青色主题');
/*!40000 ALTER TABLE `sys_theme_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_theme_value`
--

DROP TABLE IF EXISTS `sys_theme_value`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_theme_value` (
  `uuid` varchar(60) NOT NULL,
  `theme_id` varchar(30) DEFAULT NULL,
  `res_id` varchar(30) DEFAULT NULL,
  `res_url` varchar(120) DEFAULT NULL,
  `res_type` varchar(5) DEFAULT NULL,
  `res_bg_color` varchar(30) DEFAULT NULL,
  `res_class` varchar(90) DEFAULT NULL,
  `group_id` char(1) DEFAULT NULL,
  `res_img` varchar(200) DEFAULT NULL,
  `sort_id` decimal(10,0) DEFAULT NULL,
  KEY `pk_sys_theme_value_01` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_theme_value`
--

LOCK TABLES `sys_theme_value` WRITE;
/*!40000 ALTER TABLE `sys_theme_value` DISABLE KEYS */;
INSERT INTO `sys_theme_value` VALUES ('1001-0101010000','1001','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1),('1001-0103010000','1001','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',3),('1001-0104010000','1001','0104010000','/v1/auth/domain/page','0','#0099CC','tile tile-wide','1','/static/images/hauth/domain.png',1),('1001-0105010000','1001','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',1),('1001-0105020000','1001','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',2),('1001-0100000000','1001','0100000000','./views/hauth/theme/default/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1),('1001-0105040000','1001','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4),('1001-0103020000','1001','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','1','/static/images/hauth/org.png',2),('54786c62-0246-11e7-9b60-a0c58951c8d5','1001','0200000000','./apps/mas/ca/views/cyan/ca.tpl','0','#666699','tile tile-wide','2','/static/images/hauth/grant.png',1),('fb944b0a-0246-11e7-9b60-a0c58951c8d5','1001','0201010000','/v1/ca/responsibility/page','0','#6fc07c','tile tile-wide','1','/static/images/ca_icon/org_info.png',1),('5046d07a-0247-11e7-9b60-a0c58951c8d5','1001','0201030000','/v1/ca/cost/direction/page','0','#FF9999','tile','1','/static/images/ca_icon/cost_direction.png',3),('7929aa2b-0247-11e7-9b60-a0c58951c8d5','1001','0201040000','/v1/ca/driver/page','0','#b4d39e','tile','1','/static/images/ca_icon/driver_info.png',4),('c93c4e93-0247-11e7-9b60-a0c58951c8d5','1001','0201060000','/v1/ca/cost/page','0','#e4d690','tile tile-wide','1','/static/images/ca_icon/cost_pool.png',6),('f02a3b32-0247-11e7-9b60-a0c58951c8d5','1001','0202010000','/v1/ca/static/radio/page','0','#009966','tile tile-wide','2','/static/images/ca_icon/static_rules.png',1),('19c73fba-0248-11e7-9b60-a0c58951c8d5','1001','0202020000','/v1/ca/amart/rules/page','0','#3399CC','tile tile-wide','2','/static/images/ca_icon/amart_rules.png',2),('66e55e26-0248-11e7-9b60-a0c58951c8d5','1001','0202040000','/v1/ca/amart/group/page','0','#99CC33','tile tile-wide','2','/static/images/ca_icon/group_rules.png',4),('8a180b66-0248-11e7-9b60-a0c58951c8d5','1001','0203010000','/v1/ca/dispatch/page','0','#666699','tile tile-wide','3','/static/images/ca_icon/dispatch_manage.png',1),('a831ec58-0248-11e7-9b60-a0c58951c8d5','1001','0203020000','/v1/ca/dispatch/history/page','0','#339999','tile tile-wide','3','/static/images/ca_icon/dispatch_history.png',2),('dd815000-0248-11e7-9b60-a0c58951c8d5','1001','0203040000','/v1/ca/cost/manage/page','0','#FFCC33','tile','3','/static/images/ca_icon/cost_query.png',3),('ba1a252f-0249-11e7-9b60-a0c58951c8d5','1001','0300000000','./apps/mas/ftp/views/theme/default/ftp.tpl','0','#009999','tile tile-wide','2','/static/images/hauth/ftp.png',2),('948f67dc-024a-11e7-9b60-a0c58951c8d5','1001','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1),('b687a0e9-024a-11e7-9b60-a0c58951c8d5','1001','0301020000','/v1/ftp/curve/manage/page','0','#336699','tile','1','/static/images/hauth/curve_manage.png',2),('5c60abdd-024b-11e7-9b60-a0c58951c8d5','1001','0301050000','/v1/ftp/rules/manage/page','0','#99CC33','tile tile-wide','1','/static/images/hauth/ftp_rules.png',3),('83792fdb-024b-11e7-9b60-a0c58951c8d5','1001','0302010000','/v1/ftp/adjust/inner/page','0','#0099CC','tile','2','/static/images/hauth/ftp_inner_adjust.png',2),('9e79b725-024b-11e7-9b60-a0c58951c8d5','1001','0302020000','/v1/ftp/adjust/outer/page','0','#CC6600','tile','2','/static/images/hauth/ftp_outer_adjust.png',3),('c864e93c-024b-11e7-9b60-a0c58951c8d5','1001','0303010000','/v1/ftp/single/calc/page','0','#FF6666','tile tile-wide','3','/static/images/hauth/ftp_single_calc.png',1),('ecfe0b20-024b-11e7-9b60-a0c58951c8d5','1001','0303020000','/v1/ftp/dispatch/manage/page','0','#009933','tile','3','/static/images/hauth/ftp_dispatch.png',2),('1797ac80-024c-11e7-9b60-a0c58951c8d5','1001','0303030000','/v1/ftp/dispatch/history/page','0','#009999','tile','3','/static/images/hauth/ftp_dispatch_history.png',3),('53c3813f-024c-11e7-9b60-a0c58951c8d5','1001','0302030000','/v1/ftp/filter/define/page','0','#FFCC33','tile tile-wide','2','/static/images/hauth/ftp_filter.png',1),('624b90c0-0278-11e7-9b60-a0c58951c8d5','1002','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1),('824c0d97-04a3-11e7-9b60-a0c58951c8d5','1001','0400000000','./apps/mas/common/views/green/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1),('b2561d1e-04a3-11e7-9b60-a0c58951c8d5','1001','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1),('cb4afcc4-04a3-11e7-9b60-a0c58951c8d5','1001','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2),('e6191fef-04a3-11e7-9b60-a0c58951c8d5','1001','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1),('f6a6448b-04a3-11e7-9b60-a0c58951c8d5','1001','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1),('48fb4303-04a4-11e7-9b60-a0c58951c8d5','1001','0301010000','/v1/ftp/curve/define/page','0','#666699','tile','1','/static/images/hauth/curve_define.png',1),('f2e81083-07d2-11e7-95d9-a0c58951c8d5','1001','0104010100','/v1/auth/domain/get','0','','','','',0),('946658e9-07d5-11e7-952f-a0c58951c8d5','1001','0104010200','/v1/auth/domain/share/page','0','','','','',0),('8024ac09-07d8-11e7-952f-a0c58951c8d5','1001','0104010300','/v1/auth/domain/update','0','','','','',0),('9705437b-07d8-11e7-952f-a0c58951c8d5','1001','0104010400','/v1/auth/domain/delete','0','','','','',0),('ad3e295c-07d8-11e7-952f-a0c58951c8d5','1001','0104010500','/v1/auth/domain/post','0','','','','',0),('c1174621-07e1-11e7-952f-a0c58951c8d5','1001','0103030100','/v1/auth/domain/share/get','0','','','','',0),('d8fccbcb-07e1-11e7-952f-a0c58951c8d5','1001','0103030200','/v1/auth/domain/share/post','0','','','','',0),('fb975107-07e1-11e7-952f-a0c58951c8d5','1001','0103030300','/v1/auth/domain/share/delete','0','','','','',0),('1c30f988-07e2-11e7-952f-a0c58951c8d5','1001','0103030400','/v1/auth/domain/share/put','0','','','','',0),('8ca386d8-07e5-11e7-952f-a0c58951c8d5','1001','0101010200','/v1/auth/handle/logs/download','0','','','','',0),('a29fba3f-07e5-11e7-952f-a0c58951c8d5','1001','0101010300','/v1/auth/handle/logs/search','0','','','','',0),('daadf91b-07e6-11e7-952f-a0c58951c8d5','1001','0103020100','/v1/auth/resource/org/get','0','','','','',0),('ee765e9a-07e6-11e7-952f-a0c58951c8d5','1001','0103020200','/v1/auth/resource/org/insert','0','','','','',0),('0574add7-07e7-11e7-952f-a0c58951c8d5','1001','0103020300','/v1/auth/resource/org/update','0','','','','',0),('1bf270aa-07e7-11e7-952f-a0c58951c8d5','1001','0103020400','/v1/auth/resource/org/delete','0','','','','',0),('3d237ba7-07e7-11e7-952f-a0c58951c8d5','1001','0103020500','/v1/auth/resource/org/download','0','','','','',0),('1bde8991-07e9-11e7-952f-a0c58951c8d5','1001','0103010100','/v1/auth/resource/get','0','','','','',0),('33b9cb0c-07e9-11e7-952f-a0c58951c8d5','1001','0103010200','/v1/auth/resource/post','0','','','','',0),('48460086-07e9-11e7-952f-a0c58951c8d5','1001','0103010300','/v1/auth/resource/update','0','','','','',0),('6bb7b2c8-07e9-11e7-952f-a0c58951c8d5','1001','0103010400','/v1/auth/resource/delete','0','','','','',0),('b8df0cd7-07e9-11e7-952f-a0c58951c8d5','1001','0103010500','/v1/auth/resource/config/theme','0','','','','',0),('7d73058c-07ec-11e7-952f-a0c58951c8d5','1001','0105010100','/v1/auth/user/get','0','','','','',0),('974ce1fd-07ec-11e7-952f-a0c58951c8d5','1001','0105010200','/v1/auth/user/post','0','','','','',0),('b58002f6-07ec-11e7-952f-a0c58951c8d5','1001','0105010300','/v1/auth/user/put','0','','','','',0),('c988bb89-07ec-11e7-952f-a0c58951c8d5','1001','0105010400','/v1/auth/user/delete','0','','','','',0),('ec5cb33a-07ec-11e7-952f-a0c58951c8d5','1001','0105010500','/v1/auth/user/modify/passwd','0','','','','',0),('00714873-07ed-11e7-952f-a0c58951c8d5','1001','0105010600','/v1/auth/user/modify/status','0','','','','',0),('a265597d-07ed-11e7-952f-a0c58951c8d5','1001','0105020100','/v1/auth/role/get','0','','','','',0),('bd264fd7-07ed-11e7-952f-a0c58951c8d5','1001','0105020200','/v1/auth/role/post','0','','','','',0),('d517aab8-07ed-11e7-952f-a0c58951c8d5','1001','0105020300','/v1/auth/role/update','0','','','','',0),('ea237b6a-07ed-11e7-952f-a0c58951c8d5','1001','0105020400','/v1/auth/role/delete','0','','','','',0),('c3bad47b-07ee-11e7-952f-a0c58951c8d5','1001','0105020500','/v1/auth/role/resource/details','0','','','','',0),('43ad2a9a-07f1-11e7-952f-a0c58951c8d5','1001','0105020510','/v1/auth/role/resource/get','0','','','','',0),('5a7d8dbf-07f1-11e7-952f-a0c58951c8d5','1001','0105020520','/v1/auth/role/resource/rights','0','','','','',0),('0f9303e2-07f2-11e7-952f-a0c58951c8d5','1001','0105040100','/v1/auth/user/roles/auth','0','','','','',0),('25165700-07f2-11e7-952f-a0c58951c8d5','1001','0105040200','/v1/auth/user/roles/revoke','0','','','','',0),('0e9aec3f-094c-11e7-952f-a0c58951c8d5','1001','0203050000','/v1/ca/driver/manage/page','0','#CC6600','tile','3','/static/images/ca_icon/driver_query.png',4),('f87a9123-0991-11e7-952f-a0c58951c8d5','1001','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1),('991641c3-0d55-11e7-964b-a0c58951c8d5','1004','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1),('99164f5c-0d55-11e7-964b-a0c58951c8d5','1004','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',3),('9916502d-0d55-11e7-964b-a0c58951c8d5','1004','0104010000','/v1/auth/domain/page','0','#0099CC','tile tile-wide','1','/static/images/hauth/domain.png',1),('991650a9-0d55-11e7-964b-a0c58951c8d5','1004','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',1),('9916512d-0d55-11e7-964b-a0c58951c8d5','1004','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',2),('9916519c-0d55-11e7-964b-a0c58951c8d5','1004','0100000000','./views/hauth/theme/cyan/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1),('99165203-0d55-11e7-964b-a0c58951c8d5','1004','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4),('9916525c-0d55-11e7-964b-a0c58951c8d5','1004','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','1','/static/images/hauth/org.png',2),('991652b2-0d55-11e7-964b-a0c58951c8d5','1004','0200000000','./apps/mas/ca/views/green/ca.tpl','0','#666699','tile tile-wide','2','/static/images/hauth/grant.png',1),('9916534b-0d55-11e7-964b-a0c58951c8d5','1004','0201010000','/v1/ca/responsibility/page','0','#6fc07c','tile tile-wide','1','/static/images/ca_icon/org_info.png',1),('9916545c-0d55-11e7-964b-a0c58951c8d5','1004','0201030000','/v1/ca/cost/direction/page','0','#FF9999','tile','1','/static/images/ca_icon/cost_direction.png',3),('991654be-0d55-11e7-964b-a0c58951c8d5','1004','0201040000','/v1/ca/driver/page','0','#b4d39e','tile','1','/static/images/ca_icon/driver_info.png',4),('991657d4-0d55-11e7-964b-a0c58951c8d5','1004','0201060000','/v1/ca/cost/page','0','#e4d690','tile tile-wide','1','/static/images/ca_icon/cost_pool.png',6),('9916933a-0d55-11e7-964b-a0c58951c8d5','1004','0202010000','/v1/ca/static/radio/page','0','#92cdd2','tile tile-wide','2','/static/images/ca_icon/static_rules.png',1),('9917f369-0d55-11e7-964b-a0c58951c8d5','1004','0202020000','/v1/ca/amart/rules/page','0','#58c0b3','tile tile-wide','2','/static/images/ca_icon/amart_rules.png',2),('9917f42d-0d55-11e7-964b-a0c58951c8d5','1004','0202040000','/v1/ca/amart/group/page','0','#ded1b0','tile tile-wide','2','/static/images/ca_icon/group_rules.png',4),('9917f48b-0d55-11e7-964b-a0c58951c8d5','1004','0203010000','/v1/ca/dispatch/page','0','#666699','tile tile-wide','3','/static/images/ca_icon/dispatch_manage.png',1),('9917f4cb-0d55-11e7-964b-a0c58951c8d5','1004','0203020000','/v1/ca/dispatch/history/page','0','#339999','tile tile-wide','3','/static/images/ca_icon/dispatch_history.png',2),('9917f598-0d55-11e7-964b-a0c58951c8d5','1004','0203040000','/v1/ca/cost/manage/page','0','#FFCC33','tile','3','/static/images/ca_icon/cost_query.png',3),('9917f676-0d55-11e7-964b-a0c58951c8d5','1004','0300000000','./apps/mas/ftp/views/theme/cyan/ftp.tpl','0','#009999','tile tile-wide','2','/static/images/hauth/ftp.png',2),('9917f6e5-0d55-11e7-964b-a0c58951c8d5','1004','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1),('9917f743-0d55-11e7-964b-a0c58951c8d5','1004','0301020000','/v1/ftp/curve/manage/page','0','#336699','tile','1','/static/images/hauth/curve_manage.png',2),('9917f7ba-0d55-11e7-964b-a0c58951c8d5','1004','0301050000','/v1/ftp/rules/manage/page','0','#99CC33','tile tile-wide','1','/static/images/hauth/ftp_rules.png',3),('9917f818-0d55-11e7-964b-a0c58951c8d5','1004','0302010000','/v1/ftp/adjust/inner/page','0','#0099CC','tile','2','/static/images/hauth/ftp_inner_adjust.png',2),('9917f869-0d55-11e7-964b-a0c58951c8d5','1004','0302020000','/v1/ftp/adjust/outer/page','0','#CC6600','tile','2','/static/images/hauth/ftp_outer_adjust.png',3),('9917f8b6-0d55-11e7-964b-a0c58951c8d5','1004','0303010000','/v1/ftp/single/calc/page','0','#FF6666','tile tile-wide','3','/static/images/hauth/ftp_single_calc.png',1),('99180aad-0d55-11e7-964b-a0c58951c8d5','1004','0303020000','/v1/ftp/dispatch/manage/page','0','#009933','tile','3','/static/images/hauth/ftp_dispatch.png',2),('99180b3a-0d55-11e7-964b-a0c58951c8d5','1004','0303030000','/v1/ftp/dispatch/history/page','0','#009999','tile','3','/static/images/hauth/ftp_dispatch_history.png',3),('99180b7a-0d55-11e7-964b-a0c58951c8d5','1004','0302030000','/v1/ftp/filter/define/page','0','#FFCC33','tile tile-wide','2','/static/images/hauth/ftp_filter.png',1),('99180bfa-0d55-11e7-964b-a0c58951c8d5','1004','0400000000','./apps/mas/common/views/cyan/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1),('99180c36-0d55-11e7-964b-a0c58951c8d5','1004','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1),('99180c72-0d55-11e7-964b-a0c58951c8d5','1004','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2),('99180ca9-0d55-11e7-964b-a0c58951c8d5','1004','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1),('99180ced-0d55-11e7-964b-a0c58951c8d5','1004','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1),('99180d2d-0d55-11e7-964b-a0c58951c8d5','1004','0301010000','/v1/ftp/curve/define/page','0','#666699','tile','1','/static/images/hauth/curve_define.png',1),('99180d65-0d55-11e7-964b-a0c58951c8d5','1004','0104010100','/v1/auth/domain/get','0','','','','',0),('99180da1-0d55-11e7-964b-a0c58951c8d5','1004','0104010200','/v1/auth/domain/share/page','0','','','','',0),('99180ddc-0d55-11e7-964b-a0c58951c8d5','1004','0104010300','/v1/auth/domain/update','0','','','','',0),('99180e14-0d55-11e7-964b-a0c58951c8d5','1004','0104010400','/v1/auth/domain/delete','0','','','','',0),('99180e4f-0d55-11e7-964b-a0c58951c8d5','1004','0104010500','/v1/auth/domain/post','0','','','','',0),('99180e87-0d55-11e7-964b-a0c58951c8d5','1004','0103030100','/v1/auth/domain/share/get','0','','','','',0),('99180ec3-0d55-11e7-964b-a0c58951c8d5','1004','0103030200','/v1/auth/domain/share/post','0','','','','',0),('99180efa-0d55-11e7-964b-a0c58951c8d5','1004','0103030300','/v1/auth/domain/share/delete','0','','','','',0),('99180f32-0d55-11e7-964b-a0c58951c8d5','1004','0103030400','/v1/auth/domain/share/put','0','','','','',0),('99180fa1-0d55-11e7-964b-a0c58951c8d5','1004','0101010200','/v1/auth/handle/logs/download','0','','','','',0),('99180fdc-0d55-11e7-964b-a0c58951c8d5','1004','0101010300','/v1/auth/handle/logs/search','0','','','','',0),('99181014-0d55-11e7-964b-a0c58951c8d5','1004','0103020100','/v1/auth/resource/org/get','0','','','','',0),('9918104b-0d55-11e7-964b-a0c58951c8d5','1004','0103020200','/v1/auth/resource/org/insert','0','','','','',0),('99181087-0d55-11e7-964b-a0c58951c8d5','1004','0103020300','/v1/auth/resource/org/update','0','','','','',0),('991810be-0d55-11e7-964b-a0c58951c8d5','1004','0103020400','/v1/auth/resource/org/delete','0','','','','',0),('991810fe-0d55-11e7-964b-a0c58951c8d5','1004','0103020500','/v1/auth/resource/org/download','0','','','','',0),('9918113a-0d55-11e7-964b-a0c58951c8d5','1004','0103010100','/v1/auth/resource/get','0','','','','',0),('99181176-0d55-11e7-964b-a0c58951c8d5','1004','0103010200','/v1/auth/resource/post','0','','','','',0),('991811ad-0d55-11e7-964b-a0c58951c8d5','1004','0103010300','/v1/auth/resource/update','0','','','','',0),('991811e1-0d55-11e7-964b-a0c58951c8d5','1004','0103010400','/v1/auth/resource/delete','0','','','','',0),('99181218-0d55-11e7-964b-a0c58951c8d5','1004','0103010500','/v1/auth/resource/config/theme','0','','','','',0),('9918124f-0d55-11e7-964b-a0c58951c8d5','1004','0105010100','/v1/auth/user/get','0','','','','',0),('9918128b-0d55-11e7-964b-a0c58951c8d5','1004','0105010200','/v1/auth/user/post','0','','','','',0),('991812c3-0d55-11e7-964b-a0c58951c8d5','1004','0105010300','/v1/auth/user/put','0','','','','',0),('991812fa-0d55-11e7-964b-a0c58951c8d5','1004','0105010400','/v1/auth/user/delete','0','','','','',0),('99181332-0d55-11e7-964b-a0c58951c8d5','1004','0105010500','/v1/auth/user/modify/passwd','0','','','','',0),('99181365-0d55-11e7-964b-a0c58951c8d5','1004','0105010600','/v1/auth/user/modify/status','0','','','','',0),('9918139c-0d55-11e7-964b-a0c58951c8d5','1004','0105020100','/v1/auth/role/get','0','','','','',0),('991813d4-0d55-11e7-964b-a0c58951c8d5','1004','0105020200','/v1/auth/role/post','0','','','','',0),('9918140b-0d55-11e7-964b-a0c58951c8d5','1004','0105020300','/v1/auth/role/update','0','','','','',0),('99181443-0d55-11e7-964b-a0c58951c8d5','1004','0105020400','/v1/auth/role/delete','0','','','','',0),('99181476-0d55-11e7-964b-a0c58951c8d5','1004','0105020500','/v1/auth/role/resource/details','0','','','','',0),('991814ad-0d55-11e7-964b-a0c58951c8d5','1004','0105020510','/v1/auth/role/resource/get','0','','','','',0),('991814f2-0d55-11e7-964b-a0c58951c8d5','1004','0105020520','/v1/auth/role/resource/rights','0','','','','',0),('9918152d-0d55-11e7-964b-a0c58951c8d5','1004','0105040100','/v1/auth/user/roles/auth','0','','','','',0),('99181569-0d55-11e7-964b-a0c58951c8d5','1004','0105040200','/v1/auth/user/roles/revoke','0','','','','',0),('991815a1-0d55-11e7-964b-a0c58951c8d5','1004','0203050000','/v1/ca/driver/manage/page','0','#CC6600','tile','3','/static/images/ca_icon/driver_query.png',4),('991815e1-0d55-11e7-964b-a0c58951c8d5','1004','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1),('becde5db-0eb9-11e7-9612-a0c58951c8d5','1001','0101010100','/v1/auth/handle/logs','0','','','','',0),('8e2d2ae7-1c0a-11e7-9d82-a0c58951c8d5','1004','0101010100','/v1/auth/handle/logs','0','','tile tile-large','','',0),('a0e208f2-20f8-11e7-966c-a0c58951c8d5','1001','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1),('b3f18e0b-20f8-11e7-966c-a0c58951c8d5','1004','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1),('6c7f5772-250a-11e7-9c7e-a0c58951c8d5','1001','01030104001','/v1/auth/resource/org/page','','','','','',0),('8a8c3203-2b27-11e7-9c7e-a0c58951c8d5','1002','0200000000','./apps/mas/ca/views/blue/ca.tpl','0','#666699','tile tile-wide','2','/static/images/hauth/grant.png',1),('9b081aec-2b27-11e7-9c7e-a0c58951c8d5','1002','0100000000','./views/hauth/theme/blue/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1),('a343cbfc-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010100','/v1/auth/handle/logs','0','','tile','','',0),('a65d91b0-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010200','/v1/auth/handle/logs/download','0','','tile','','',0),('a8854ec0-2b27-11e7-9c7e-a0c58951c8d5','1002','0101010300','/v1/auth/handle/logs/search','0','','tile','','',0),('aabbbd36-2b27-11e7-9c7e-a0c58951c8d5','1002','01030104001','/v1/auth/resource/org/page','0','','tile','','',0),('af0e054c-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',3),('b1314131-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010100','/v1/auth/resource/get','0','','tile','','',0),('b3c7c6a6-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010200','/v1/auth/resource/post','0','','tile','','',0),('b6372ff3-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010300','/v1/auth/resource/update','0','','tile','','',0),('b8d3d1c1-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010400','/v1/auth/resource/delete','0','','tile','','',0),('bb9fc76f-2b27-11e7-9c7e-a0c58951c8d5','1002','0103010500','/v1/auth/resource/config/theme','0','','tile','','',0),('bea9df22-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020100','/v1/auth/resource/org/get','0','','tile','','',0),('c15e0f8b-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020200','/v1/auth/resource/org/insert','0','','tile','','',0),('c37806f8-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020300','/v1/auth/resource/org/update','0','','tile','','',0),('c59c3303-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020400','/v1/auth/resource/org/delete','0','','tile','','',0),('c77c6ed0-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020500','/v1/auth/resource/org/download','0','','tile','','',0),('cc8891d2-2b27-11e7-9c7e-a0c58951c8d5','1002','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','1','/static/images/hauth/org.png',2),('d1f01d28-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010000','/v1/auth/domain/page','0','#0099CC','tile tile-wide','1','/static/images/hauth/domain.png',1),('d4bfa83c-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010100','/v1/auth/domain/get','0','','tile','','',0),('d767f63e-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010200','/v1/auth/domain/share/page','0','','tile','','',0),('da84a5e1-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030100','/v1/auth/domain/share/get','0','','tile','','',0),('dc65642a-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030200','/v1/auth/domain/share/post','0','','tile','','',0),('de8f9fcb-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030300','/v1/auth/domain/share/delete','0','','tile','','',0),('e0a10dc4-2b27-11e7-9c7e-a0c58951c8d5','1002','0103030400','/v1/auth/domain/share/put','0','','tile','','',0),('e2e782c4-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010300','/v1/auth/domain/update','0','','tile','','',0),('e4e17463-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010400','/v1/auth/domain/delete','0','','tile','','',0),('e777d2c2-2b27-11e7-9c7e-a0c58951c8d5','1002','0104010500','/v1/auth/domain/post','0','','tile','','',0),('eb13f0e9-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010100','/v1/auth/user/get','0','','tile','','',0),('ed148f2a-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010200','/v1/auth/user/post','0','','tile','','',0),('ef613f0c-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010300','/v1/auth/user/put','0','','tile','','',0),('f19af335-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010400','/v1/auth/user/delete','0','','tile','','',0),('f3959708-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010500','/v1/auth/user/modify/passwd','0','','tile','','',0),('f5a0999f-2b27-11e7-9c7e-a0c58951c8d5','1002','0105010600','/v1/auth/user/modify/status','0','','tile','','',0),('f94b4a93-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020100','/v1/auth/role/get','0','','tile','','',0),('fbcd8b0b-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',2),('fdb44348-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020200','/v1/auth/role/post','0','','tile','','',0),('ff9f6773-2b27-11e7-9c7e-a0c58951c8d5','1002','0105020300','/v1/auth/role/update','0','','tile','','',0),('0287ee48-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020400','/v1/auth/role/delete','0','','tile','','',0),('052dc4ac-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020500','/v1/auth/role/resource/details','0','','tile','','',0),('0875a5f3-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020510','/v1/auth/role/resource/get','0','','tile','','',0),('0a964ef9-2b28-11e7-9c7e-a0c58951c8d5','1002','0105020520','/v1/auth/role/resource/rights','0','','tile','','',0),('0e4ca28b-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4),('107e273d-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040100','/v1/auth/user/roles/auth','0','','tile','','',0),('12cd5409-2b28-11e7-9c7e-a0c58951c8d5','1002','0105040200','/v1/auth/user/roles/revoke','0','','tile','','',0),('51a5bff2-2b28-11e7-9c7e-a0c58951c8d5','1002','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',1),('e464ee50-2b2a-11e7-9c7e-a0c58951c8d5','1002','0201010000','/v1/ca/responsibility/page','0','#FF6666','tile tile-wide','1','/static/images/ca_icon/org_info.png',1),('e6c1fb99-2b2a-11e7-9c7e-a0c58951c8d5','1002','0201030000','/v1/ca/cost/direction/page','0','#FF9999','tile','1','/static/images/ca_icon/cost_direction.png',3),('e8d836db-2b2a-11e7-9c7e-a0c58951c8d5','1002','0201040000','/v1/ca/driver/page','0','#CCCC33','tile','1','/static/images/ca_icon/driver_info.png',4),('eac811b6-2b2a-11e7-9c7e-a0c58951c8d5','1002','0201060000','/v1/ca/cost/page','0','#CC6633','tile tile-wide','1','/static/images/ca_icon/cost_pool.png',6),('edc39c65-2b2a-11e7-9c7e-a0c58951c8d5','1002','0202010000','/v1/ca/static/radio/page','0','#009966','tile tile-wide','2','/static/images/ca_icon/static_rules.png',1),('efab4483-2b2a-11e7-9c7e-a0c58951c8d5','1002','0202020000','/v1/ca/amart/rules/page','0','#3399CC','tile tile-wide','2','/static/images/ca_icon/amart_rules.png',2),('f1b7d81a-2b2a-11e7-9c7e-a0c58951c8d5','1002','0202040000','/v1/ca/amart/group/page','0','#99CC33','tile tile-wide','2','/static/images/ca_icon/group_rules.png',4),('f3a437e0-2b2a-11e7-9c7e-a0c58951c8d5','1002','0203010000','/v1/ca/dispatch/page','0','#666699','tile tile-wide','3','/static/images/ca_icon/dispatch_manage.png',1),('f6004f8a-2b2a-11e7-9c7e-a0c58951c8d5','1002','0203020000','/v1/ca/dispatch/history/page','0','#339999','tile tile-wide','3','/static/images/ca_icon/dispatch_history.png',2),('fa02435a-2b2a-11e7-9c7e-a0c58951c8d5','1002','0203040000','/v1/ca/cost/manage/page','0','#FFCC33','tile','3','/static/images/ca_icon/cost_query.png',3),('fc7a0545-2b2a-11e7-9c7e-a0c58951c8d5','1002','0203050000','/v1/ca/driver/manage/page','0','#CC6600','tile','3','/static/images/ca_icon/driver_query.png',4),('015376ca-2b2b-11e7-9c7e-a0c58951c8d5','1002','0400000000','./apps/mas/common/views/blue/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1),('07ab049a-2b2b-11e7-9c7e-a0c58951c8d5','1002','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1),('098dd130-2b2b-11e7-9c7e-a0c58951c8d5','1002','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2),('0bbabbfb-2b2b-11e7-9c7e-a0c58951c8d5','1002','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1),('0db2afab-2b2b-11e7-9c7e-a0c58951c8d5','1002','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1),('4fd8fdcf-2b42-11e7-9c7e-a0c58951c8d5','1002','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1),('5dcfdfc0-2b42-11e7-9c7e-a0c58951c8d5','1002','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1),('60c6e788-2b42-11e7-9c7e-a0c58951c8d5','1002','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1),('d4605f79-2b43-11e7-9c7e-a0c58951c8d5','1003','0100000000','./views/hauth/theme/apple/sysconfig.tpl','0','#FF6600','tile tile-wide','1','/static/images/hauth/system.png',1),('d7867a7a-2b43-11e7-9c7e-a0c58951c8d5','1003','0101010000','/v1/auth/HandleLogsPage','0','#336699','tile tile-large','3','/static/images/hauth/logs_shen.png',1),('dd972a84-2b43-11e7-9c7e-a0c58951c8d5','1003','0103010000','/v1/auth/resource/page','0','#666699','tile','1','/static/images/hauth/menus.png',3),('e007d284-2b43-11e7-9c7e-a0c58951c8d5','1003','0103020000','/v1/auth/resource/org/page','0','#FF6666','tile','1','/static/images/hauth/org.png',2),('e224205c-2b43-11e7-9c7e-a0c58951c8d5','1003','0104010000','/v1/auth/domain/page','0','#0099CC','tile tile-wide','1','/static/images/hauth/domain.png',1),('e4ac3710-2b43-11e7-9c7e-a0c58951c8d5','1003','0105010000','/v1/auth/user/page','0','#CC6600','tile tile-wide','2','/static/images/hauth/user_manager.png',1),('e716b0a1-2b43-11e7-9c7e-a0c58951c8d5','1003','0105020000','/v1/auth/role/page','0','#FFCC33','tile','2','/static/images/hauth/role_manager.png',2),('ea4b0eda-2b43-11e7-9c7e-a0c58951c8d5','1003','0105040000','/v1/auth/batch/page','0','#339999','tile','2','/static/images/hauth/grant.png',4),('09f47f4e-2b44-11e7-9c7e-a0c58951c8d5','1003','0200000000','./apps/mas/ca/views/apple/ca.tpl','0','#666699','tile tile-wide','2','/static/images/hauth/grant.png',1),('14d34e0f-2b44-11e7-9c7e-a0c58951c8d5','1003','0201010000','/v1/ca/responsibility/page','0','#6fc07c','tile tile-wide','1','/static/images/ca_icon/org_info.png',1),('186b1431-2b44-11e7-9c7e-a0c58951c8d5','1003','0201030000','/v1/ca/cost/direction/page','0','#FF9999','tile','1','/static/images/ca_icon/cost_direction.png',3),('1ac87e71-2b44-11e7-9c7e-a0c58951c8d5','1003','0201040000','/v1/ca/driver/page','0','#b4d39e','tile','1','/static/images/ca_icon/driver_info.png',4),('1d205fdc-2b44-11e7-9c7e-a0c58951c8d5','1003','0201060000','/v1/ca/cost/page','0','#e4d690','tile tile-wide','1','/static/images/ca_icon/cost_pool.png',6),('215c357e-2b44-11e7-9c7e-a0c58951c8d5','1003','0203010000','/v1/ca/dispatch/page','0','#666699','tile tile-wide','3','/static/images/ca_icon/dispatch_manage.png',1),('239acba3-2b44-11e7-9c7e-a0c58951c8d5','1003','0203020000','/v1/ca/dispatch/history/page','0','#339999','tile tile-wide','3','/static/images/ca_icon/dispatch_history.png',2),('2824ed08-2b44-11e7-9c7e-a0c58951c8d5','1003','0203040000','/v1/ca/cost/manage/page','0','#FFCC33','tile','3','/static/images/ca_icon/cost_query.png',3),('2a58bc0b-2b44-11e7-9c7e-a0c58951c8d5','1003','0203050000','/v1/ca/driver/manage/page','0','#CC6600','tile','3','/static/images/ca_icon/driver_query.png',4),('3403b3b7-2b44-11e7-9c7e-a0c58951c8d5','1003','1100000000','./views/help/default/syshelp.tpl','0','#0099CC','tile tile-wide','3','/static/images/hauth/help.png',1),('37ceac85-2b44-11e7-9c7e-a0c58951c8d5','1003','1101010000','/v1/help/system/help','0','#339999','tile tile-wide','1','/static/images/hauth/sys_help.png',1),('3a0e741e-2b44-11e7-9c7e-a0c58951c8d5','1003','1101020000','/v1/auth/swagger/page','1','#339999','tile tile-wide','2','/static/images/hauth/api.png',1),('ba15af88-2b44-11e7-9c7e-a0c58951c8d5','1003','0400000000','./apps/mas/common/views/apple/dimension.tpl','0','#FFCC33','tile tile-wide','3','/static/images/hauth/system.png',1),('eb87e6f6-2b44-11e7-9c7e-a0c58951c8d5','1002','0300000000','./apps/mas/ftp/views/theme/blue/ftp.tpl','0','#009999','tile tile-wide','2','/static/images/hauth/ftp.png',2),('f4f2f6ed-2b44-11e7-9c7e-a0c58951c8d5','1003','0300000000','./apps/mas/ftp/views/theme/apple/ftp.tpl','0','#009999','tile tile-wide','2','/static/images/hauth/ftp.png',2),('ffed3495-2b44-11e7-9c7e-a0c58951c8d5','1002','0301010000','/v1/ftp/curve/define/page','0','#666699','tile','1','/static/images/hauth/curve_define.png',1),('04322974-2b45-11e7-9c7e-a0c58951c8d5','1003','0301010000','/v1/ftp/curve/define/page','0','#666699','tile','1','/static/images/hauth/curve_define.png',1),('080a1969-2b45-11e7-9c7e-a0c58951c8d5','1002','0301020000','/v1/ftp/curve/manage/page','0','#336699','tile','1','/static/images/hauth/curve_manage.png',2),('09f345ff-2b45-11e7-9c7e-a0c58951c8d5','1003','0301020000','/v1/ftp/curve/manage/page','0','#336699','tile','1','/static/images/hauth/curve_manage.png',2),('0d1637eb-2b45-11e7-9c7e-a0c58951c8d5','1002','0301050000','/v1/ftp/rules/manage/page','0','#99CC33','tile tile-wide','1','/static/images/hauth/ftp_rules.png',3),('0ea14dd4-2b45-11e7-9c7e-a0c58951c8d5','1003','0301050000','/v1/ftp/rules/manage/page','0','#99CC33','tile tile-wide','1','/static/images/hauth/ftp_rules.png',3),('10f8e018-2b45-11e7-9c7e-a0c58951c8d5','1002','0302010000','/v1/ftp/adjust/inner/page','0','#0099CC','tile','2','/static/images/hauth/ftp_inner_adjust.png',2),('12d82536-2b45-11e7-9c7e-a0c58951c8d5','1003','0302010000','/v1/ftp/adjust/inner/page','0','#0099CC','tile','2','/static/images/hauth/ftp_inner_adjust.png',2),('14c02da6-2b45-11e7-9c7e-a0c58951c8d5','1002','0302020000','/v1/ftp/adjust/outer/page','0','#CC6600','tile','2','/static/images/hauth/ftp_outer_adjust.png',3),('165ca720-2b45-11e7-9c7e-a0c58951c8d5','1003','0302020000','/v1/ftp/adjust/outer/page','0','#CC6600','tile','2','/static/images/hauth/ftp_outer_adjust.png',3),('18c97590-2b45-11e7-9c7e-a0c58951c8d5','1002','0302030000','/v1/ftp/filter/define/page','0','#FFCC33','tile tile-wide','2','/static/images/hauth/ftp_filter.png',1),('1aa86a37-2b45-11e7-9c7e-a0c58951c8d5','1003','0302030000','/v1/ftp/filter/define/page','0','#FFCC33','tile tile-wide','2','/static/images/hauth/ftp_filter.png',1),('209d20b2-2b45-11e7-9c7e-a0c58951c8d5','1002','0303010000','/v1/ftp/single/calc/page','0','#FF6666','tile tile-wide','3','/static/images/hauth/ftp_single_calc.png',1),('22a8ba05-2b45-11e7-9c7e-a0c58951c8d5','1003','0303010000','/v1/ftp/single/calc/page','0','#FF6666','tile tile-wide','3','/static/images/hauth/ftp_single_calc.png',1),('258f0649-2b45-11e7-9c7e-a0c58951c8d5','1002','0303020000','/v1/ftp/dispatch/manage/page','0','#009933','tile','3','/static/images/hauth/ftp_dispatch.png',2),('270af058-2b45-11e7-9c7e-a0c58951c8d5','1003','0303020000','/v1/ftp/dispatch/manage/page','0','#009933','tile','3','/static/images/hauth/ftp_dispatch.png',2),('29cadde3-2b45-11e7-9c7e-a0c58951c8d5','1002','0303030000','/v1/ftp/dispatch/history/page','0','#009999','tile','3','/static/images/hauth/ftp_dispatch_history.png',3),('2ba922e4-2b45-11e7-9c7e-a0c58951c8d5','1003','0303030000','/v1/ftp/dispatch/history/page','0','#009999','tile','3','/static/images/hauth/ftp_dispatch_history.png',3),('32cb5534-2b45-11e7-9c7e-a0c58951c8d5','1003','0401000000','/v1/common/depart/page','0','#6fc07c','tile tile-wide','1','/static/images/common_icon/department.png',1),('350dd891-2b45-11e7-9c7e-a0c58951c8d5','1003','0402000000','/v1/common/product/page','0','#92cdd2','tile tile-wide','1','/static/images/common_icon/product.png',2),('3e3b64e8-2b45-11e7-9c7e-a0c58951c8d5','1003','0403000000','/v1/common/glaccount/page','0','#ed9f86','tile tile-large','2','/static/images/common_icon/gl_account.png',1),('40813c9f-2b45-11e7-9c7e-a0c58951c8d5','1003','0404000000','/v1/common/isocurrency/page','0','#67accd','tile tile-large','3','/static/images/common_icon/iso_currency.png',1),('f58252ab-2b47-11e7-9c7e-a0c58951c8d5','1003','0202010000','/v1/ca/static/radio/page','0','#009966','tile tile-wide','2','/static/images/ca_icon/static_rules.png',1),('f7d09d3b-2b47-11e7-9c7e-a0c58951c8d5','1003','0202020000','/v1/ca/amart/rules/page','0','#3399CC','tile tile-wide','2','/static/images/ca_icon/amart_rules.png',2),('fa0a85e5-2b47-11e7-9c7e-a0c58951c8d5','1003','0202040000','/v1/ca/amart/group/page','0','#ded1b0','tile tile-wide','2','/static/images/ca_icon/group_rules.png',4),('f0ccdcc7-4666-11e7-9beb-a0c58951c8d5','1001','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-wide','2','/static/images/dispatch_icon/etl.png',1),('fa2728fc-4666-11e7-9beb-a0c58951c8d5','1002','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-wide','2','/static/images/dispatch_icon/etl.png',1),('fc276e32-4666-11e7-9beb-a0c58951c8d5','1003','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-wide','2','/static/images/dispatch_icon/etl.png',1),('fe6968d2-4666-11e7-9beb-a0c58951c8d5','1004','0500000000','./views/dispatch/index.tpl','0','#009966','tile tile-wide','2','/static/images/dispatch_icon/etl.png',1),('d3478cbe-4671-11e7-9beb-a0c58951c8d5','1001','0501010000','/v1/dispatch/argument/page','0','#FF6666','tile tile-wide','1','/static/images/dispatch_icon/arg_define.png',1),('1c7f5512-4672-11e7-9beb-a0c58951c8d5','1001','0502010000','/v1/dispatch/task/page','0','#FF9999','tile tile-wide','2','/static/images/dispatch_icon/task.png',1),('3b92e08b-4672-11e7-9beb-a0c58951c8d5','1001','0502020000','/v1/dispatch/group/page','0','#CC6633','tile','2','/static/images/dispatch_icon/group.png',2),('6402d0f4-4672-11e7-9beb-a0c58951c8d5','1001','0503010000','/v1/dispatch/batch/page','0','#009966','tile tile-wide','3','/static/images/dispatch_icon/batch.png',1),('78c703b2-4672-11e7-9beb-a0c58951c8d5','1002','0501010000','/v1/dispatch/argument/page','0','#FF6666','tile tile-wide','1','/static/images/dispatch_icon/arg_define.png',1),('7ac7fbb6-4672-11e7-9beb-a0c58951c8d5','1003','0501010000','/v1/dispatch/argument/page','0','#FF6666','tile tile-wide','1','/static/images/dispatch_icon/arg_define.png',1),('7c98873a-4672-11e7-9beb-a0c58951c8d5','1004','0501010000','/v1/dispatch/argument/page','0','#FF6666','tile tile-wide','1','/static/images/dispatch_icon/arg_define.png',1),('80454951-4672-11e7-9beb-a0c58951c8d5','1002','0502010000','/v1/dispatch/task/page','0','#FF9999','tile tile-wide','2','/static/images/dispatch_icon/task.png',1),('82027c48-4672-11e7-9beb-a0c58951c8d5','1003','0502010000','/v1/dispatch/task/page','0','#FF9999','tile tile-wide','2','/static/images/dispatch_icon/task.png',1),('83debaa2-4672-11e7-9beb-a0c58951c8d5','1004','0502010000','/v1/dispatch/task/page','0','#FF9999','tile tile-wide','2','/static/images/dispatch_icon/task.png',1),('87220100-4672-11e7-9beb-a0c58951c8d5','1002','0502020000','/v1/dispatch/group/page','0','#CC6633','tile','2','/static/images/dispatch_icon/group.png',2),('88d9c13f-4672-11e7-9beb-a0c58951c8d5','1003','0502020000','/v1/dispatch/group/page','0','#CC6633','tile','2','/static/images/dispatch_icon/group.png',2),('8a806181-4672-11e7-9beb-a0c58951c8d5','1004','0502020000','/v1/dispatch/group/page','0','#CC6633','tile','2','/static/images/dispatch_icon/group.png',2),('8d108700-4672-11e7-9beb-a0c58951c8d5','1002','0503010000','/v1/dispatch/batch/page','0','#009966','tile tile-wide','3','/static/images/dispatch_icon/batch.png',1),('8f2d5f2b-4672-11e7-9beb-a0c58951c8d5','1003','0503010000','/v1/dispatch/batch/page','0','#009966','tile tile-wide','3','/static/images/dispatch_icon/batch.png',1),('9153d196-4672-11e7-9beb-a0c58951c8d5','1004','0503010000','/v1/dispatch/batch/page','0','#009966','tile tile-wide','3','/static/images/dispatch_icon/batch.png',1),('f44f54e6-46b0-11e7-9beb-a0c58951c8d5','1001','0503020000','/v1/dispatch/batch/monitoring/page','0','#0099CC','tile tile-wide','3','/static/images/dispatch_icon/monitoring.png',2),('0ba33d1a-46b1-11e7-9beb-a0c58951c8d5','1002','0503020000','/v1/dispatch/batch/monitoring/page','0','#0099CC','tile tile-wide','3','/static/images/dispatch_icon/monitoring.png',2),('0d5ea3fc-46b1-11e7-9beb-a0c58951c8d5','1003','0503020000','/v1/dispatch/batch/monitoring/page','0','#0099CC','tile tile-wide','3','/static/images/dispatch_icon/monitoring.png',2),('0f1882f2-46b1-11e7-9beb-a0c58951c8d5','1004','0503020000','/v1/dispatch/batch/monitoring/page','0','#0099CC','tile tile-wide','3','/static/images/dispatch_icon/monitoring.png',2),('e4e9263d-46b1-11e7-9beb-a0c58951c8d5','1001','0501020000','/v1/dispatch/system/config/page','0','#FF9966','tile tile-wide','1','/static/images/dispatch_icon/arg_system.png',2),('ed067953-46b1-11e7-9beb-a0c58951c8d5','1002','0501020000','/v1/dispatch/system/config/page','0','#FF9966','tile tile-wide','1','/static/images/dispatch_icon/arg_system.png',2),('eea5c6c0-46b1-11e7-9beb-a0c58951c8d5','1003','0501020000','/v1/dispatch/system/config/page','0','#FF9966','tile tile-wide','1','/static/images/dispatch_icon/arg_system.png',2),('f03b11ac-46b1-11e7-9beb-a0c58951c8d5','1004','0501020000','/v1/dispatch/system/config/page','0','#FF9966','tile tile-wide','1','/static/images/dispatch_icon/arg_system.png',2);
/*!40000 ALTER TABLE `sys_theme_value` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_info`
--

DROP TABLE IF EXISTS `sys_user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_info` (
  `user_id` varchar(30) NOT NULL,
  `user_name` varchar(300) DEFAULT NULL,
  `user_create_date` datetime DEFAULT NULL,
  `user_owner` varchar(30) DEFAULT NULL,
  `user_email` varchar(30) DEFAULT NULL,
  `user_phone` decimal(15,0) DEFAULT NULL,
  `org_unit_id` varchar(66) DEFAULT NULL,
  `user_maintance_date` datetime DEFAULT NULL,
  `user_maintance_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_user_org_idx` (`org_unit_id`),
  CONSTRAINT `fk_sys_user_org` FOREIGN KEY (`org_unit_id`) REFERENCES `sys_org_info` (`org_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_info`
--

LOCK TABLES `sys_user_info` WRITE;
/*!40000 ALTER TABLE `sys_user_info` DISABLE KEYS */;
INSERT INTO `sys_user_info` VALUES ('431243','2143214','2017-06-28 10:51:09','admin','hzwy23@163.com',18986110550,'vertex_root_join_vertex_root','2017-06-28 10:51:09','admin'),('admin','超级管理员','2016-01-01 00:00:00','sys','hzwy23@163.com',18986110550,'vertex_root_join_vertex_root','2017-06-28 10:54:54','admin'),('caadmin','CA业务管理员','2017-03-18 14:32:22','admin','hzwy23@163.com',18986110550,'mas_join_34124','2017-03-18 14:32:22','admin'),('demo','演示账号','2017-03-01 21:21:38','admin','hzwy23@sina.com',18986110551,'mas_join_4542346','2017-04-24 20:53:42','admin');
/*!40000 ALTER TABLE `sys_user_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_status_attr`
--

DROP TABLE IF EXISTS `sys_user_status_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_status_attr` (
  `status_id` char(1) NOT NULL,
  `status_desc` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`status_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_status_attr`
--

LOCK TABLES `sys_user_status_attr` WRITE;
/*!40000 ALTER TABLE `sys_user_status_attr` DISABLE KEYS */;
INSERT INTO `sys_user_status_attr` VALUES ('0','正常'),('1','锁定'),('2','失效');
/*!40000 ALTER TABLE `sys_user_status_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_theme`
--

DROP TABLE IF EXISTS `sys_user_theme`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_theme` (
  `user_id` varchar(30) NOT NULL,
  `theme_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `pk_sys_user_theme_01` FOREIGN KEY (`user_id`) REFERENCES `sys_user_info` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_theme`
--

LOCK TABLES `sys_user_theme` WRITE;
/*!40000 ALTER TABLE `sys_user_theme` DISABLE KEYS */;
INSERT INTO `sys_user_theme` VALUES ('431243','1001'),('admin','1002'),('caadmin','1002'),('demo','1002');
/*!40000 ALTER TABLE `sys_user_theme` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'test'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-06-28 21:10:56
