package work91

import (
	"sync"
	"testing"
)

func TestWithdrawConcurrent(t *testing.T) {
	//存入1w
	Deposit(10000)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			Withdrawal(amount)
			wg.Done()
		}(i)
	}
	wg.Wait()
	if got, want := Balance(), 5050; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
