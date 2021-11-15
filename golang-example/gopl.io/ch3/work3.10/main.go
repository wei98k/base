package main

import (
	"bytes"
	"fmt"
)

// 练习 3.10： 编写一个非递归版本的comma函数，
// 使用bytes.Buffer代替字符串链接操作。

// 思路：循环字符串，以字符串数组下标%3 可以整除增加符号,

func main() {
	fmt.Println(comma("1231234343"))
}

func comma(s string) string {
	if s == "" {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
