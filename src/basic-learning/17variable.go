package main

import (
	"fmt"
)

//X this is X
type X string
type action func(message string, integer int, f1 func(string) string)

var Message string

func main() {
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

	fmt.Println("====================")

	a := 1
	b := 2
	a, b = func() int {
		fmt.Println("first func , a , b = ", a, b)
		return b
	}(), func() int {
		fmt.Println("second func , a , b = ", a, b)
		return a
	}()
	fmt.Println("a,b", a, b)

	c := 3
	d := 4
	c, d = d, c
	fmt.Println("c,d", c, d)
}
