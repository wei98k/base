package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 程序是findlinks的改进版本。
// 修改后的findlinks可以自己发起HTTP请求，
// 这样我们就不必再运行fetch。
// 因为HTTP请求和解析操作可能会失败，
// 因此findlinks声明了2个返回值：
// 链接列表和错误信息。一般而言，
// HTML的解析器可以处理HTML页面的错误结点，
// 构造出HTML页面结构，所以解析HTML很少失败。
// 这意味着如果findlinks函数失败了，
// 很可能是由于I/O的错误导致的。

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}
