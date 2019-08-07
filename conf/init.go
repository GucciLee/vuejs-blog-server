package conf

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"vuejs-blog-server/utils"
	"io/ioutil"
	"os"
	"log"
)

func init()  {
	var env string = "conf/env.conf"
	var envExample string = "conf/env.example.conf"
	// 如果 evn.conf 文件不存在，就创建
	if (!utils.Exists(env)){
		data, err := ioutil.ReadFile(envExample)
		if err != nil {
			log.Fatalln("读取 "+ envExample +" 文件失败")
		}
		err = ioutil.WriteFile(env, data, os.ModePerm)
		if err != nil {
			log.Fatalln("拷贝 "+ env +" 文件失败")
		}
	}

	// 读取 env 配置文件
	envConf, err := config.NewConfig("ini", "conf/env.conf")
	if err != nil {
		log.Fatalln("读取 "+ env +" 文件失败")
	}

	fmt.Println(envConf.String("mysql::username"))
}