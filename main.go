package main

import (
	"github.com/astaxie/beego/plugins/cors"
	_ "vuejs-blog-server/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		// beego.BConfig.WebConfig.DirectoryIndex = true
		// beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		// https://godoc.org/github.com/astaxie/beego/plugins/auth
		// https://godoc.org/github.com/astaxie/beego/plugins/cors
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			// 允许跨域
			AllowOrigins: []string{"http://localhost:8080"},
		}))
	}

	beego.Run()
}