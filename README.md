# ByteDanceCamp_tiktok
# 一、项目介绍
## 功能
- 视频推送、视频投稿、发布列表、用户注册、用户登录、点赞、点赞列表、评论、评论列表、关注、关注列表、粉丝列表
## 特点
基于apifox提供的接口进行开发，使用青训营提供的APK进行Demo测试， 功能完整实现，接口匹配良好。

代码结构采用 MVC （Controller,Dao,Service）项目结构清晰，代码符合规范

使用 JWT 进行用户token的校验

使用 Gorm 对 MySQL 进行 ORM 操作；

## 使用提示
clone下来之后修改servce/conf的PlayURL为本机ip地址，修改dao/init_db的端口，user，password 以及创建表或者修改表名（该项目的表名为douyin）
```git
git clone git@github.com:paopaoshuaige/ByteDanceCamp_tiktok.git
go mod tidy
./main.go 
```

## 鸣谢
<a href="https://youthcamp.bytedance.com/" target="_blank">字节跳动青训营</a>
