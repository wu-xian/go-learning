package main

import (
	"fmt"
)

func main_1 {
	fmt.Println("closure")

	demoFunction2()
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func demoFunction2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
