-- 添加初始用户 账号:admin 密码:password
insert into user (id,username,password,create_time,update_time,create_user,update_user) values (1,'admin','d0a4fd3a28beeb8804051095bc319584',0,0,0,0);

-- 初始化角色表
insert into oauth_role (id,name,create_time,create_user,update_time,update_user) values (1,'超级管理员',0,0,0,0);

-- 初始化资源表
insert into oauth_resource (id,client,method,url,create_time,create_user,update_time,update_user) values (1,'超级管理员',0,0,0,0);