package interview

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMyWaitTimeout(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	if MyWaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}
