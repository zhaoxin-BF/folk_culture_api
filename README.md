# folk_culture_api 民俗文化资源管理系统 后台


## 说明：
* 采用golang语言，gin web 框架来进行的项目开发
* 采用go Mudles 进行包管理工具
* 数据库采用远程云端数据库Mysql

## 数据库设计代码
```
# 民俗文化资源管理系统 folk culture：民俗文化
drop database if exists folk_culture_system;
create database folk_culture_system;
use folk_culture_system;
drop table if exists users_table;                     #用户信息账户表
create table users_table(
    user_id int not null primary key auto_increment,  #1、用户ID   区分唯一的
    user_name varchar(50) ,                           #2、用户名字
    user_account varchar(11) not null unique ,        #3、用户账户  账户注册唯一
    user_password varchar(11),                        #4、用户密码
    user_tel varchar(11) not null,                    #5、用户电话
    user_type int,                                    #6、用户类型 0普通用户 1管理员 2超级管理员
    create_time int,                                  #7、用户注册时间
    user_context varchar(200)                         #8、备注信息
);

drop table if exists tags_table;                      #分类标签表
create table tags_table(
    tag_id int not null primary key auto_increment,   #1、类别id
    tag_name varchar(50),                             #2、类别名字
    tag_size int,                                     #3、类别资源数量
    tag_context varchar(200)                          #4、备注信息
);

drop table if exists resources_table;                 #资源信息表
create table resources_table(
    resource_id int not null primary key auto_increment, #1、资源id 自增
    tag_id int ,                                         #2、分类标签
    tag_name varchar(50),                                #3、分类的名字
    resource_name varchar(100),                          #4、资源名称
    description varchar(200),                            #5、资源描述
    author varchar(40),                                  #6、资源作者
    time varchar(20),                                    #7、资源年代
    nation varchar(40),                                  #8、资源民族
    region varchar(50),                                  #9、资源所属地域
    copyright varchar(100),                              #10、资源所属版权
    url       varchar(100),                              #11、资源连接
    status int ,                                         #12、资源状态 0 表示审核通过， 1 表示审核中 2 表示未通过审核

    create_time int,                                     #13、创建时间int
    screate_time varchar(50),                            #14、创建时间string

    update_time int,                                     #15、更新时间int
    supdate_time varchar(50),                            #16、更新时间string

    resource_context text,                               #17、资源备注信息

    upload_id int,                                       #18、上传人id
    upload_user varchar(50),                             #19、上传人的姓名
    check_id int  ,                                       #20、审核者id
    check_name varchar(50)                                #21、审核者名字
);
```
