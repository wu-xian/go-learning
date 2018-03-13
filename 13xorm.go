package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	xorm "github.com/go-xorm/xorm"
)

type test struct {
	Id    int64  `xorm:"'Id'"`
	Name  string `xorm:"'Name'"`
	Value string `xorm:"'Value'"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:1234@tcp(192.168.1.199:3306)/test")
	if err != nil {
		fmt.Println(err)
	}

	t := test{}
	t.Id = 1
	result, err := engine.Table("test").Get(&t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(t)

	//ts := make([]*test, 0)
	var ts []test
	engine.Find(&ts)
	for _, t1 := range ts {
		fmt.Println("#", t1)
	}

	t1 := test{
		Id:    0,
		Name:  "asd",
		Value: "999",
	}
	t2 := test{
		Id:    0,
		Name:  "2222222",
		Value: "9222222222222222222299",
	}
	t3 := test{
		Id:    0,
		Name:  "33333333333333",
		Value: "993333333339",
	}
	sess := engine.NewSession()
	tts := []test{t1, t2, t3}
	for _, _t := range tts {
		sess.Insert(_t)
	}
	sess.Commit()
	defer sess.Close()
>>>>>>> d42bc28a14e9e126652bfa372480008e6aa50a5c
}
