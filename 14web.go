package main

import (
	"fmt"
	"net/http"
)

func main14() {
	fmt.Println("go web!")
	http.HandleFunc("/", processRequest)
	http.HandleFunc("/hello", processRequestHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func processRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println("start request")
}

func processRequestHello(r http.ResponseWriter, req *http.Request) {
	fmt.Fprint(r, "hello ")
}
