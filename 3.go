package main

import (
	"fmt"
)

func main() {
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
}
