package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	xorm "github.com/go-xorm/xorm"
)

type user struct {
	Id       string `xorm:"char(36) notnull unique 'Id'"`
	UserName string `xorm:"varchar(50) notnull 'UserName'`
	Password string `xorm:"varchar(50) notnull 'Password'`
	RealName string `xorm:"varchar(50) notnull 'RealName'`
}

func main() {
	fmt.Println("13")
	engine, err := xorm.NewEngine("mysql", "root:wuxian@/lamp")
	if err != nil {
		fmt.Println(err)
	}
	var users = []user{}
	//user.Id = "7efe47a5-4083-4e83-b019-8bd3f053ff24"
	err = engine.Table("user").Select("user.Id,user.UserName,user.Password,user.RealName").Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Println(user.Id, "|", user.UserName, "|", user.Password, "|", user.RealName)
	}

}
