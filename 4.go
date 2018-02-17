package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a)
	c, _ := strconv.ParseInt("99", 10, 64)
	d := strconv.Itoa(a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
}
