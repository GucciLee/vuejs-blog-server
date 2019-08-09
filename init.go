package main

import (
	"vuejs-blog-server/utils/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/validation"
	"vuejs-blog-server/models"
)

func init()  {
	initConf()
	initLogs()
	initModels()
	initRequest()
}

// 1. 配置文件初始化
// EnvConfig.String("")
var EnvConfig config.Configer
func initConf()  {
	var env string = "conf/env.conf"
	var envExample string = "conf/env.example.conf"

	// 如果 evn.conf 文件不存在，就创建
	if (!filepath.Exists(env)){
		filepath.CopyFile(envExample, env)
		beego.Error("文件：" + env + "创建成功，请配置它已让程序正常运行。。。")
	}

	// 读取 env 配置文件
	var err error
	EnvConfig, err = config.NewConfig("ini", env)
	if err != nil {
		beego.Error(err)
	}
}

// 2. 日志系统初始化
func initLogs()  {
	// github.com/astaxie/beego/logs/log.go ==> line280
	// writeMsg 方法中
	// `_, filename := path.Split(file)` 修正为 `filename := file`

	// 设置日志文件
	beego.SetLogger("file", `{"filename":"logs/logs.log"}`)

	// 只输出到文件
	beego.BeeLogger.DelLogger("console")

	// 设置级别
	beego.SetLevel(beego.LevelNotice)

	// 输出文件名和行号
	beego.SetLogFuncCall(true)
}

// 3. 数据库初始化
func initModels()  {
	// 获取 env::mysql 配置信息
	mysql := map[string]string{
		"username": EnvConfig.String("mysql::username"),
		"password": EnvConfig.String("mysql::password"),
		"tcp": EnvConfig.String("mysql::tcp"),
		"dbname": EnvConfig.String("mysql::dbname"),
		"charset": EnvConfig.String("mysql::charset"),
	}

	// set default database
	// 连接数据库
	dataSource := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v", mysql["username"], mysql["password"], mysql["tcp"], mysql["dbname"], mysql["charset"])
	orm.RegisterDataBase("default", "mysql", dataSource, 30)

	// register model
	// 在表不存在的时候会创建表
	registeModels()

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

// 3.1 在表不存在的时候会创建表
func registeModels()  {
	orm.RegisterModel(new(models.Users))
}

// 4. 表单验证初始化
func initRequest()  {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "不能为空",
		"Min":          "Minimum is %d",
		"Max":          "Maximum is %d",
		"Range":        "Range is %d to %d",
		"MinSize":      "Minimum size is %d",
		"MaxSize":      "Maximum size is %d",
		"Length":       "Required length is %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	})
}