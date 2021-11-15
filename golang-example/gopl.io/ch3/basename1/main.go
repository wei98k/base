package main

import (
	"bufio"
	"fmt"
	"os"
)

// basename函数灵感源于Unix shell的同名工具。在我们实现的版本中，
// basename(s)将看起来像是系统路径的前缀删除，同时将看似文件类型的后缀名部分删除：
func main() {
	// 创建一个命令行交换窗口
	input := bufio.NewScanner(os.Stdin)
	// 等待用户输出
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
