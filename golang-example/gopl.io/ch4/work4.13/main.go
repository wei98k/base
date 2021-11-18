package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// 练习 4.13： 使用开放电影数据库的JSON服务接口，
// 允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。
// 编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。

// 思路
// 1. 搜索URL的响应内容
// 2. 匹配的话返回其海报的图片地址 然后复制文件流生成文件

var (
	title  = flag.String("t", "", "")
	apikey = flag.String("apikey", "", "")
)

const api = "https://www.omdbapi.com/"

type result struct {
	PosterURL string `json:"Poster"`
}

func main() {
	flag.Parse()
	url := api + "?t=" + url.QueryEscape(*title) + "&apikey=" + url.QueryEscape(*apikey)
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get url: %v fail. err: %v", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var res result
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Fprintf(os.Stderr, "json decode fail. err: %v", err)
		os.Exit(1)
	}

	poster, err := http.Get(res.PosterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http get poster url: %v fail. err: %v", res.PosterURL, err)
		os.Exit(1)
	}
	defer poster.Body.Close()

	f, err := os.Create(*title + ".jpeg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file fail. err: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = io.Copy(f, poster.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io copy fail. err: %v", err)
		os.Exit(1)
	}
}
