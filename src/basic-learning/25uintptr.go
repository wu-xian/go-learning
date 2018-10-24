package main

import (
	"fmt"
)

func main() {
	const PtrSize = 4 << (^uintptr(0) >> 63)
	fmt.Println(PtrSize)
	fmt.Println(^uintptr(0))
	fmt.Println(^uintptr(1))
	fmt.Println(uint(0x01 << 63))
	fmt.Println(uint(0xffffffffffffffff))
	fmt.Println(uint(0xffffffff))
}
