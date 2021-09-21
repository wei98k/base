package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestMyFib(t *testing.T) {
	go MySpinner(100 * time.Microsecond)
	const n = 45
	fibN := MyFib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
