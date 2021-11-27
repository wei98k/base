package main

import (
	"fmt"
	"os"
	"time"
)

// 现在我们让这个程序支持在倒计时中，
// 用户按下return键时直接中断发射流程。

func main() {
	// create abort channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// 打印结果
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
