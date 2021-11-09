package popcount

import "testing"

var x uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClean(x)
	}
}

// # go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: example/gopl.io/ch2/work2.5
// cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
// BenchmarkPopCount-4             1000000000               0.3845 ns/op
// BenchmarkPopCountClean-4        57180925                17.70 ns/op
// PASS
