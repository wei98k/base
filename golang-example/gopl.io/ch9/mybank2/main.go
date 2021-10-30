package main

import (
	"fmt"
	"sync"
)

var balance int

var deposits = make(chan int) //存款用chancel
var balances = make(chan int) //接收余额用chancel

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func main() {
	go teller()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(300)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := Withdraw(401)
		if !res {
			fmt.Println("取款失败")
		}
	}()
	wg.Wait()
	b := Balance()
	fmt.Println(b)
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}
