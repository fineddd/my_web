DROP DATABASE IF EXISTS `my_web`;
CREATE DATABASE `my_web` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `my_web`;
SET names utf8;

CREATE TABLE `platform` (
    `id` int(10) UNSIGNED NOT NULL, \
    `name` varchar(20) NOT NULL, \
    `note` varchar(100) NOT NULL, \
    `srcs` varchar(20000),
    PRIMARY KEY (`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `server` (
    `id` int(10) UNSIGNED NOT NULL, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `name` varchar(20) NOT NULL, \
    `address` varchar(100) NOT NULL, \
    PRIMARY KEY(`id`,`pfid`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `user` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `name` varchar(20) NOT NULL, \
    `password` varchar(20) NOT NULL, \
    `right` int(10) UNSIGNED NOT NULL, \
    `address` varchar(20) NOT NULL, \
    `lastlogintime` bigint(10) NOT NULL, \
    `status` int(10) UNSIGNED NOT NULL, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `right` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `name` varchar(20) NOT NULL, \
    `val` bigint(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `notice` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL,\
    `name` varchar(20) NOT NULL, \
    `content` varchar(100) NOT NULL, \
    `starttime` bigint(10) NOT NULL, \
    `endtime` bigint(10) NOT NULL, \
    `status` int(10) NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `compensation` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL, \
    `global` int(10) NOT NULL, \
    `playernames` varchar(100) NOT NULL, \
    `title` varchar(20) NOT NULL, \
    `contenttype` int(10) NOT NULL, \
    `content` varchar(100) NOT NULL, \
    `attachments` varchar(100) NOT NULL, \
    `sendtime` bigint(10) NOT NULL, \
    `savetime` bigint(10) NOT NULL, \
    `sender` int(10) NOT NULL, \
    `status` int(10) NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `operating_active` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL, \
    `activeid` int(10) UNSIGNED NOT NULL, \
    `viplevel` int(10) UNSIGNED NOT NULL, \
    `plans` varchar(200) NOT NULL, \
    `skin` int(10) UNSIGNED NOT NULL, \
    `starttime` bigint(10) NOT NULL, \
    `endtime` bigint(10) NOT NULL, \
    `status` int(10) NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `gmcmd_record` (
    `id` int (10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL, \
    `cmd` varchar(200) NOT NULL, \
    `sender` int(10) NOT NULL, \
    `status` int(10) NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `cdk_batch` (
    `id` int(10) UNSIGNED NOT NULL, \
    `type` int(10) UNSIGNED NOT NULL, \
    `batch` int(10) UNSIGNED NOT NULL, \
    `name` varchar(100) NOT NULL, \
    `note` varchar(100) NOT NULL, \
    `objs` varchar(100) NOT NULL, \
    `count` int(10) UNSIGNED NOT NULL, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `starttime` bigint(10) UNSIGNED NOT NULL, \
    `endtime` bigint(10) UNSIGNED NOT NULL, \
    `createtime` bigint(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `cdk_key` (
    `id` varchar(20) NOT NULL, \
    `batchid` int(10) UNSIGNED NOT NULL, \
    `idx` int(10) UNSIGNED NOT NULL, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL, \
    `playerid` int(10) UNSIGNED NOT NULL, \
    `used` int(10) UNSIGNED NOT NULL, \
    `usedtime` bigint(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `cdk_hd_batch` (
    `id` varchar(20) NOT NULL, \
    `name` varchar(100) NOT NULL, \
    `note` varchar(100) NOT NULL, \
    `objs` varchar(100) NOT NULL, \
    `starttime` bigint(10) UNSIGNED NOT NULL, \
    `endtime` bigint(10) UNSIGNED NOT NULL, \
    `createtime` bigint(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `cdk_hd_key` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `batchid` varchar(20) NOT NULL, \
    `pfid` int(10) UNSIGNED NOT NULL, \
    `servid` int(10) UNSIGNED NOT NULL, \
    `playerid` int(10) UNSIGNED NOT NULL, \
    `usedtime` bigint(10) UNSIGNED NOT NULL, \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `server_data` ( \
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, \
    `date` bigint(10) NOT NULL, \
    `pfid` int(10) NOT NULL,  \
    `servid` int(10) NOT NULL, \
    `src` varchar(20),\
    `new_player_count` int(10) NOT NULL, \
    `login_count` int(10) NOT NULL, \
    `max_online_count` int(10) NOT NULL, \
    `avg_online_count` int(10) NOT NULL, \
    `next_day_retain` int(10) NOT NULL, \
    `seven_day_retain` int(10) NOT NULL, \
    `month_day_retain` int(10) NOT NULL, \
    `recharge_count` int(10) NOT NULL, \
    `recharge_player` int(10) NOT NULL, \
    `recharge_number` int(10) NOT NULL, \
    `total_connect_count` int(10) NOT NULL, \
    `total_player_count` int (10) NOT NULL, \
    `total_recharge_count` int(10) NOT NULL, \
    `total_recharge_player` int(10) NOT NULL, \
    `total_recharge_number` int(10) NOT NULL, \
    `total_consume_count` bigint(10) NOT NULL, \
    KEY(`pfid`), \
    KEY(`date`), \
    KEY(`date`, `pfid`), \
    KEY(`date`,`pfid`,`servid`), \
    PRIMARY KEY(`id`) \
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

insert into `platform` set `id`=0, `name`='内部',`note`='内部平台';
insert into `right` set `name`='超级管理员',`val`=4294967295;
insert into `user` set `name`='admin',`password`='123456',`right`=1,address='',lastlogintime=0,status=0,pfid=0;
