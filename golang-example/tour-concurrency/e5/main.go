package main

import "fmt"

/*

select语句让一个goroutine在多个通信操作中等待。

一个选择语句会阻塞，直到它的一个案例可以运行，然后它就执行这个案例。如果有多个准备好了，它就随机选择一个。
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// 
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 创建通道c
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}