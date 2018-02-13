package main

import (
	"fmt"
	"learn/packages"
)

type smoker struct {
	id    int
	name  string
	inUse bool
}

func main2() {
	i := 1
	var ptr *int
	var smoker1 = smoker{1, "wu-xian", true}
	var smoker2 = smoker1

	fmt.Println(smoker1.id, smoker1.name, smoker1.inUse)
	fmt.Println("===================")
	smoker2.id = 3
	fmt.Println(smoker1.id, smoker1.name, smoker1.inUse)
	fmt.Println(smoker2.id, smoker2.name, smoker2.inUse)

	ptr = &i
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	packages.Add(1, 2)
	if i == 1 {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}
	fmt.Println("ptr", ptr)
	fmt.Println(packages.PI)
}
