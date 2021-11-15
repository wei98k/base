package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 练习 3.11：
// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

// 思路：
// 小数点处理 以小数点为分割 处理好后在把这两段字符串连接起来
// 判断第一个是否为正负号 如果是保留不变 从下标1开始
func main() {
	fmt.Println(comma("1234567.890123"))
}

func comma(s string) string {
	if s == "" {
		return ""
	}
	str := []byte(s)
	var buf bytes.Buffer
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		buf.WriteByte(s[0])
		str = str[1:]
	}
	// pointIndex := strings.LastIndex(s, ".")
	pointIndex := bytes.LastIndex(str, []byte("."))
	for i := 0; i < pointIndex; i++ {
		if (pointIndex-i)%3 == 0 && i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}

	buf.WriteString(string(str[pointIndex:]))
	return buf.String()
}

func comma2(s string) string {
	if s == "" {
		return ""
	}
	str := []byte(s)
	var buf bytes.Buffer
	sign := str[0]
	if sign == '+' || sign == '-' {
		buf.WriteByte(sign)
		str = str[1:]
	}

	last := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			last = str[i:]
			str = str[:i]
		}
	}

	for i := 0; i < len(str); i++ {
		if (len(str)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}
	buf.WriteString(string(last))
	return buf.String()
}
