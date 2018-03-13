package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("go web!")
	http.HandleFunc("/", processRequest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func processRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println("start request")
}
