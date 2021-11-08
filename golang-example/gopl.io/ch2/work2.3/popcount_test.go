package popcount

import "testing"

var x uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

// # go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: example/gopl.io/ch2/work2.3
// cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
// BenchmarkPopCount-4             1000000000               0.3544 ns/op
// BenchmarkPopCountLoop-4         367668963                3.177 ns/op
// PASS
