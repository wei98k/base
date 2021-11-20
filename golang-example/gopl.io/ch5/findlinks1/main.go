package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// visit函数遍历HTML的节点树，
// 从每一个anchor元素的href属性获得link,
// 将这些links存入字符串数组中，并返回这个字符串数组。

// run: ../fetch https://golang.org | ./main
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
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
