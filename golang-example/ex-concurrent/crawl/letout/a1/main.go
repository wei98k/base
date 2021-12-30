package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {}()
	go func() {}()
	fmt.Printf("now num goroutine: %d\n", runtime.NumGoroutine())
}
