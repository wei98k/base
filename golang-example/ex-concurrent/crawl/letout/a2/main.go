package main

// https://medium.com/golangspec/goroutine-leak-400063aef468

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "net/http/pprof"
)

// func main() {
// 	for i := 0; i < 4; i++ {
// 		r := queryAll()
// 		fmt.Println(r)
// 		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
// 	}
// }

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(rep http.ResponseWriter, req *http.Request) {
	var a int
	for i := 0; i < 4; i++ {
		a += queryAll()
		// fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
	rep.Write([]byte(fmt.Sprintf("%d", a)))
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
