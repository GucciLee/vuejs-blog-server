package logs

import (
	"github.com/astaxie/beego"
)


func init(){
	// github.com/astaxie/beego/logs/log.go ==> line280
	// writeMsg 方法中
	// `_, filename := path.Split(file)` 修正为 `filename := file`

	// 设置日志文件
	beego.SetLogger("file", `{"filename":"logs/logs.log"}`)

	// 只输出到文件
	beego.BeeLogger.DelLogger("console")

	// 设置级别
	// beego.SetLevel(beego.LevelNotice)

	// 输出文件名和行号
	beego.SetLogFuncCall(true)
}