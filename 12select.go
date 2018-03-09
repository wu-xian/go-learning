package main

import (
	"fmt"
	"time"
)

var i = 0

func main() {
	ch1 := make(chan string)
	//ch2 := make(chan string)

	go func() {
		fmt.Println("started anonymous func")
		time.Sleep(2 * time.Second)
		ch1 <- "in anonymous"
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case c := <-ch1:
			{
				fmt.Println("get it ", c, i)
				return
			}
		default:
			{
				fmt.Println("can not get anymore", i)
				i++
			}
		}
	}
	fmt.Println("invalid return")
}
