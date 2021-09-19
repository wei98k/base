package interview

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestVisit(t *testing.T) {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)
	wait := &sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()

				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.Visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wait.Wait()
	fmt.Println("success: ", success)
}
