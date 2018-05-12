package main

import (
	"fmt"
)

func main() {
	var value uint8 = 0x08
	fmt.Printf("%o \n", value)

	fmt.Println("==> %x %d %o")
	a := 0x10
	fmt.Printf("%x \n", a)
	fmt.Printf("%d \n", a)
	fmt.Printf("%o \n", a)
	fmt.Println("=========")

	fmt.Println("==> &   |")
	b := 0x0A                    //0000 1010
	c := 0x01                    //0000 0001
	fmt.Printf("b&c %x \n", b&c) //0000 0000
	fmt.Printf("b|c %x \n", b|c) //0000 1011
	fmt.Printf("^b %x \n", ^b)   //1111 0101
	fmt.Printf("^c %x \n", ^c)   //1111 1110
	fmt.Printf("b^c %x \n", b^c) //0000 1011
	fmt.Println("=========")
}
