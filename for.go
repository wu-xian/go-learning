package main

import "fmt"
import "net/http"

func main1() {
	fmt.Printf("hello go")

	var a int

	for a < 100 {
		fmt.Printf("%d \n", a)
		a++
	}

	for a < 100 {
		fmt.Printf("%d \n", a)
		a++
	}

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("hahahah" + request.URL.String()))
		responseWriter.Write([]byte("Host:" + request.Host))
		responseWriter.Write([]byte("Method" + request.Method))
		responseWriter.Write([]byte("URL.Hostname" + request.URL.Hostname()))
		responseWriter.Write([]byte("URL.RawQuery" + request.URL.RawQuery))
	})

	http.ListenAndServe(":5000", nil)

}
