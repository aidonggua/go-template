create database if not exists go_template default charset utf8mb4 collate utf8mb4_unicode_ci;
use go_template;
create table if not exists user
(
    id           bigint auto_increment   not null primary key,
    account      varchar(64)             not null comment '账号',
    password     varchar(255)            not null comment '密码',
    name         varchar(255) default '' not null comment '姓名',
    gmt_create   timestamp               not null comment '创建时间',
    gmt_modified timestamp               not null comment '更新时间',
    deleted      tinyint(1)   default 0  not null comment '0:未删除 1:已删除',
    unique unique_account (account)
) comment '用户表';