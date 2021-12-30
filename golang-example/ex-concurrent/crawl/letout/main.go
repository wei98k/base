package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func main() {
	// num := 6
	// for index := 0; index < num; index++ {
	// 	resp, _ := http.Get("https://www.baidu.com")
	// 	_, _ = ioutil.ReadAll(resp.Body)
	// 	resp.Body.Close()
	// }
	resp, _ := http.Get("https://www.baidu.com")
	_, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	resp1, _ := http.Get("https://news.qq.com/")
	_, _ = ioutil.ReadAll(resp1.Body)
	resp1.Body.Close()

	fmt.Printf("now num goroutine: %d\n", runtime.NumGoroutine())
}
