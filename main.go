package main

import (
	_ "vuejs-blog-server/logs"
	_ "vuejs-blog-server/conf"
	_ "vuejs-blog-server/routers"
	_ "vuejs-blog-server/models"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
