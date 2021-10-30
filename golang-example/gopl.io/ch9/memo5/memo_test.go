package memo5_test

import (
	"example/gopl.io/ch9/memo5"
	"example/gopl.io/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo5.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo5.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
