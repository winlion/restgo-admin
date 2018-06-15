/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : restgo-admin

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2018-06-15 12:40:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `config`
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `name` varchar(255) NOT NULL,
  `value` varchar(1024) DEFAULT NULL,
  `label` varchar(40) DEFAULT NULL,
  `format` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES ('guestid', '3', '用户作为游客身份访问默认角色ID', '^\\d+$');
INSERT INTO `config` VALUES ('role', '1', '用户注册默认角色id', '^\\d+$');

-- ----------------------------
-- Table structure for `ref_role_res`
-- ----------------------------
DROP TABLE IF EXISTS `ref_role_res`;
CREATE TABLE `ref_role_res` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleid` int(11) DEFAULT NULL,
  `resid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of ref_role_res
-- ----------------------------
INSERT INTO `ref_role_res` VALUES ('13', '2', '1');
INSERT INTO `ref_role_res` VALUES ('14', '2', '2');
INSERT INTO `ref_role_res` VALUES ('15', '3', '1');
INSERT INTO `ref_role_res` VALUES ('18', '1', '1');
INSERT INTO `ref_role_res` VALUES ('19', '1', '2');
INSERT INTO `ref_role_res` VALUES ('20', '3', '2');
INSERT INTO `ref_role_res` VALUES ('22', '1', '3');
INSERT INTO `ref_role_res` VALUES ('23', '1', '4');
INSERT INTO `ref_role_res` VALUES ('24', '1', '5');
INSERT INTO `ref_role_res` VALUES ('25', '1', '6');

-- ----------------------------
-- Table structure for `resource`
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `patern` varchar(40) DEFAULT NULL,
  `name` varchar(40) DEFAULT NULL,
  `res_type` varchar(10) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES ('1', '^\\s+$', '用户管理', 'mod', '0', '1');
INSERT INTO `resource` VALUES ('2', '/user/list.shtml', '用户管理页面', 'page', '1', '1');
INSERT INTO `resource` VALUES ('3', '/user/search', '搜索用户', 'api', '1', '1');
INSERT INTO `resource` VALUES ('4', '^\\s+$', '系统设置', 'mod', '0', '1');
INSERT INTO `resource` VALUES ('5', '/setting/role.shtml', '角色设置', 'page', '4', '1');
INSERT INTO `resource` VALUES ('6', '/setting/site.shtml', '参数设置', 'page', '4', '1');

-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', 'admin', '1');
INSERT INTO `role` VALUES ('2', 'user', '1');
INSERT INTO `role` VALUES ('3', 'guest', '1');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account` varchar(40) DEFAULT NULL,
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机号',
  `passwd` varchar(40) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `avatar` varchar(180) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `nick_name` varchar(40) DEFAULT NULL,
  `ticket` varchar(40) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `stat` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin@qq.com', '18600000000', 'd060812a3a1af12643a74a4d3b6d492d', 'admin@qq.com', '', '2018-02-23 11:32:32', 'winlion', '', '1', '0');
