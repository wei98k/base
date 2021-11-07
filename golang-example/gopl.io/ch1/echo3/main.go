// 打印命令行参数
// 连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用strings包
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
