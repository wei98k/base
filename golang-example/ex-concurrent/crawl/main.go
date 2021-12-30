package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var (
	baseUrl  = "https://books.studygolang.com/gopl-zh/"
	basePath = "out/"
	baseDir  = ""
)

var tokens = make(chan struct{}, 20)

// crawl gitbook
func main() {

	var wg sync.WaitGroup

	//read base url
	resp, err := http.Get(baseUrl)

	hostname := resp.Request.Host
	baseDir = basePath + hostname + "/"

	// 创建抓取地址目录
	os.MkdirAll(baseDir, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 解析第一个页面
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

	// 保存第一个页面出现的资源文件
	downAssets(doc)

	start := time.Now()

	// 定义用于保存链接的变量
	var links []string

	// 递归遍历第一个页面的节点信息
	visitNode := func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {

			for _, a := range n.Attr {

				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}

				if link.Hostname() == "books.studygolang.com" {

					split := bytes.Split([]byte(a.Val), []byte("/"))
					if len(split) == 2 {
						if string(split[1]) != "" {

							// 创建子目录 不管目录存不存在都执行创建目录操作
							subDir := baseDir + string(split[0])
							os.Mkdir(subDir, fs.ModePerm)
							f := subDir + "/" + string(split[1])

							// HttpGet2(link.String(), f)
							// 直接加go是无法下载到文件的
							// go HttpGet2(link.String(), f)
							// 效果同上
							// go func(l string, f string) {
							// 	HttpGet2(l, f)
							// }(link.String(), f)
							// 执行时间略快一点
							// wg.Add(1)
							// go func(l string, f string) {
							// 	defer wg.Done()
							// 	HttpGet2(l, f)
							// }(link.String(), f)

							wg.Add(1)
							go HttpGet3(link.String(), f, &wg)

						}
					}
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	// fmt.Printf("now num process: %d\n", runtime.NumCgoCall())
	// fmt.Printf("now num goroutine: %d\n", runtime.NumGoroutine())

	wg.Wait()
	// 期待时间是最长一个下载页面的时间
	fmt.Printf("%.2fs run time\n", time.Since(start).Seconds())
}

// 递归遍历HTML的node
// 参数1： *html.Node
// 参数2和3：func(n *html.Node) 匿名函数，参数是*html.Node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	// 每次都是去第一个子子节点，当子节点等于nil的时候
	// 那么就取同胞的下一个节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// down文件
func HttpGet(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页body内容
	// 下载文件有问题
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		//读取结束，或者出问题
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])

	}
	return
}

// down文件
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

// down3文件
func HttpGet3(url string, fileName string, w *sync.WaitGroup) (result string, err error) {
	start := time.Now()

	tokens <- struct{}{}

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

	<-tokens

	w.Done()

	fmt.Printf("down page %s:  %.2fs sub run time\n", fileName, time.Since(start).Seconds())
	fmt.Printf("now num goroutine: %d\n", runtime.NumGoroutine())
	return
}

// 抓取文章资源内容
func downAssets(n *html.Node) {
	// 传入当前网页 html.node
	// 遍历node查找 link javascript src
	// 根据url创建相对路径文件夹
	// 组合图片的完整的URL
	// 保存图片到本地

	// css 标签是 link rel="stylesheet" href=""
	// js 标签 script src=""
	// image 标签 img src=""

	assetsNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			// 查找link标签
			if n.Data == "link" {
				//循环link属性

				for _, link := range n.Attr {
					if link.Key == "href" {
						//通过.css获取css文件
						if strings.LastIndex(link.Val, ".css") > 0 {
							// 创建相对目录 - 只要路径部分
							d := filepath.Dir(link.Val)
							os.MkdirAll(baseDir+d, fs.ModePerm)
							// 组合完整URL地址
							fullUrl := baseUrl + link.Val
							// fmt.Println(fullUrl)
							// https://books.studygolang.com/gopl-zh/
							// 保存文件到本地
							HttpGet2(fullUrl, baseDir+link.Val)
						}

					}
				}
			}

			if n.Data == "script" {
				//循环link属性
				for _, link := range n.Attr {
					if link.Key == "src" {
						fmt.Println(link.Val)
						d := filepath.Dir(link.Val)
						os.MkdirAll(baseDir+d, fs.ModePerm)
						// 组合完整URL地址
						fullUrl := baseUrl + link.Val
						HttpGet2(fullUrl, baseDir+link.Val)
					}
				}
			}

			if n.Data == "img" {
				//循环link属性
				for _, link := range n.Attr {
					if link.Key == "src" {
						if !strings.Contains(link.Val, "http") {
							fmt.Println(link.Val)
							d := filepath.Dir(link.Val)
							os.MkdirAll(baseDir+d, fs.ModePerm)
							// 组合完整URL地址
							fullUrl := baseUrl + link.Val
							HttpGet2(fullUrl, baseDir+link.Val)
						}
					}
				}
			}

		}
	}
	forEachNode(n, assetsNode, nil)
}
