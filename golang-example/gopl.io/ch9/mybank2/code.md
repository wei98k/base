
```golang
package main

import (
	"fmt"
	"sync"
)

var balance int

var deposits = make(chan int) //存款用channel
var balances = make(chan int) //接收余额用channel

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
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := Withdraw(200)
		if !res {
			fmt.Println("取款失败")
		}
	}()
	wg.Wait()
	b := Balance()
	fmt.Println(b)
}

/*
解决:
1.总余额限定在一个goroutine中,通过channel通讯
2.channel是会阻塞同一时间的多个goroutine的
*/
func teller() {
	var balance int //总余额限定在一个goroutine中
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

//取款用函数
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false // insufficient funds
	}
	return true
}
```