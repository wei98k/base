package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// 随机生成文件
// 打包文件
// 解压文件

func main() {

	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	fmt.Print(io.ErrClosedPipe)
}
