package main

import (
	"example/gopl.io/ch5/links"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// 练习 8.7： 完成一个并发程序来创建一个线上网站的本地镜像，
// 把该站点的所有可达的页面都抓取到本地硬盘。
// 为了省事，我们这里可以只取出现在该域下的所有页面
// （比如golang.org开头，译注：外链的应该就不算了。）
// 当然了，出现在页面里的链接你也需要进行一些处理，
// 使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。

// go run main.go -base https://books.studygolang.com/

var (
	base = flag.String("base", "https://books.studygolang.com/", "")
)

var wg sync.WaitGroup

func main() {
	flag.Parse()
	for _, url := range crawl(*base) {
		wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			download(*base, url)
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
}

func download(base, url string) {
	if !strings.HasPrefix(url, base) {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	dir := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}

	filename := dir + "index.html"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
