package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

var logger *logs.BeeLogger

func main() {
	logger = logs.NewLogger(1000)
	logger.SetLogger("file", `{"filename":"test.log"}`)
	fmt.Println("started")
	logger.Info("log test")
}
