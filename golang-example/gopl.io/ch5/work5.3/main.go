package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 练习 5.3： 编写函数输出所有text结点的内容。
// 注意不要访问<script>和<style>元素，
// 因为这些元素对浏览者是不可见的。

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, a := range n.Attr {
			texts = append(texts, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}
	return texts
}
