create table sys_user
(
    `id`       bigint auto_increment primary key,
    `username` varchar(32)  not null comment '用户名',
    `pwd`      varchar(64)  not null comment '密码',
    `salt`     varchar(64)  not null comment '密码盐',
    `email`    varchar(64)  not null default '' comment '邮箱',
    `nickname` varchar(32)  not null comment '昵称',
    `avatar`   varchar(256) not null default '' comment '头像',
    `phone`    varchar(32)  not null default '' comment '手机号',
    `status`   varchar(32)  not null default 'normal' comment '用户状态',
    `utime`    timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`    timestamp    not null default current_timestamp comment '创建时间',
    unique uk_u (`username`),
    index idx_u_s (`username`, `status`)
)
    engine = innodb
    default charset = utf8mb4 comment '用户信息表';

create table sys_config
(
    `id`         bigint auto_increment primary key,
    `group`      varchar(32)           default '' not null comment '分组',
    `group_name` varchar(32)           default '' not null comment '分组名称',
    `key`        varchar(64)  not null comment '配置键',
    `value`      text         not null comment '配置值',
    `status`     varchar(16)  not null comment 'normal disable del',
    `remark`     varchar(512) not null comment '备注',
    `utime`      timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`      timestamp    not null default current_timestamp comment '创建时间',
    unique uk_k (`key`)
)
    engine = innodb
    default charset = utf8mb4 comment '系统配置';

create table sys_dict
(
    `id`         bigint auto_increment primary key,
    `group`      varchar(32)  not null comment '分组',
    `group_name` varchar(32)  not null comment '分组名称',
    `value`      varchar(64)  not null comment '数据值',
    `label`      varchar(128) not null comment '显示标签',
    `status`     varchar(16)  not null comment 'normal disable del',
    `remark`     varchar(512) not null default '' comment '备注',
    `utime`      timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`      timestamp    not null default current_timestamp comment '创建时间',
    unique uk_g_v (`group`, `value`)
)
    engine = innodb
    default charset = utf8mb4 comment '系统字典表';



