
CREATE TABLE IF NOT EXISTS su_computer (
    id integer PRIMARY KEY AUTOINCREMENT, -- 自增ID
    name varchar ( 64 )  NOT NULL, -- 名称
    username varchar ( 64 ) NOT NULL DEFAULT '', -- 用户名
    password varchar ( 64 ) NOT NULL DEFAULT '', -- 用户密码
    lan_hostname varchar ( 64 ) NOT NULL DEFAULT '', -- 局域网地址
    wan_hostname varchar ( 64 ) NOT NULL DEFAULT '', -- 广域网地址
    address char ( 12 ) NOT NULL DEFAULT '', -- 物理地址
    power_status tinyint NOT NULL DEFAULT 0, -- 电源状态
    scan_time bigint NOT NULL DEFAULT 0, -- 最后一次扫描时间
    creator integer not null, -- 创建人
    updator integer not null, -- 更新人
    create_time bigint not null, -- 创建时间
    update_time bigint not null, -- 更新时间
    delete_time bigint not null default 0 -- 删除时间
);

CREATE TABLE IF NOT EXISTS su_computer_port (
    id integer PRIMARY KEY AUTOINCREMENT,
    computer_id integer not null, -- 主机ID
    port integer not null default 0, -- 网络端口 
    protocol tinyint not null default 0, -- 应用协议
    tranport tinyint not null default 0, -- 传输协议
    use_type tinyint not null default 0, -- 用途
    desc text not null default '', -- 备注
    creator integer not null, -- 创建人
    updator integer not null, -- 更新人
    create_time bigint not null, -- 创建时间
    update_time bigint not null, -- 更新时间
    delete_time bigint not null default 0 -- 删除时间
);

CREATE TABLE IF NOT EXISTS su_user (
    id integer PRIMARY KEY AUTOINCREMENT, -- 自增ID
    name varchar ( 64 )  NOT NULL, -- 名称
    username varchar ( 64 ) NOT NULL DEFAULT '', -- 用户名
    password varchar ( 64 ) NOT NULL DEFAULT '', -- 用户密码
    role tinyint not null default 0, -- 用户角色
    status tinyint not null default 0, -- 用户状态
    create_time bigint not null, -- 创建时间
    update_time bigint not null, -- 更新时间
    delete_time bigint not null default 0 -- 删除时间
);