package main

import (
	"fmt"
	"strings"
)

// 练习5.16：编写多参数版本的strings.Join。

func main() {
	fmt.Println(join("aaa", "kkk", "llll"))
}

func join(strs ...string) string {
	if len(strs) < 2 {
		return ""
	}
	sep := strs[len(strs)-1]
	last := strs[len(strs)-2]

	tempStrs := strs[:len(strs)-2]

	var b strings.Builder
	for _, s := range tempStrs {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(last)

	return b.String()
}
