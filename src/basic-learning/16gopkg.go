package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func main16() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Println(err)
	}
	server := cfg.Section("server")
	serverName := server.Key("Name").MustString("haha")
	fmt.Println(serverName)

}
