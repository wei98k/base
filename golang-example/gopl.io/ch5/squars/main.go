package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	// 在main程序没有结束前x变量存在内存中
	var x int
	return func() int {
		// 每次调用都会影响到上一个函数的x
		x++
		return x * x
	}
}
