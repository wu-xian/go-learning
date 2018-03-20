package main

import (
	"fmt"
)

//X this is X
type X string
type action func(message string, integer int, f1 func(string) string)

var Message string

func main17() {
	var x X = "hello"
	fmt.Println(x)
	f2 := func(s string) string {
		fmt.Println("in sub sub func")
		fmt.Println("s", s)
		return "string sub func return"
	}
	var action1 action = func(msg string, i int, ff func(string) string) {
		fmt.Println("in sub func", msg, i)
		f := ff(msg)
		fmt.Println(f)
	}
	action1("hey", -999, f2)
	Message = "message!!!!"
	fmt.Println(Message)
	s1 := "asdddf"
	fmt.Println(X(s1))
}
