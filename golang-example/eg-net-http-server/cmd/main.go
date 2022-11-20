package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		//输出请求方式、请求参数
		fmt.Printf("Method: %v\n", req.Method)
		fmt.Printf("Content-Type: %v\n", req.Header.Get("Content-Type"))
		fmt.Printf("Body: %v\n", req.Body)
		fmt.Printf("Form: %v\n", req.Form)
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
