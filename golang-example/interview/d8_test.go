package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestProc(t *testing.T) {
	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					Proc()
				}()
			}
		}
	}()
	select {}
}
