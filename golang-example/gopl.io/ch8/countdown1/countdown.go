package main

import (
	"fmt"
	"time"
)

// 基于Select多路复用
//

func main() {
	fmt.Println("commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
