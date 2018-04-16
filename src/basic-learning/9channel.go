package main

import (
	"fmt"
)

func main9() {
	fmt.Println("9.channel")
	interger := make(chan int, 2)

	go readChan(interger)
	fmt.Println("send channel before")
	interger <- 99
	interger <- 98
	fmt.Println("end of line")
	for {
		msg, _ := fmt.Scanln()
		interger <- msg
	}
}

func readChan(ch <-chan int) {
	var result int
	var count int = 0
	for {
		result = <-ch
		fmt.Println("read chan ", count, result)
		count++
		if result != 0 {
			fmt.Println(result)
		}
	}
	fmt.Println("gorutine is closed")
}
