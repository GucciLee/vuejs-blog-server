## 项目预览
* 暂时未放到线上，敬请期待。。。

## 项目简介
* 基于 `go mod` 管理包依赖
* 使用 `beego` 框架开发

## 安装
> 如果你是一个程序员，并且是一个 Golang 新手，请先阅读文章：[Golang go.mod 环境搭建](http://www.learnku.net/blog/articles/147)，再来进行构建此程序。

```
# 拉取项目
git clone git@github.com:learnku/vuejs-blog-server.git

# 拉取包依赖
go mod tidy

# 运行服务器
bee run

# 访问 http://localhost:8888
```


## 工具
```
// 快速生成代码
bee generate scaffold users -fields="id:int64,username:string,email:string,password:string,avatar:string,address:string" -driver=mysql -conn="root:root@tcp(127.0.0.1:3306)/vuejs_blog_server"
```