package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 循环接收后关闭通道
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range读取通道
	for i := range c {
		fmt.Println(i)
	}
}