package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"vuejs-blog-server/conf"
	"github.com/astaxie/beego"
)

func init()  {
	// 获取 env::mysql 配置信息
	mysql := map[string]string{
		"username": conf.EnvConfig.String("mysql::username"),
		"password": conf.EnvConfig.String("mysql::password"),
		"tcp": conf.EnvConfig.String("mysql::tcp"),
		"dbname": conf.EnvConfig.String("mysql::dbname"),
		"charset": conf.EnvConfig.String("mysql::charset"),
	}

	// set default database
	// 连接数据库
	dataSource := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v", mysql["username"], mysql["password"], mysql["tcp"], mysql["dbname"], mysql["charset"])
	orm.RegisterDataBase("default", "mysql", dataSource, 30)

	// register model
	// 在表不存在的时候会创建表
	orm.RegisterModel(new(Users))

	// create table
	// @params name default 代表mysql数据库
	// @params force 强制创建表
	// @params verbose 打印更多信息
	orm.RunSyncdb("default", false, true)

	// 调试模式
	if beego.BConfig.RunMode == "dev" {
		// 打印查询语句
		// orm.Debug = true
	}
}