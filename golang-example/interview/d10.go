package interview

import (
	"fmt"
	"time"
)

func MySpinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func MyFib(x int) int {
	if x < 2 {
		return x
	}
	return MyFib(x-1) + MyFib(x-2)
}
