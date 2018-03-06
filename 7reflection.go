package main

import (
	"fmt"
	"reflect"
)

func main7() {
	var pi float64 = 3.1415926575323846
	fValue := reflect.ValueOf(&pi)
	fType := reflect.TypeOf(&pi)
	fmt.Println("value", fValue)
	fmt.Println("type", fType)

	fmt.Println("can set:", fValue.CanSet())
	fValue.Elem().SetFloat(2.81)

	fmt.Println("pi:", pi)
}
