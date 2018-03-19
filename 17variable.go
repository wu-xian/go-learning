package main

import (
	"fmt"
)

//X this is X
type X string
type action func(message string, integer int)

var Message string

func main() {
	var x X = "hello"
	fmt.Println(x)
	var action1 action = func(msg string, i int) { fmt.Println("in sub func", msg, i) }
	action1("hey", 23)
	Message = "message!!!!"
	fmt.Println(Message)
}
