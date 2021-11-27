package main

import (
	"example/gopl.io/ch5/links"
	"fmt"
	"log"
	"os"
)

// 让我们重写crawl函数，
// 将对links.Extract的调用操作用获取、
// 释放token的操作包裹起来，
// 来确保同一时间对其只有20个调用。
// 信号量数量和其能操作的IO资源数量应保持接近。

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}
