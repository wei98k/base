package memo3_test

import (
	"example/gopl.io/ch9/memo3"
	"example/gopl.io/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Concurrent(t, m)
}
