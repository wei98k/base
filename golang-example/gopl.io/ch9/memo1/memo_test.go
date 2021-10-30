package memo1_test

import (
	"example/gopl.io/ch9/memo1"
	"example/gopl.io/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo1.New(httpGetBody)

	memotest.Sequential(t, m)
}

//测试线程安全问题
func TestConcurrent(t *testing.T) {
	m := memo1.New(httpGetBody)
	memotest.Concurrent(t, m)
}
