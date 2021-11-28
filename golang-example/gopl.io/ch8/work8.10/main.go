package main

import (
	"example/gopl.io/ch8/work8.10/links"
	"fmt"
	"log"
	"os"
)

// 练习 8.10： HTTP请求可能会因http.Request结构体中
// Cancel channel的关闭而取消。
// 修改8.6节中的web crawler来支持取消http请求。
// （提示：http.Get并没有提供方便地定制一个请求的方法。
// 你可以用http.NewRequest来取而代之，
// 设置它的Cancel字段，
// 然后用http.DefaultClient.Do(req)来进行这个http请求。）

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	cancelled := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancelled)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancelled)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string, cancelled chan struct{}) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancelled)
	if err != nil {
		log.Print(err)
	}
	return list
}
