// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for k, v := range os.Args[1:] {
		// s += "key: " + string(k) + " value: " + v + " \n"
		// %v 通用默认  %d 数字 %s 字符串
		s += fmt.Sprintf("key: %d, value: %v", k, v)
	}
	fmt.Println(s)
}
