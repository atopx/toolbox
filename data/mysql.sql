CREATE TABLE `su_computer` (
	`id` INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	`name` VARCHAR ( 64 ) NOT NULL default '' COMMENT '名称',
	`username` VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录用户',
	`password` VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录密码',
	`lan_hostname` VARCHAR ( 64 ) NOT NULL default '' COMMENT '局域网地址',
	`wan_hostname` VARCHAR ( 64 ) NOT NULL default '' COMMENT '广域网地址',
	`address` CHAR ( 12 ) NOT NULL default '' COMMENT '物理地址',
	`power_status` TINYINT NOT NULL DEFAULT 0 COMMENT '电源状态',
	`scan_time` BIGINT NOT NULL DEFAULT 0 COMMENT '最后一次扫描时间',
	`creator` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	`updater` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	`create_time` BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	`update_time` BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
  `delete_time` BIGINT NOT NULL COMMENT '删除时间 时间戳：秒' 
) ENGINE = INNODB COMMENT = '主机表';


CREATE TABLE `su_computer_port` (
	`id` INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	`computer_id` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '主机ID',
	`port` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '网络端口',
	`protocol` TINYINT NOT NULL DEFAULT 0 COMMENT '应用协议',
	`tranport` TINYINT NOT NULL DEFAULT 0 COMMENT '传输协议',
	`use_type` TINYINT NOT NULL DEFAULT 0 COMMENT '用途',
	`descr` varchar(1024)  NOT NULL default '' COMMENT '描述',
	`creator` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	`updater` INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	`create_time` BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	`update_time` BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
  `delete_time` BIGINT NOT NULL COMMENT '删除时间 时间戳：秒' 
) ENGINE = INNODB COMMENT = '主机端口表';


CREATE TABLE `su_user` (
	`id` INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	`name` VARCHAR ( 64 ) NOT NULL default '' COMMENT '姓名',
	`username` VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录用户',
	`password` VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录密码',
	`role` TINYINT NOT NULL DEFAULT 0 COMMENT '用户角色',
	`status` TINYINT NOT NULL DEFAULT 0 COMMENT '用户状态',
	`create_time` BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	`update_time` BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
  `delete_time` BIGINT NOT NULL COMMENT '删除时间 时间戳：秒' 
) ENGINE = INNODB COMMENT = '用户基础表';
