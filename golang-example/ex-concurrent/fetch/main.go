package main

import (
	"fmt"
	"time"
	"net/http"
	"golang.org/x/net/html"
)

// 并发获取多个URL并统计每个请求的响应时间


// main 主线程

// fetch函数
// 
func main() {
	// 执行时间计算
	// start := time.Now()

	// ch := make(chan string)

	// url := [...]string{"https://www.sina.com.cn", 
	// 					"https://www.baidu.com", 
	// 					"https://www.qq.com", 
	// 					"https://news.163.com",
	// 					"https://www.taobao.com",
	// 					"https://gocn.vip",
	// 					"https://juejin.cn",
	// 					"https://u.geekbang.org",
	// 					"https://www.lottery.gov.cn"}

	// for _, v := range url {

	// 	go func (ch chan<- string) {
			
	// 		startSub := time.Now()

	// 		_, err := http.Get(v)

	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		secs := time.Since(startSub).Seconds()
	// 		ch <- fmt.Sprintf("sub time %.2fs", secs)
	// 	}(ch)
	// }

	// for range url {
	// 	fmt.Println(<-ch)
	// }

	// fmt.Printf("%.2fs 执行时间\n", time.Since(start).Seconds())

	// 结论
	// 使用并发请求前，主线程执行时间约等于请求时间的总和
	// 使用并发请求后，主线程执行时间等于最长那个请求时间

	//=============================================

	getHref()

}

func noCacheGet() {
	// 测试结果没有得到想要的答案

	start := time.Now()

	request, _ := http.NewRequest("GET", "https://segmentfault.com", nil)

	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("Cache-Control", "must-revalidate")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Cache-Control", "no-store")
	request.Header.Set("Expires", "0")

	_, err := (&http.Client{}).Do(request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%.2fs 执行时间\n", time.Since(start).Seconds())
}


func getHref() {

	// var href = []string
	resp, _ := http.Get("https://www.baidu.com")

	
	doc, _ := html.Parse(resp.Body)

	fmt.Println(doc)

	resp.Body.Close()

	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, a := range doc.Attr {
			if a.Key != "href" {
				continue
			}
			fmt.Println(a.Val)
		}
	}

	fmt.Println("fetch html ....")
}







