package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	xorm "github.com/go-xorm/xorm"
)

type user struct {
	Id       string `xorm:"'Id'"`
	UserName string `xorm:"'UserName'`
	Password string `xorm:"'Password'`
	RealName string `xorm:"'RealName'`
}

func main13() {
	fmt.Println("13")
	engine, err := xorm.NewEngine("mysql", "root:wuxian@/lamp")
	if err != nil {
		fmt.Println(err)
	}
	var users = make([]*user, 0)
	//user.Id = "7efe47a5-4083-4e83-b019-8bd3f053ff24"
	err = engine.Table("user").Select("user.Id,user.UserName,user.Password,user.RealName").Cols("Id", "UserName", "Password", "RealName").Find(&users)
	// sess, _ := engine.Query(`select
	// 		user.Id as Id,
	// 		user.UserName as UserName,
	// 		user.Password as Password,
	// 		user.RealName as RealName
	// 	 from user`).Cols("").Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users {
		fmt.Println(user.Id, "|", user.UserName, "|", user.Password, "|", user.RealName)
	}

}
