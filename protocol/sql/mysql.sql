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
    creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
    updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
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
	mac_address CHAR ( 12 ) NOT NULL default '' COMMENT '物理地址',
	power_status INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '电源状态',
	scan_time BIGINT NOT NULL DEFAULT 0 COMMENT '最后一次扫描时间',
	creator INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '创建人',
	updater INT ( 11 ) NOT NULL DEFAULT 0 COMMENT '更新人',
	create_time BIGINT NOT NULL COMMENT '创建时间 时间戳：秒',
	update_time BIGINT NOT NULL COMMENT '最后更新时间 时间戳：秒',
    delete_time BIGINT NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '主机表';
ALTER TABLE computer ADD UNIQUE address_idx(mac_address);

create table folder
(
    id          int primary key auto_increment comment '主键',
    parent_id   int          default 0  not null comment '父级ID',
    source      int          default 0  not null comment '来源,枚举',
    name        varchar(128) default '' not null comment '名称',
    path        varchar(256) default '' not null comment '存储路径,虚拟文件夹为空',
    creator     int          default 0  not null comment '创建人',
    updater     int          default 0  not null comment '更新人',
    create_time bigint                  not null comment '创建时间 时间戳：秒',
    update_time bigint                  not null comment '最后更新时间 时间戳：秒',
    delete_time bigint                  not null comment '删除时间 时间戳：秒',
    constraint source_name_idx
        unique (name, source)
) comment '文件夹';

create table label
(
    id          int primary key auto_increment comment '主键',
    source      int          default 0  not null comment '来源,枚举',
    name        varchar(128) default '' not null comment '名称',
    creator     int          default 0  not null comment '创建人',
    updater     int          default 0  not null comment '更新人',
    create_time bigint                  not null comment '创建时间 时间戳：秒',
    update_time bigint                  not null comment '最后更新时间 时间戳：秒',
    delete_time bigint                  not null comment '删除时间 时间戳：秒',
    constraint source_name_idx
        unique (name, source)
) comment '标签';

create table note
(
    id          int primary key auto_increment comment '主键',
    folder_id   int                           not null comment '文件夹ID',
    topic_id    int          default 0        not null comment '主题ID',
    sign        varchar(64)  default ''       not null comment '签名:内容+时间',
    title       varchar(128) default '未命名'  not null comment '标题',
    public      tinyint(1)   default '0'      not null comment '公开的',
    content     mediumtext                    not null comment '笔记内容',
    creator     int          default 0        not null comment '创建人',
    updater     int          default 0        not null comment '更新人',
    create_time bigint                        not null comment '创建时间 时间戳：秒',
    update_time bigint                        not null comment '最后更新时间 时间戳：秒',
    delete_time bigint                        not null comment '删除时间 时间戳：秒'
) comment '笔记';
create index folder_idx on note(folder_id);
create index topic_idx on note(topic_id);
create fulltext index note_full
    on note(title, content);

create table note_topic
(
    id          int primary key auto_increment comment '主键',
    name        varchar(128) default '' not null comment '名称',
    creator     int          default 0  not null comment '创建人',
    updater     int          default 0  not null comment '更新人',
    create_time bigint                  not null comment '创建时间 时间戳：秒',
    update_time bigint                  not null comment '最后更新时间 时间戳：秒',
    delete_time bigint                  not null comment '删除时间 时间戳：秒'

) comment '笔记主题';
create index note_topic_name_idx on note_topic(name);

create table note_label
(
    id          int primary key auto_increment comment '主键',
    note_id     int           not null comment '笔记ID',
    label_id    int           not null comment '标签ID',
    creator     int default 0 not null comment '创建人',
    updater     int default 0 not null comment '更新人',
    create_time bigint        not null comment '创建时间 时间戳：秒',
    update_time bigint        not null comment '最后更新时间 时间戳：秒',
    delete_time bigint        not null comment '删除时间 时间戳：秒',
    constraint note_label_idx
        unique (note_id, label_id)
) comment '笔记标签';

create table file
(
    id          int primary key auto_increment comment '主键',
    folder_id   int                           not null comment '文件夹ID',
    sign        varchar(64)  default ''       not null comment '内容签名',
    name        varchar(128) default '未命名'  not null comment '文件名',
    public      tinyint(1)   default '0'      not null comment '公开的',
    share_url   varchar(128) default ''       not null comment '共享链接',
    creator     int          default 0        not null comment '创建人',
    updater     int          default 0        not null comment '更新人',
    create_time bigint                        not null comment '创建时间 时间戳：秒',
    update_time bigint                        not null comment '最后更新时间 时间戳：秒',
    delete_time bigint                        not null comment '删除时间 时间戳：秒'
) comment '文件';
create index file_folder_idx on file(folder_id);
create index file_sign_idx on file(sign);
create index file_name_idx on file(name);
create unique index folder_file_sign_idx on file(folder_id, sign);

create table book
(
    id          int primary key auto_increment comment '主键',
    name        varchar(128)  default ''     not null comment '名称',
    author      varchar(128)  default ''     not null comment '作者',
    src         varchar(256)  default ''     not null comment '链接',
    cover       varchar(256)  default ''     not null comment '封面',
    label       varchar(32)   default ''     not null comment '标签/分类',
    intro       varchar(4096) default ''     not null comment '简介',
    state       varchar(32)   default 'NONE' not null comment '状态',
    last_modify bigint        default 0      not null comment '最后修改时间 时间戳：秒',
    create_time bigint        default 0      not null comment '创建时间 时间戳：秒',
    update_time bigint        default 0      not null comment '更新时间 时间戳：秒',
    delete_time bigint        default 0      not null comment '删除时间 时间戳：秒',
    constraint src_idx unique (src)
);

create table chapter
(
    id          int primary key auto_increment comment '主键',
    book_id     int                         not null comment 'book-id',
    code        int                         not null comment '章节code',
    src         varchar(256) default ''     not null comment '链接',
    title       varchar(256) default ''     not null comment '章节标题',
    state       varchar(32)  default 'NONE' not null comment '状态',
    content     text comment '章节内容',
    create_time bigint       default 0      not null comment '创建时间 时间戳：秒',
    update_time bigint       default 0      not null comment '更新时间 时间戳：秒',
    delete_time bigint       default 0      not null comment '删除时间 时间戳：秒',
    constraint books unique (book_id, code),
    constraint src_idx unique (src)
);
