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

// 练习5.13： 修改crawl，使其能保存发现的页面，
// 必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。
// 假设初始页面在golang.org下，
// 就不要保存vimeo.com下的页面。

var (
	base = flag.String("base", "https://www.sulinehk.com", "")
)

var wg sync.WaitGroup

func main() {
	flag.Parse()
	for _, url := range crawl(*base) {
		wg.Add(1)
		go download(*base, url)
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()
	<-done
}

func download(base, url string) {
	defer wg.Done()

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
