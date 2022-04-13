package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 测试抓取链接
	// https://draveness.me/golang/docs/part4-advanced/ch09-stdlib/golang-json/ b.html
	//
	HttpGet2("https://top.chinaz.com/alltop/", "c.html")
}

func HttpGet2(url string, fileName string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页body内容
	f, err2 := os.Create(fileName)
	if err2 != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return
}
