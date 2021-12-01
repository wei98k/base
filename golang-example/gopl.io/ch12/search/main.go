package main

import (
	"example/gopl.io/ch12/params"
	"fmt"
	"log"
	"net/http"
)

// Unpack将请求参数填充到合适的结构体成员中，
// 这样我们可以方便地通过合适的类型类来访问这些参数。

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels    []string `http:"l"`
		MaxResult int      `http:"max"`
		Exact     bool     `http:"x"`
	}
	data.MaxResult = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
