package log

import "github.com/astaxie/beego/logs"

var Logger *logs.BeeLogger

func InitLogger() {
	Logger = logs.NewLogger(1000)
	//Logger.SetLogger("console", "")
	Logger.SetLogger("file", `{"filename":"test.log"}`)
}
