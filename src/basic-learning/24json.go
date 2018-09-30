package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B string `json:"b"`
	C int    `json:"c"`
	D struct {
		E float32 `json:"e"`
		F string  `json:"f"`
	} `json:"d"`
}

func main() {
	a := A{
		B: "bbbb",
		C: 20,
		D: struct {
			E float32 `json:"ee"`
			F string  `json:"ff"`
		}{
			E: 0.3,
			F: "ffffffff",
		},
	}
	str, err := json.Marshal(a)
	fmt.Println(err)
	fmt.Println(string(str))
}
