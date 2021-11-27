package main

import (
	"example/gopl.io/ch5/links"
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

// 练习 8.6： 为并发爬虫增加深度限制。也就是说，
// 如果用户设置了depth=3，
// 那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。

type work struct {
	url   string
	depth int
}

func main() {
	worklist := make(chan []work)
	var n int
	n++
	go func() {
		var works []work
		for _, url := range os.Args[1:] {
			works = append(works, work{url, 1})
		}
		worklist <- works
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		works := <-worklist
		for _, w := range works {
			if !seen[w.url] {
				seen[w.url] = true
				n++
				go func(w work) {
					worklist <- crawl(w)
				}(w)
			}
		}
	}
}

func crawl(w work) []work {
	fmt.Printf("depth: %d, url: %s\n", w.depth, w.url)

	if w.depth >= 3 {
		return nil
	}

	tokens <- struct{}{}
	urls, err := links.Extract(w.url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	var works []work
	for _, url := range urls {
		works = append(works, work{url, w.depth + 1})
	}
	return works
}
