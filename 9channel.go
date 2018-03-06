package main

import (
	"fmt"
)

func main() {
	fmt.Println("9.channel")
	interger := make(chan int)

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

func readChan(ch chan int) {
	var result int
	for {
		result = <-ch
		if result == 0 {
			break
		}
		fmt.Println(result)
	}
}
