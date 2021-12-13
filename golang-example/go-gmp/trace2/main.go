package main

import (
	"fmt"
	"time"
)

//cmd
// go build main.go
// GODEBUG=schedtrace=1000 ./main
func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("are you ok")
	}
}
