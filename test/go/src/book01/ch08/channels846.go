package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "welcome my world"
	fmt.Println(cap(ch)) // 查看初始容量大小
	fmt.Println(len(ch)) // 查看使用容量空间
}
