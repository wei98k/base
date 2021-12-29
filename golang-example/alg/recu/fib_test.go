package recu

import (
	"testing"
)

func TestFib(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := fib(i)
		t.Log(n)
	}
}