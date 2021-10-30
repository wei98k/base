package blank_test

import (
	blank "example/gopl.io/ch9/blank1"
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	//Alice
	go func() {
		blank.Deposit(200)
		fmt.Println("=", blank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		blank.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := blank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
