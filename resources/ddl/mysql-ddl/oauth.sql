create table `oauth` (
    `id` int unsigned primary key auto_increment comment '用户ID',
    `user_id` int unsigned not null comment '用户ID',
    `client` varchar(32) not null comment '客户端类型',
    `access_token` text comment 'access_token',
    `access_token_expires_time` bigint unsigned not null default 0 comment 'access_token 过期时间',
    `refresh_token` text comment 'refresh_token',
    `refresh_token_expires_time` bigint unsigned not null default 0 comment 'refresh_token 过期时间',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment 'oauth';