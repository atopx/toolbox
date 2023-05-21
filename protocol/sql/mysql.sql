CREATE TABLE user
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    username    VARCHAR(64) DEFAULT '' NOT NULL COMMENT '登录用户',
    password    VARCHAR(64) DEFAULT '' NOT NULL COMMENT '登录密码',
    status      INT         DEFAULT 0  NOT NULL COMMENT '用户状态',
    create_time BIGINT      DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT      DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT      DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户基础表';
CREATE UNIQUE INDEX uname_index ON user (username);

CREATE TABLE role
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    name        VARCHAR(64) DEFAULT '' NOT NULL COMMENT '名称',
    nature      INT         DEFAULT 0  NOT NULL COMMENT '角色特征',
    creator     INT         DEFAULT 0  NOT NULL COMMENT '创建人',
    updater     INT         DEFAULT 0  NOT NULL COMMENT '更新人',
    create_time BIGINT      DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT      DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT      DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '角色表';

CREATE TABLE user_role_ref
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    user_id     INT    DEFAULT 0 NOT NULL COMMENT '用户ID',
    role_id     INT    DEFAULT 0 NOT NULL COMMENT '角色ID',
    creator     INT    DEFAULT 0 NOT NULL COMMENT '创建人',
    updater     INT    DEFAULT 0 NOT NULL COMMENT '更新人',
    create_time BIGINT DEFAULT 0 NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT DEFAULT 0 NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT DEFAULT 0 NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户角色关联表';
CREATE UNIQUE INDEX user_role_index ON user_role_ref (user_id, role_id);


CREATE TABLE access
(
    id     INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    path   VARCHAR(64) DEFAULT '' NOT NULL COMMENT '资源路径',
    method VARCHAR(8)  DEFAULT '' NOT NULL COMMENT '请求方法',
    status INT         DEFAULT 0  NOT NULL DEFAULT 0 COMMENT '资源状态'
) ENGINE = INNODB COMMENT = '请求资源表';
CREATE UNIQUE INDEX access_request ON access (path, method);


CREATE TABLE auth_token
(
    id            INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    user_id       INT          DEFAULT 0  NOT NULL COMMENT '用户ID',
    access_token  VARCHAR(255) DEFAULT '' NOT NULL COMMENT '登录token',
    refresh_token VARCHAR(255) DEFAULT '' NOT NULL COMMENT '登录token',
    issued_time   BIGINT       DEFAULT 0  NOT NULL COMMENT '签发时间 时间戳：秒',
    expire_time   BIGINT       DEFAULT 0  NOT NULL COMMENT '过期时间 时间戳：秒',
    create_time   BIGINT       DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time   BIGINT       DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time   BIGINT       DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '用户令牌表';
CREATE UNIQUE INDEX access_token_idx ON auth_token (access_token);
CREATE UNIQUE INDEX refresh_token_idx ON auth_token (refresh_token);

CREATE TABLE permission
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    access_id   INT    DEFAULT 0 NOT NULL COMMENT '资源ID',
    role_id     INT    DEFAULT 0 NOT NULL COMMENT '角色ID',
    creator     INT    DEFAULT 0 NOT NULL COMMENT '创建人',
    updater     INT    DEFAULT 0 NOT NULL COMMENT '更新人',
    create_time BIGINT DEFAULT 0 NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT DEFAULT 0 NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT DEFAULT 0 NOT NULL COMMENT '删除时间 时间戳：秒'
) ENGINE = INNODB COMMENT = '角色资源权限表';
CREATE UNIQUE INDEX role_access ON permission (access_id, role_id);


create table label
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    code        VARCHAR(36)  DEFAULT 0  NOT NULL COMMENT '来源,枚举',
    name        VARCHAR(128) DEFAULT '' NOT NULL COMMENT '名称',
    creator     INT          DEFAULT 0  NOT NULL COMMENT '创建人',
    updater     INT          DEFAULT 0  NOT NULL COMMENT '更新人',
    create_time BIGINT       DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT       DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT       DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
) COMMENT '标签';
CREATE INDEX code_idx ON label (code);

create table note
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    title       VARCHAR(128) DEFAULT '' NOT NULL COMMENT '标题',
    content     MEDIUMTEXT              null COMMENT '笔记内容',
    creator     INT          DEFAULT 0  NOT NULL COMMENT '创建人',
    updater     INT          DEFAULT 0  NOT NULL COMMENT '更新人',
    create_time BIGINT       DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT       DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT       DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
) COMMENT '笔记';
create fullTEXT INDEX note_full ON note (title, content);

create table note_label
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    note_id     INT    DEFAULT 0 NOT NULL COMMENT '笔记ID',
    label_id    INT    DEFAULT 0 NOT NULL COMMENT '标签ID',
    label_type  INT    DEFAULT 0 NOT NULL COMMENT '标签类型',
    creator     INT    DEFAULT 0 NOT NULL COMMENT '创建人',
    updater     INT    DEFAULT 0 NOT NULL COMMENT '更新人',
    create_time BIGINT DEFAULT 0 NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT DEFAULT 0 NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT DEFAULT 0 NOT NULL COMMENT '删除时间 时间戳：秒'
) COMMENT '笔记标签';
create index note_idx on note_label(note_id);
create index label_idx on note_label(note_id);
create unique index note_label_type_idx on note_label(note_id, label_id, label_type);

create table book
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    name        VARCHAR(128) DEFAULT '' NOT NULL COMMENT '名称',
    author      VARCHAR(128) DEFAULT '' NOT NULL COMMENT '作者',
    src         VARCHAR(256) DEFAULT '' NOT NULL COMMENT '链接',
    cover       VARCHAR(256) DEFAULT '' NOT NULL COMMENT '封面',
    label       VARCHAR(32)  DEFAULT '' NOT NULL COMMENT '标签/分类',
    intro       TEXT                    null COMMENT '简介',
    state       INT          DEFAULT 0  NOT NULL COMMENT '任务状态',
    last_modify BIGINT       DEFAULT 0  NOT NULL COMMENT '最后修改时间 时间戳：秒',
    create_time BIGINT       DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT       DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT       DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
);
CREATE UNIQUE INDEX book_src_idx ON book (src);


create table chapter
(
    id          INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    book_id     INT          DEFAULT 0  NOT NULL COMMENT 'book-id',
    code        INT          DEFAULT 0  NOT NULL COMMENT '章节code',
    src         VARCHAR(256) DEFAULT '' NOT NULL COMMENT '链接',
    title       VARCHAR(256) DEFAULT '' NOT NULL COMMENT '章节标题',
    state       INT          DEFAULT 0  NOT NULL COMMENT '任务状态',
    content     TEXT                    null COMMENT '章节内容',
    create_time BIGINT       DEFAULT 0  NOT NULL COMMENT '创建时间 时间戳：秒',
    update_time BIGINT       DEFAULT 0  NOT NULL COMMENT '更新时间 时间戳：秒',
    delete_time BIGINT       DEFAULT 0  NOT NULL COMMENT '删除时间 时间戳：秒'
);
CREATE UNIQUE INDEX chapter_src_idx ON chapter (src);
CREATE UNIQUE INDEX book_chapter_code ON chapter (book_id, code);
