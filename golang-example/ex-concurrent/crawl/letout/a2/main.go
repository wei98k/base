package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		r := queryAll()
		fmt.Println(r)
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() { ch <- query() }()
	}
	// 原因在于 channel 均已经发送了（每次发送 3 个），但是在接收端并没有接收完全（只返回 1 个 ch），所诱发的 Goroutine 泄露。
	// r := <-ch
	// fmt.Println(r)
	return <-ch
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
