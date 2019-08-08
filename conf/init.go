package conf

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"vuejs-blog-server/utils/filepath"
	"log"
)

func init()  {
	var env string = "conf/env.conf"
	var envExample string = "conf/env.example.conf"
	// 如果 evn.conf 文件不存在，就创建

	if (!filepath.Exists(env)){
		filepath.CopyFile(env, envExample)
	}

	// 读取 env 配置文件
	envConf, err := config.NewConfig("ini", env)
	if err != nil {
		log.Fatalln("读取 "+ env +" 文件失败")
	}

	fmt.Println(envConf.String("mysql::username"))
}