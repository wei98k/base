package memo4_test

import (
	"example/gopl.io/ch9/memo4"
	"example/gopl.io/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Concurrent(t, m)
}
