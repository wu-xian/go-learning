package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main4() {
	var a int = 65
	b := string(a)
	c, _ := strconv.ParseInt("99", 10, 64)
	d := strconv.Itoa(a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)

	maper1 := make(map[string]string)
	maper1["talker"] = "balabalabala..."

	ss := "http://localhost:9090"
	idx := strings.LastIndex(ss, ":")
	fmt.Println(idx)
	fmt.Println(ss[:idx])
	fmt.Println(ss[idx:])
	fmt.Println(maper1["talker"])
}
