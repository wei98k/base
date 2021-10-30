```golang
package main

import (
	"fmt"
	"sync"
)

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

/*
问题:
1.在Alice运行期间 balance = balance + amount 这一步运算可能会被Bob中间挤占
2.当运行到balance + amount的时候,Bob的正好赶到,然后继续运行blance=
3.此时Bob的增加的数据会丢失
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// Alice:
	go func() {
		defer wg.Done()
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()
	wg.Add(1)
	// Bob:
	go func() {
		defer wg.Done()
		Deposit(100)
	}()
	wg.Wait()
	res := Balance()
	fmt.Println(res)
}

```