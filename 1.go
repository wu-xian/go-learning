package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("asdasd")
	fmt.Println("jajajajaja")
	fmt.Print("wulauwuasla")

	myArray := []int{1, 3, 5, 6, 67, 234, 56744, 567, 234, 123}

	for i, a := range myArray {
		fmt.Println(i, a)
	}

	fmt.Println("==================")

	result,sum:=anomousFunc(myArray)

	fmt.Println(result)
	fmt.Println(sum)
}

func anomousFunc(params []int) (results string,sum int) {
	sum1:=0
	result:=""
	for	_,param := range(params){
		sum1+=param
		result+=strconv.Itoa(param)
	}
	results = result
	sum = sum1
	return
}