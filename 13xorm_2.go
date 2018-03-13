package main

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
)

type test2 struct {
	Id   int    `xorm:"'Id'"`
	Name string `xorm:"'Name'"`
}

func main13_2() {
	t := new(test2)
	t.Id = 0
	t.Name = "wuxian" + strconv.Itoa(t.Id)
	engine, _ := xorm.NewEngine("mysql", "root:wuxian@/test")
	engine.Table("test2").Insert(&t)
	fmt.Println("ok")
}
