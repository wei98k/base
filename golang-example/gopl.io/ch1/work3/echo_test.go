package echo

import (
	"testing"
	"time"
)

const n = 5000000

// 1. 如何在命令行下运行下面测试用例? ok
// 2. 为什么要用testing.B BenchmarkEcho命名是否确定的?

func BenchmarkEcho1(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo1()
	}
	b.Log(time.Since(now))
}

func BenchmarkEcho2(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo2()
	}
	b.Log(time.Since(now))
}

func BenchmarkEcho3(b *testing.B) {
	b.N = n
	now := time.Now()
	for i := 0; i < b.N; i++ {
		echo3()
	}
	b.Log(time.Since(now))
}
