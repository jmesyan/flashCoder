SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `flash_interface`;
DROP TABLE IF EXISTS `flash_behavior`;
DROP TABLE IF EXISTS `flash_coding`;
DROP TABLE IF EXISTS `flash_template`;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `flash_interface` (
    `fid` int(11) NOT NULL,
    `name` varchar(20) NOT NULL,
    `head` int(11) NOT NULL,
    `tail` int(11) NOT NULL,
    `desc` varchar(255) NOT NULL
);

CREATE TABLE `flash_behavior` (
    `bid` int(11) NOT NULL,
    `name` varchar(20) NOT NULL,
    `behavior` text(0) NOT NULL
);

CREATE TABLE `flash_coding` (
    `id` int(11) NOT NULL,
    `tid` int(11) NOT NULL,
    `code` text(0) NOT NULL
);

CREATE TABLE `flash_template` (
    `id` int(11) NOT NULL,
    `name` varchar(20) NOT NULL,
    `flow` text(0) NOT NULL
);
