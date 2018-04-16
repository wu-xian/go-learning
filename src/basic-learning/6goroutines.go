package main

import (
	"fmt"
)

func main6() {
	go printString("what happened")

	for i := 0; i < 10; i++ {
		fmt.Println("in main", i)
	}
}

func printString(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println("in goroutine", message)
	}
}
