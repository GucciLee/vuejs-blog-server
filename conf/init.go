package conf

import (
	"os"
	"vuejs-blog-server/utils/filepath"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
)
func exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func init()  {
	var env string = "conf/env.conf"
	var envExample string = "conf/env.example.conf"

	// 如果 evn.conf 文件不存在，就创建
	if (!filepath.Exists(env)){
		filepath.CopyFile(envExample, env)
		beego.Error("文件：" + env + "创建成功，请配置它已让程序正常运行。。。")
	}

	// 读取 env 配置文件
	envConf, err := config.NewConfig("ini", env)
	if err != nil {
		beego.Error(err)
	}
	beego.Warn(envConf.String("mysql::username"))
}