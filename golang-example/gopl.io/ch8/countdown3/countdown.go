package main

import (
	"fmt"
	"os"
	"time"
)

// 下面让我们的发射程序打印倒计时。
// 这里的select语句会使每次循环迭代等待一秒来执行退出操作。

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commentcing countdown. press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//Do nothing
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
