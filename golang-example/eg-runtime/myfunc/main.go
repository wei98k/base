package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(1)
	// num := runtime.GOMAXPROCS(2)
	// fmt.Println(num)
	// fmt.Println(runtime.NumCPU())

	fmt.Printf("now num process: %d\n", runtime.NumCgoCall())
	go func() {
		fmt.Println("test ... test")
	}()
	fmt.Printf("now num goroutine: %d\n", runtime.NumGoroutine())
}
