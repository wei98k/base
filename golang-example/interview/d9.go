package interview

import (
	"sync"
	"time"
)

func MyWaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() {
		ch <- true
	})
	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
