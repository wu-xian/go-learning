package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		waitGroup.Add(i)
		go process(i, waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Done.")
}

func process(i int, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println(i, "WaitDone")
	wg.Done()
}
