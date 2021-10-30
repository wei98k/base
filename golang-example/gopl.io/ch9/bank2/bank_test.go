package bank2_test

import (
	"example/gopl.io/ch9/bank2"
	"sync"
	"testing"
)

func TestBank2(t *testing.T) {
	//模拟存1 存1000次
	var n sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank2.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank2.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
