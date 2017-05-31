/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : flashcoder

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-05-31 23:31:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for flash_behavior
-- ----------------------------
DROP TABLE IF EXISTS `flash_behavior`;
CREATE TABLE `flash_behavior` (
  `bid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `opid` int(11) unsigned NOT NULL COMMENT '操作id',
  `bname` varchar(50) NOT NULL COMMENT '行为名称',
  `paramsdef` text NOT NULL COMMENT '默认参数',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updtime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`bid`,`bname`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_behavior
-- ----------------------------
INSERT INTO `flash_behavior` VALUES ('1', '0', 'hello', '', null, '1496014152', '1496014152');
INSERT INTO `flash_behavior` VALUES ('2', '0', 'OpenFile', '{\"path\":\"a.txt\"}', null, '1496066314', '1496066314');
INSERT INTO `flash_behavior` VALUES ('3', '0', 'WriteFile', '{\"content\":\"this is a text\"}', null, '1496066380', '1496066380');
INSERT INTO `flash_behavior` VALUES ('4', '0', 'CloseFile', '{\"\":\"\"}', null, '1496103734', '1496103734');

-- ----------------------------
-- Table structure for flash_cron
-- ----------------------------
DROP TABLE IF EXISTS `flash_cron`;
CREATE TABLE `flash_cron` (
  `crid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `second` int(5) NOT NULL DEFAULT '0' COMMENT '秒',
  `minute` int(5) NOT NULL DEFAULT '0' COMMENT '分',
  `hour` int(5) NOT NULL DEFAULT '0' COMMENT '时',
  `day` int(5) NOT NULL DEFAULT '0' COMMENT '日期',
  `month` int(5) NOT NULL DEFAULT '0' COMMENT '月份',
  `week` int(5) NOT NULL DEFAULT '0' COMMENT '周',
  `tid` int(11) NOT NULL DEFAULT '0' COMMENT '任务id',
  PRIMARY KEY (`crid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_cron
-- ----------------------------

-- ----------------------------
-- Table structure for flash_operate
-- ----------------------------
DROP TABLE IF EXISTS `flash_operate`;
CREATE TABLE `flash_operate` (
  `opid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `opname` varchar(30) NOT NULL DEFAULT '' COMMENT '操作名称',
  `optag` varchar(100) NOT NULL COMMENT '操作标识',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`opid`,`opname`,`optag`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_operate
-- ----------------------------
INSERT INTO `flash_operate` VALUES ('1', '打开文件', 'OpenFile', '基本打开文件操作', '1496236427');
INSERT INTO `flash_operate` VALUES ('2', '写入文件', 'WriteFile', null, '1496236427');
INSERT INTO `flash_operate` VALUES ('3', '关闭文件', 'CloseFile', null, '1496236427');

-- ----------------------------
-- Table structure for flash_task
-- ----------------------------
DROP TABLE IF EXISTS `flash_task`;
CREATE TABLE `flash_task` (
  `tid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '任务id',
  `tcate` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-基础任务 2-复合任务',
  `tsubs` text COMMENT '子任务id-基础任务为0',
  `bids` text COMMENT '行为id-复合认为为0',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updtime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task
-- ----------------------------

-- ----------------------------
-- Table structure for flash_task_behavior
-- ----------------------------
DROP TABLE IF EXISTS `flash_task_behavior`;
CREATE TABLE `flash_task_behavior` (
  `tbid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `bid` int(11) NOT NULL DEFAULT '0' COMMENT '行为id',
  `tid` int(11) NOT NULL DEFAULT '0' COMMENT '任务id',
  `ctid` int(11) NOT NULL DEFAULT '0' COMMENT '复合任务id',
  `border` int(11) NOT NULL DEFAULT '0' COMMENT '行为时序',
  `torder` int(11) NOT NULL DEFAULT '0' COMMENT '基础任务时序',
  `paramsin` text NOT NULL COMMENT '输入参数',
  PRIMARY KEY (`tbid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task_behavior
-- ----------------------------
