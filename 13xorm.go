package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	xorm "github.com/go-xorm/xorm"
)

type Test struct {
	Id   int64  `xorm:"'Id'"`
	Name string `xorm:"'Name'"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:1234@192.168.1.199:3306/test")
	if err != nil {
		fmt.Println(err)
	}

	t := Test{}
	t.Id = 1
	result, err := engine.Table("test").Get(&t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
