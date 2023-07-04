create table user
(
    id          bigint unsigned primary key auto_increment comment '主键',
    username    varchar(64)     default ''  not null comment '登录用户',
    password    varchar(64)     default ''  not null comment '登录密码',
    status      varchar(24)     default ''  not null comment '用户状态',
    role        varchar(24)     default ''  not null comment '用户角色',
    avatar      varchar(128)    default ''  not null comment '头像',
    email       varchar(64)     default ''  not null comment '邮箱',
    phone       varchar(12)     default ''  not null comment '手机号',
    create_time bigint unsigned default '0' not null comment '创建时间 时间戳：秒',
    update_time bigint unsigned default '0' not null comment '更新时间 时间戳：秒',
    delete_time bigint unsigned default '0' not null comment '删除时间 时间戳：秒',
    constraint uname_index unique (username)
) comment '用户基础表';


create table auth_token
(
    id            bigint unsigned primary key auto_increment comment '主键',
    user_id       bigint unsigned default 0  not null comment '用户ID',
    access_token  varchar(255)    default '' not null comment '登录token',
    refresh_token varchar(255)    default '' not null comment '登录token',
    issued_time   bigint unsigned default 0  not null comment '签发时间 时间戳：秒',
    expire_time   bigint unsigned default 0  not null comment '过期时间 时间戳：秒',
    create_time   bigint unsigned default 0  not null comment '创建时间 时间戳：秒',
    update_time   bigint unsigned default 0  not null comment '更新时间 时间戳：秒',
    delete_time   bigint unsigned default 0  not null comment '删除时间 时间戳：秒',
    constraint access_token_idx unique (access_token),
    constraint refresh_token_idx unique (refresh_token)
) comment '用户令牌表';
create index user_idx on auth_token (user_id);

create table file
(
    id          bigint unsigned primary key auto_increment comment '主键',
    owner       bigint unsigned not null default 0 comment '拥有者ID',
    parent_id   bigint unsigned               not null default 0 comment '父级文件夹',
    name        varchar(64)                   not null default '' comment '文件名',
    feature     enum ('FILE', 'FOLDER') not null default 'FILE' comment '文件特征',
    md5         varchar(45)                   not null default '' comment 'md5',
    size        bigint unsigned               not null default 0 comment '文件大小, KB',
    path        varchar(1024)                 not null default '' comment '文件路径',
    content_type varchar(128)                  not null default '' comment '文件类别',
    password     varchar(64)                   not null default '' comment '文件密码',
    create_time bigint unsigned               not null default '0' comment '创建时间 时间戳：秒',
    update_time bigint unsigned               not null default '0' comment '更新时间 时间戳：秒',
    delete_time bigint unsigned               not null default '0' comment '删除时间 时间戳：秒',
    creator     bigint unsigned default '0'      not null comment '创建人',
    updater     bigint unsigned default '0'      not null comment '修改人'
);
create index parent_idx on file (parent_id);


create table note_label
(
    id          bigint unsigned primary key auto_increment comment '主键',
    name        varchar(64)     not null default '' comment 'LABEL名称',
    note_id     bigint unsigned not null default 0 comment '笔记ID',
    create_time bigint unsigned not null default '0' comment '创建时间 时间戳：秒',
    update_time bigint unsigned not null default '0' comment '更新时间 时间戳：秒',
    delete_time bigint unsigned not null default '0' comment '删除时间 时间戳：秒'
);
create index note_idx on note_label (note_id);

create table note
(
    id          bigint unsigned primary key auto_increment comment '主键',
    title       varchar(128)    not null default '未命名' comment '笔记名称',
    folder_id   bigint unsigned not null default 0 comment '文件夹ID',
    topic       varchar(128)    not null default '' comment '主题名称',
    content     mediumtext      null comment '笔记内容',
    create_time bigint unsigned not null default '0' comment '创建时间 时间戳：秒',
    update_time bigint unsigned not null default '0' comment '更新时间 时间戳：秒',
    delete_time bigint unsigned not null default '0' comment '删除时间 时间戳：秒',
    creator     bigint unsigned default '0'      not null comment '创建人',
    updater     bigint unsigned default '0'      not null comment '修改人'
);
create index folder_idx on note (folder_id);
create fulltext index note_full on note (title, content);
