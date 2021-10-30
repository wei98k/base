package bank3_test

import (
	"example/gopl.io/ch9/bank3"
	"sync"
	"testing"
)

func TestBank3(t *testing.T) {
	var n sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank3.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank3.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
