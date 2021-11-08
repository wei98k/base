package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 练习 1.10： 找一个数据量比较大的网站，
// 用本小节中的程序调研网站的缓存策略，
// 对每个URL执行两遍请求，查看两次时间是否有较大的差别，
// 并且每次获取到的响应内容是否一致，修改本节中的程序，
// 将响应结果输出，以便于进行对比。

// go run main.go https://news.sina.com.cn https://news.qq.com
// url: https://news.qq.com
// resp1: 0.32s    6510
// resp2: 0.02s    6510
// url: https://news.sina.com.cn
// resp1: 0.34s  481662
// resp2: 0.02s  481662

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchTwoTimes(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetchTwoTimes(url string, ch chan<- string) {
	ch <- "url: " + url + "\nresp1: " + fetch(url) + "\n" + "resp2: " + fetch(url)
}

func fetch(url string) string {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprint(err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs %7d", secs, nbytes)
}
