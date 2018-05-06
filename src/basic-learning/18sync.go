package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main18() {
	locker := new(sync.RWMutex)
	i := 1
	for j := 0; j < 100; j++ {
		go func() {
			locker.Lock()
			i++
			fmt.Println(strconv.Itoa(i))
			locker.Unlock()
		}()
	}
	fmt.Println("==", strconv.Itoa(i))
	var ss = ""

	fmt.Println(ss)
}
