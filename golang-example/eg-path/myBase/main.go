package main

import (
	"fmt"
	"path"
)

func main() {
	url := "https://books.studygolang.com/gopl-zh/ch0/ch0-01.html"
	n := path.Base(url)
	fmt.Println(n)
}
