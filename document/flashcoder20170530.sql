/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : flashcoder

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-05-30 12:47:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for flash_behavior
-- ----------------------------
DROP TABLE IF EXISTS `flash_behavior`;
CREATE TABLE `flash_behavior` (
  `bid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '行为名称',
  `paramsin` text NOT NULL COMMENT '输入参数',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `updtime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`bid`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_behavior
-- ----------------------------
INSERT INTO `flash_behavior` VALUES ('1', 'hello', '', '1496014152', '1496014152');
INSERT INTO `flash_behavior` VALUES ('2', 'OpenFile', '{\"path\":\"a.txt\"}', '1496066314', '1496066314');
INSERT INTO `flash_behavior` VALUES ('3', 'WriteFile', '{\"content\":\"this is a text\"}', '1496066380', '1496066380');
INSERT INTO `flash_behavior` VALUES ('4', 'CloseFile', '{\"\":\"\"}', '1496103734', '1496103734');

-- ----------------------------
-- Table structure for flash_coding
-- ----------------------------
DROP TABLE IF EXISTS `flash_coding`;
CREATE TABLE `flash_coding` (
  `id` int(11) NOT NULL,
  `tid` int(11) NOT NULL,
  `code` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_coding
-- ----------------------------

-- ----------------------------
-- Table structure for flash_interface
-- ----------------------------
DROP TABLE IF EXISTS `flash_interface`;
CREATE TABLE `flash_interface` (
  `fid` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `head` int(11) NOT NULL,
  `tail` int(11) NOT NULL,
  `desc` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_interface
-- ----------------------------

-- ----------------------------
-- Table structure for flash_template
-- ----------------------------
DROP TABLE IF EXISTS `flash_template`;
CREATE TABLE `flash_template` (
  `id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `flow` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of flash_template
-- ----------------------------
