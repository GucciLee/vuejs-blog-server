## 项目预览
* 暂时未放到线上，敬请期待。。。

## 项目简介
* 基于 `go mod` 管理包依赖
* 使用 `beego` 框架开发

## 安装
```
# 拉取项目
git clone git@github.com:learnku/vuejs-blog-server.git

# 初始化一个 module 名为 vuejs-blog-server 【这一步可以跳过】 
# go mod init vuejs-blog-server

# 拉取包依赖
go mod tidy

# 运行服务器
bee run
```


## 工具
```
// 快速生成代码
bee generate scaffold users -fields="id:int64,username:string,email:string,password:string,avatar:string,address:string" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/vuejs_blog_server"
```