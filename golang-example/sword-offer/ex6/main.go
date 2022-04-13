package main

import "fmt"

// 定义链表结构
type NodeList struct {
	Val  int
	Next *NodeList
}

// 递归打印链表
func printListFromTailToHead(head *NodeList) {
	if head != nil {
		printListFromTailToHead(head.Next)
		fmt.Printf("%d -> ", head.Val)
	}
}

func main() {
	n3 := &NodeList{3, nil}
	n2 := &NodeList{2, n3}
	n1 := &NodeList{1, n2}

	fmt.Printf("\n NodeList 1 -> 2 -> 3 \n")

	// 测试
	fmt.Printf("\n Output: ")
	printListFromTailToHead(n1)
	fmt.Printf("\n \n")
}
