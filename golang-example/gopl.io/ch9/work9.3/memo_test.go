package memo

import (
	"example/gopl.io/ch9/work9.3/memotest"
	"testing"
	"time"
)

const timeout = 5 * time.Minute

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(done)
	}()
	memotest.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	done := make(chan struct{})
	go func() {
		time.Sleep(timeout)
		close(done)
	}()
	memotest.Concurrent(t, m, done)
}
