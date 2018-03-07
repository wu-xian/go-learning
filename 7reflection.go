package main

import (
	"fmt"
	"reflect"
)

type man struct {
	name   string
	id     int
	parent *man
}

func main7() {
	var x int = 3
	typeOf := reflect.TypeOf(x)
	valueOf := reflect.ValueOf(x)
	m1 := new(man)
	m1.id = 123
	m1.name = "wuxian"
	m1.parent = nil

	m2 := &man{"heihei", 233, m1}
	fmt.Println("typeof:", typeOf)
	fmt.Println("valueof:", valueOf)

	fmt.Println("typeof man1:", reflect.TypeOf(m1))
	fmt.Println("valueof man1:", reflect.ValueOf(m1))

	fmt.Println("typeof man2:", reflect.TypeOf(m2))
	fmt.Println("valueof man2:", reflect.ValueOf(m2))

	var pi float64 = 3.1415926575323846
	fValue := reflect.ValueOf(&pi)
	fType := reflect.TypeOf(&pi)
	fmt.Println("value", fValue)
	fmt.Println("type", fType)

	fmt.Println("can set:", fValue.CanSet())
	fValue.Elem().SetFloat(2.81)

	fmt.Println("pi:", pi)

}
