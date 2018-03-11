package main

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
)

type test struct {
	Id   int    `xorm:"'Id'"`
	Name string `xorm:"'Name'"`
}

func main() {
	t := new(test)
	t.Id = 0
	t.Name = "wuxian" + strconv.Itoa(t.Id)
	engine, _ := xorm.NewEngine("mysql", "root:wuxian@/test")
	engine.Table("test").Insert(&t)
	fmt.Println("ok")
}
