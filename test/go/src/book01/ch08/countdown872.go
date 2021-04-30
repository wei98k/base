package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i :=0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x, i)
		case ch <-i:
			fmt.Println(i)
		}
	}
}
