/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : flashcoder

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-06-11 12:49:48
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_behavior
-- ----------------------------
INSERT INTO `flash_behavior` VALUES ('7', '1', '打开文件', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"b.txt\"}]', '打开一个文件', '1496499317', '1496502218');
INSERT INTO `flash_behavior` VALUES ('10', '2', '写入文件', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"today a happy day, i want to solve a problem\"}]', '把数据写入一个文件', '1496502347', '1496529218');
INSERT INTO `flash_behavior` VALUES ('11', '3', '关闭文件', '[]', '关闭一个文件句柄', '1496502399', '1496502399');
INSERT INTO `flash_behavior` VALUES ('12', '1', '打开文件2', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"hello.txt\"}]', '这是备注吗', '1496537986', '1496675344');
INSERT INTO `flash_behavior` VALUES ('13', '2', '写入文件2', '[{\"type\":\"1\",\"name\":\"content\",\"value\":\"这是什么内容呢，我也想了好半天\"}]', '这是一个备注', '1496538017', '1496673557');
INSERT INTO `flash_behavior` VALUES ('14', '1', '', '[]', '', '1497155937', '1497155937');

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

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
INSERT INTO `flash_operate` VALUES ('2', '写入文件', 'WriteFile', '写入文件', '1496236427');
INSERT INTO `flash_operate` VALUES ('3', '关闭文件', 'CloseFile', '关闭文件句柄', '1496236427');

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task
-- ----------------------------
INSERT INTO `flash_task` VALUES ('1', '写入文件测试', '1', null, '[{\"ItemId\":7,\"ItemName\":\"打开文件\"},{\"ItemId\":10,\"ItemName\":\"写入文件\"},{\"ItemId\":11,\"ItemName\":\"关闭文件\"}]', '1496585469', '1496585469');
INSERT INTO `flash_task` VALUES ('2', '重复写入文件测试', '2', '[{\"ItemId\":1,\"ItemName\":\"写入文件测试\"},{\"ItemId\":1,\"ItemName\":\"写入文件测试\"}]', null, '1496670375', '1496670375');

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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_task_behavior
-- ----------------------------
INSERT INTO `flash_task_behavior` VALUES ('1', '7', '1', '0', '0', '0', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"cbc.txt\"}]');
INSERT INTO `flash_task_behavior` VALUES ('2', '10', '1', '0', '1', '0', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"这是定时脚本内容测试 \\\\r\\\\n\"}]');
INSERT INTO `flash_task_behavior` VALUES ('3', '11', '1', '0', '2', '0', '[]');
INSERT INTO `flash_task_behavior` VALUES ('4', '7', '1', '2', '0', '0', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"d.txt\"}]');
INSERT INTO `flash_task_behavior` VALUES ('5', '10', '1', '2', '1', '0', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"today a happy day, i want to solve a problem\"}]');
INSERT INTO `flash_task_behavior` VALUES ('6', '11', '1', '2', '2', '0', '[]');
INSERT INTO `flash_task_behavior` VALUES ('7', '7', '1', '2', '0', '1', '[{\"type\":\"1\",\"name\":\"path\",\"value\":\"b.txt\"}]');
INSERT INTO `flash_task_behavior` VALUES ('8', '10', '1', '2', '1', '1', '[{\"type\":\"3\",\"name\":\"content\",\"value\":\"today a happy day, i want to solve a problem\"}]');
INSERT INTO `flash_task_behavior` VALUES ('9', '11', '1', '2', '2', '1', '[]');
