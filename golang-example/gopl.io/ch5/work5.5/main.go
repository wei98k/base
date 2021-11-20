package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// 练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）

var (
	images int
	words  int
)

func main() {
	err := CountWordsAndImages("https://www.sulinehk.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("images: %d\nwords: %d\n", images, words)
}

func CountWordsAndImages(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	// 函数结束后 关闭http连接
	defer resp.Body.Close()
	// 解析响应的HTML文本
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing HTML: %s", err)
	}
	// 调用统计单词和IMG标签数量函数
	countWordsAndImages(doc)
	return nil
}

func countWordsAndImages(n *html.Node) {
	// 从解析好的html中，判断节点是否为text标签
	if n.Type == html.TextNode {
		// 将字符串以空格标识分割单词
		words += len(strings.Fields(n.Data))
	}
	// 判断节点标签为img
	if n.Type == html.ElementNode && n.Data == "img" {
		// 累加img的数量
		images++
	}
	// 如果下一层还有的话, 递归循环
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordsAndImages(c)
	}
}
