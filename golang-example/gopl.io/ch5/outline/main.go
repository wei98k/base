package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 在函数outline中，我们通过递归的方式遍历整个HTML结点树，
// 并输出树的结构。在outline内部，每遇到一个HTML元素标签，
// 就将其入栈，并输出。

// outline有入栈操作，但没有相对应的出栈操作。
// 当outline调用自身时，被调用者接收的是stack的拷贝。
// 被调用者对stack的元素追加操作，修改的是stack的拷贝，
// 其可能会修改slice底层的数组甚至是申请一块新的内存空间进行扩容；
// 但这个过程并不会修改调用方的stack。
// 因此当函数返回时，调用方的stack与其调用自身之前完全一致。

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//  ../fetch https://golang.org | go run main.go

// out context:

// [html]
// [html head]
// [html head meta]
// [html head meta]
// [html head meta]
// [html head meta]
// [html head title]
// [html head link]
// [html head link]
// [html head link]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html head script]
// [html body]
// [html body header]
// [html body header div]
// [html body header div a]
// [html body header nav]
// [html body header nav a]
// [html body header nav a img]
// [html body header nav button]
// [html body header nav button div]
// [html body header nav ul]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body header nav ul li]
// [html body header nav ul li a]
// [html body main]
// [html body main div]
// [html body main div div]
// [html body main div div]
// [html body main div div section]
// [html body main div div section h1]
// [html body main div div section h1 strong]
// [html body main div div section h1 strong]
// [html body main div div section h1 strong]
// [html body main div div section i]
// [html body main div div section a]
// [html body main div div section a img]
// [html body main div div section p]
// [html body main div div section p br]
// [html body main div div section]
// [html body main div div section div]
// [html body main div div section div h2]
// [html body main div div section div a]
// [html body main div div section div]
// [html body main div div section div textarea]
// [html body main div div section div]
// [html body main div div section div pre]
// [html body main div div section div pre noscript]
// [html body main div div section div]
// [html body main div div section div select]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div select option]
// [html body main div div section div div]
// [html body main div div section div div button]
// [html body main div div section div div div]
// [html body main div div section div div div button]
// [html body main div div section div div div a]
// [html body main div div section]
// [html body main div div section h2]
// [html body main div div section div]
// [html body main div div section div a]
// [html body main div div section]
// [html body main div div section h2]
// [html body main div div section div]
// [html body main div div section div iframe]
// [html body main div script]
// [html body footer]
// [html body footer div]
// [html body footer div img]
// [html body footer div ul]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div a]
// [html body script]
