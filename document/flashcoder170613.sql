/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : flashcoder

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-06-13 05:55:49
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_behavior
-- ----------------------------
INSERT INTO `flash_behavior` VALUES ('7', '1', '打开文件', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"b.txt\"}]', '打开一个文件', '1496499317', '1496502218');
INSERT INTO `flash_behavior` VALUES ('10', '2', '写入文件', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"today a happy day, i want to solve a problem\"}]', '把数据写入一个文件', '1496502347', '1496529218');
INSERT INTO `flash_behavior` VALUES ('11', '3', '关闭文件', '[]', '关闭一个文件句柄', '1496502399', '1496502399');
INSERT INTO `flash_behavior` VALUES ('15', '4', '提示休息', '[{\"type\":\"3\",\"name\":\"msg\",\"value\":\"主人你该休息了，加油哦！\"}]', '休息提醒', '1497190209', '1497190209');

-- ----------------------------
-- Table structure for flash_cron
-- ----------------------------
DROP TABLE IF EXISTS `flash_cron`;
CREATE TABLE `flash_cron` (
  `crid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `second` varchar(10) NOT NULL DEFAULT '*' COMMENT '秒：0-59 * / , -',
  `minute` varchar(10) NOT NULL DEFAULT '*' COMMENT '分：0-59 * / , -',
  `hour` varchar(10) NOT NULL DEFAULT '*' COMMENT '时：0-23 * / , -',
  `day` varchar(10) NOT NULL DEFAULT '*' COMMENT '每月第几天 ：1-31 * / , - ?',
  `month` varchar(10) NOT NULL DEFAULT '*' COMMENT '月份：1-12   * / , -',
  `week` varchar(10) NOT NULL DEFAULT '*' COMMENT '每周第几天：0-6  * / , - ?',
  `tid` int(11) NOT NULL DEFAULT '0' COMMENT '任务id',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0-开启 1-关闭',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`crid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_cron
-- ----------------------------
INSERT INTO `flash_cron` VALUES ('1', '*/2', '*', '*', '*', '*', '*', '4', '1', '1497304370');

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_operate
-- ----------------------------
INSERT INTO `flash_operate` VALUES ('1', '打开文件', 'OpenFile', '打开文件', '1497180376');
INSERT INTO `flash_operate` VALUES ('2', '写入文件', 'WriteFile', '写入文件', '1496236427');
INSERT INTO `flash_operate` VALUES ('3', '关闭文件', 'CloseFile', '关闭文件句柄', '1496236427');
INSERT INTO `flash_operate` VALUES ('4', '提示框', 'MsgTip', '提示内容', '1497190102');

-- ----------------------------
-- Table structure for flash_task
-- ----------------------------
DROP TABLE IF EXISTS `flash_task`;
CREATE TABLE `flash_task` (
  `tid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '任务id',
  `tname` varchar(200) NOT NULL COMMENT '任务名称',
  `tcate` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1-基础任务 2-复合任务',
  `tsubs` text COMMENT '子任务id-基础任务为0',
  `bids` text COMMENT '行为id-复合认为为0',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updtime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task
-- ----------------------------
INSERT INTO `flash_task` VALUES ('3', '写数据到文件', '1', null, '[{\"ItemId\":7,\"ItemName\":\"打开文件\"},{\"ItemId\":10,\"ItemName\":\"写入文件\"},{\"ItemId\":11,\"ItemName\":\"关闭文件\"}]', '1497180500', '1497180500');
INSERT INTO `flash_task` VALUES ('4', '提醒休息', '1', null, '[{\"ItemId\":15,\"ItemName\":\"提示休息\"}]', '1497190228', '1497190228');

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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task_behavior
-- ----------------------------
INSERT INTO `flash_task_behavior` VALUES ('10', '7', '3', '0', '0', '0', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"aaa.txt\"}]');
INSERT INTO `flash_task_behavior` VALUES ('11', '10', '3', '0', '1', '0', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"today a happy day, it a new begin in my wold; \"}]');
INSERT INTO `flash_task_behavior` VALUES ('12', '11', '3', '0', '2', '0', '[]');
INSERT INTO `flash_task_behavior` VALUES ('13', '15', '4', '0', '0', '0', '[{\"type\":\"3\",\"name\":\"msg\",\"value\":\"主人你该休息了，加油哦！\"}]');
