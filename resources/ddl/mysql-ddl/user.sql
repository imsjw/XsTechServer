create table `user` (
    `id` int unsigned primary key auto_increment comment '用户ID',
    `username` varchar(32) not null unique comment '用户名',
    `password` varchar(32) not null comment '用户的登陆密码',
    `create_time` bigint unsigned not null comment '创建时间',
    `create_user` int unsigned not null comment '创建用户',
    `update_time` bigint unsigned not null comment '更新时间',
    `update_user` int unsigned not null comment '更新用户'
) charset utf8mb4 collate utf8mb4_unicode_ci engine=innodb comment '用户表';