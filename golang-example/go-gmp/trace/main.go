package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// go run main.go
// go tool trace trace.out
// 浏览器访问: http://127.0.0.1:57837/trace
func main() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main goroutine
	fmt.Println("are you ok")
}
