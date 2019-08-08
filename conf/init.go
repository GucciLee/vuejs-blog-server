package conf

import (
	"vuejs-blog-server/utils/filepath"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
)

// conf.EnvConfig.String("")
var EnvConfig config.Configer

func init()  {
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