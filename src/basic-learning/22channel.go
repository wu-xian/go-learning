package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10)

	go func() {
		defer fmt.Println("Done?")
		for m := range ch {
			defer fmt.Println("defer:", m) //won't work
			//fmt.Println(m, ok)
			time.Sleep(1 * time.Second)
			fmt.Println("processed:", m)
		}
		fmt.Println("Done.")
	}()

	ch <- "cmd.1"
	time.Sleep(2 * time.Second)
	close(ch)
	ch <- "cmd.2" //won't be processed sometimes

	time.Sleep(4 * time.Second)
}
