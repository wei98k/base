package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// 练习5.17：编写多参数版本的ElementsByTagName，
// 函数接收一个HTML结点树以及任意数量的标签名，
// 返回与这些标签名匹配的所有元素。下面给出了2个例子：

func main() {
	resp, err := http.Get("https://www.sulinehk.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	nodes := ElementByTagName(doc, "h1")
	for _, n := range nodes {
		fmt.Println(n.Data)
		fmt.Println(n.Attr)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, data := range name {
			if doc.Data == data {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementByTagName(c, name...)...)
	}
	return nodes
}
