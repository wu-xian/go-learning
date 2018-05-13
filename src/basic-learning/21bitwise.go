package main

import (
	"fmt"
)

func main21() {
	var value uint8 = 0x08
	fmt.Printf("%o \n", value)

	fmt.Println("==> %x %d %o %b")
	a := 0x10
	fmt.Printf("%%x %x \n", a)
	fmt.Printf("%%d %d \n", a)
	fmt.Printf("%%o %o \n", a)
	fmt.Printf("%%b %b \n", a)
	fmt.Printf("%%08b %08b \n", a)
	fmt.Println("=========")

	fmt.Println("==> &   |   ^")
	b := 0x0A                                      //0000 1010
	c := 0x01                                      //0000 0001
	fmt.Printf("%08b %08b = %08b \n", b, c, b&c)   //0000 0000
	fmt.Printf("%08b | %08b =%08b \n", b, c, b|c)  //0000 1011
	fmt.Printf("^%08b = %08b \n", b, ^b)           //1111 0101
	fmt.Printf("^%08b = %08b \n", c, ^c)           //1111 1110
	fmt.Printf("%08b ^ %08b = %08b \n", b, c, b^c) //0000 1011
	fmt.Println("=========")

	var d uint8 = 2
	fmt.Printf("%08b\n", ^d)

	fmt.Println("==> byte type")
	var e byte = 0xAA
	fmt.Printf("%08b \n", e)
	fmt.Println("=========")
}
