package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main8() {
	var message string = ""
	var args string = ""
	fmt.Println("input two words:")

	args = strings.Join(os.Args, "#")
	file, _ := os.Open("tester.txt")
	defer file.Close()
	fmt.Println(message, "@", args)

	reader := bufio.NewReader(file)
	txt, _ := reader.ReadString('\n')
	fmt.Println("file content:")
	fmt.Println(txt)
}
