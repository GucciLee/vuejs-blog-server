package models

import (
		"github.com/astaxie/beego/orm"
		_ "github.com/go-sql-driver/mysql"
	)

func init()  {
	// set default database
	// 连接数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/vuejs_blog_server?charset=utf8", 30)

	// register model
	// 在表不存在的时候会创建表
	orm.RegisterModel(new(Users))

	// create table
	// @params name default 代表mysql数据库
	// @params force 打印更多信息
	// @params verbose 强制创建表
	orm.RunSyncdb("default", true, true)
}