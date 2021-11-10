package mandelbrot

import (
	"image/color"
	"testing"
)

func bench(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMandelbrotComplex64(b *testing.B) {
	bench(b, mandelbrot64)
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	bench(b, mandelbrot128)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	bench(b, mandelbrotBigFloat)
}

func BenchmarkMandelbrotBigRat(b *testing.B) {
	bench(b, mandelbrotBigRat)
}

// # go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: example/gopl.io/ch3/work3.8
// cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
// BenchmarkMandelbrotComplex64-4          19240090             60.27 ns/op
// BenchmarkMandelbrotComplex128-4         20285808             57.55 ns/op
// BenchmarkMandelbrotBigFloat-4            1803558            738.3 ns/op
// BenchmarkMandelbrotBigRat-4               521125           2284 ns/op
// PASS
