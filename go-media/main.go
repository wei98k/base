package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//文件上传-http
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	file, handler, err := r.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./uploads/files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
