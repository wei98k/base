package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// Alice:
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	// Bob:
	go func() {
		defer wg.Done()
		Deposit(100)
	}()
	wg.Wait()
	fmt.Println(Balance())
}
