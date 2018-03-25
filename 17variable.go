package main

import (
	"fmt"
)

//X this is X
type X string
type action func(message string, integer int, f1 func(string))

var Message string

func main17() {
	var x X = "hello"
	fmt.Println(x)
	f2 := func(s string) {
		fmt.Println("in sub sub func")
		fmt.Println("s", s)
	}
	var action1 action = func(msg string, i int, ff func(string)) {
		fmt.Println("in sub func", msg, i)
		ff(msg)
	}
	action1("hey", -999, f2)
	Message = "message!!!!"
	fmt.Println(Message)
	s1 := "asdddf"
	fmt.Println(X(s1))
}
