package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1. 创建一个文件text当前目录 输入几行内容
// 2. 当前目前执行命令 `cat text | go run main.go`

// tip: 如何是在命令行直接运行`go run main.go`, 程序是走到Scan()一直等待用户的输入 不会往下走的
// Sacn() 终止情况 1. ctrl+d 2. 文件 `cat input | go run main.go`
func main() {
	// 标准输入
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
