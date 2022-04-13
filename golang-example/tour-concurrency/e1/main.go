package main

import (
	"fmt"
	"time"
)

// hello word by concurrency

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world-")
	say("world")
}