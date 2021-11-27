package mandelbrot

import (
	"runtime"
	"testing"
)

func BeanchmarkSerialRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialRender()
	}
}

func benchConCurrentRender(b *testing.B, workers int) {
	for i := 0; i < b.N; i++ {
		ConcurrentRender(workers)
	}
}

func Benchmark1(b *testing.B) {
	benchConCurrentRender(b, 1)
}

func BenchmarkMaxProces(b *testing.B) {
	benchConCurrentRender(b, runtime.GOMAXPROCS(-1))
}

func Benchmark8(b *testing.B) {
	benchConCurrentRender(b, 8)
}

func Benchmark16(b *testing.B) {
	benchConCurrentRender(b, 16)
}
func Benchmark32(b *testing.B) {
	benchConCurrentRender(b, 32)
}
func Benchmark64(b *testing.B) {
	benchConCurrentRender(b, 64)
}
func Benchmark128(b *testing.B) {
	benchConCurrentRender(b, 128)
}
