package main

import "fmt"

// 练习5.14： 使用breadthFirst遍历其他数据结构。
// 比如，topoSort例子中的课程依赖关系（有向图）、
// 个人计算机的文件层次结构（树）；
// 你所在城市的公交或地铁线路（无向图）

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}
	breathFirst(keys)
}

func breathFirst(worklist []string) {
	n := 1
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				fmt.Printf("%d: %s\n", n, item)
				n++

				worklist = append(worklist, prereqs[item]...)
			}
		}
	}
}
