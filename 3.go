package main

import (
	"fmt"
)

func main3() {
	fmt.Println("hello go")
	var slice1 = make([]int, 0)
	var ptrSlice = &slice1
	fmt.Println(&ptrSlice)
	_ = append(slice1, 1)
	fmt.Println(ptrSlice)
	_ = append(slice1, 3)
	fmt.Println(ptrSlice)
	_ = append(slice1, 5)

	fmt.Println("====================")
	for _, item := range slice1 {
		fmt.Print(item)
		fmt.Print("  ")
	}

	fmt.Println("***********************")

LOOP:
	for i := 0; i < 10; i++ {
		if i < 5 {
			fmt.Println(i)
		} else {
			break LOOP
		}
	}

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer in sub func i = ", i)
		}()
		fs[i] = func() { fmt.Println("closuere i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}
