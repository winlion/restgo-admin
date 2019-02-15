# Host: localhost  (Version: 5.5.53-log)
# Date: 2019-02-15 15:26:27
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "config"
#

DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `name` varchar(20) NOT NULL,
  `value` varchar(1024) DEFAULT NULL,
  `label` varchar(40) DEFAULT NULL,
  `format` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

#
# Data for table "config"
#

/*!40000 ALTER TABLE `config` DISABLE KEYS */;
INSERT INTO `config` VALUES ('guestid','3','用户作为游客身份访问默认角色ID','^\\d+$'),('role','1','用户注册默认角色id','^\\d+$');
/*!40000 ALTER TABLE `config` ENABLE KEYS */;

#
# Structure for table "ref_role_res"
#

DROP TABLE IF EXISTS `ref_role_res`;
CREATE TABLE `ref_role_res` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleid` int(11) DEFAULT NULL,
  `resid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

#
# Data for table "ref_role_res"
#

/*!40000 ALTER TABLE `ref_role_res` DISABLE KEYS */;
INSERT INTO `ref_role_res` VALUES (13,2,1),(14,2,2),(15,3,1),(18,1,1),(19,1,2),(20,3,2),(22,1,3),(23,1,4),(24,1,5),(25,1,6);
/*!40000 ALTER TABLE `ref_role_res` ENABLE KEYS */;

#
# Structure for table "resource"
#

DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `patern` varchar(40) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  `res_type` varchar(10) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

#
# Data for table "resource"
#

/*!40000 ALTER TABLE `resource` DISABLE KEYS */;
/*!40000 ALTER TABLE `resource` ENABLE KEYS */;

#
# Structure for table "role"
#

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

#
# Data for table "role"
#

/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'admin',1),(2,'user',1),(3,'guest',1);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;

#
# Structure for table "user"
#

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account` varchar(40) DEFAULT NULL,
  `mobile` varchar(20) DEFAULT NULL,
  `passwd` varchar(40) DEFAULT NULL,
  `avatar` varchar(180) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `nick_name` varchar(40) DEFAULT NULL,
  `ticket` varchar(40) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

#
# Data for table "user"
#

/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','18600000000','d060812a3a1af12643a74a4d3b6d492d','admin@qq.com','2018-02-23 11:32:32','winlion',NULL,1,'admin@qq.com',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
