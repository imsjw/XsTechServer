create table `oauth_auth` (
    `id` int unsigned primary key auto_increment comment '唯一标识符',
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
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '授权表';

create table `oauth_role` (
    `id` int unsigned primary key auto_increment comment '唯一标识符',
    `name` varchar(32) not null comment '角色名称',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '角色表';

create table `oauth_resource` (
    `id` int unsigned primary key auto_increment comment '唯一标识符',
    `method` varchar(32) not null comment '访问URL',
    `url` varchar(32) not null comment '访问URL',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '资源表';

create table `oauth_role_resource` (
    `id` int unsigned primary key auto_increment comment '唯一标识符',
    `role_id` int unsigned not null comment '角色ID',
    `resource_id` int unsigned not null comment '资源ID',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '角色资源关联表';

create table `oauth_user_role` (
    `id` int unsigned primary key auto_increment comment '唯一标识符',
    `user_id` int unsigned not null comment '用户ID',
    `role_id` int unsigned not null comment '角色ID',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '用户角色关联表';