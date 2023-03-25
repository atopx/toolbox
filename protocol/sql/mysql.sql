CREATE DATABASE IF NOT EXISTS super DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE user (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    username VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录用户',
    password VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录密码',
    status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '用户状态',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户基础表';
ALTER TABLE user ADD UNIQUE uname_index (username);

CREATE TABLE role (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    name VARCHAR ( 64 ) NOT NULL default '' COMMENT '名称',
    nature INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '角色特征',
    creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
    updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '角色表';

CREATE TABLE user_role_ref (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    user_id INT ( 11 ) NOT NULL COMMENT '用户ID',
    role_id INT ( 11 ) NOT NULL COMMENT '角色ID',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户角色关联表';
ALTER TABLE user_role_ref ADD UNIQUE user_role_index (user_id, role_id);


CREATE TABLE access (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    path VARCHAR ( 64 ) NOT NULL default '' COMMENT '资源路径',
    handler VARCHAR ( 64 ) NOT NULL default '' COMMENT '处理器',
    method VARCHAR ( 8 ) NOT NULL COMMENT '请求方法',
    status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '资源状态',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '请求资源表';
ALTER TABLE access ADD UNIQUE access_request (path, method);

CREATE TABLE auth_token (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    user_id INT ( 11 ) NOT NULL COMMENT '用户ID',
    access_token VARCHAR (255) NOT NULL COMMENT '登录token',
    refresh_token VARCHAR (255) NOT NULL COMMENT '登录token',
    issued_time BIGINT NOT NULL COMMENT '签发时间 时间戳：秒',
    expire_time BIGINT NOT NULL COMMENT '过期时间 时间戳：秒',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户令牌表';
ALTER TABLE auth_token ADD UNIQUE INDEX access_token_idx(access_token);
ALTER TABLE auth_token ADD UNIQUE INDEX refresh_token_idx(refresh_token);

CREATE TABLE permission (
	id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	access_id INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '资源ID',
	role_id INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '角色ID',
	creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '角色资源权限表';
ALTER TABLE permission ADD UNIQUE role_access (access_id, role_id);

CREATE TABLE computer (
	id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	name VARCHAR ( 64 ) NOT NULL default '' COMMENT '名称',
	username VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录用户',
	password VARCHAR ( 64 ) NOT NULL default '' COMMENT '登录密码',
	lan_hostname VARCHAR ( 64 ) NOT NULL default '' COMMENT '局域网地址',
	wan_hostname VARCHAR ( 64 ) NOT NULL default '' COMMENT '广域网地址',
	address CHAR ( 12 ) NOT NULL default '' COMMENT '物理地址',
	power_status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '电源状态',
	scan_time BIGINT NOT NULL DEFAULT 0 COMMENT '最后一次扫描时间',
	creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '主机表';
ALTER TABLE computer ADD UNIQUE address_idx(address);


CREATE TABLE computer_port (
	id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
	computer_id INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '主机ID',
	port INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '网络端口',
	protocol INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '应用协议',
	tranport INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '传输协议',
	use_type INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '用途',
	descr varchar(1024)  NOT NULL default '' COMMENT '描述',
	creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '主机端口表';
ALTER TABLE computer_port ADD UNIQUE cport (computer_id, port);
ALTER TABLE computer_port ADD INDEX cpid(computer_id);

CREATE TABLE novel (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    name varchar(32) NOT NULL default '' COMMENT '书名',
    source varchar(128) NOT NULL default '' COMMENT '书源',
    author varchar(64) NOT NULL default '' COMMENT '作者',
    intro varchar(4096) NOT NULL default '' COMMENT '简介',
    status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '状态',
    lasted BIGINT NOT NULL COMMENT '(源)最近更新时间 时间戳：秒',
    scan_num INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '扫描次数',
    error_num INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '异常次数',
    no_change_num INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '未发生变化次数',
    creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
    updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '书表';
ALTER TABLE novel ADD UNIQUE novel_source (source);


CREATE TABLE novel_chapter (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    novel_id INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '书ID',
    name varchar(32) NOT NULL default '' COMMENT '章节名称',
    source varchar(128) NOT NULL default '' COMMENT '章节源',
    status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '章节状态',
    content text not null COMMENT '章节内容',
    message varchar(256) NOT NULL default '' COMMENT '异常消息',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '书-章节表';
ALTER TABLE novel_chapter ADD UNIQUE chapter_source (source);


CREATE TABLE novel_task (
    id INT ( 11 ) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    novel_id INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '书ID',
    status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '章节状态',
    create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '书-任务表';
