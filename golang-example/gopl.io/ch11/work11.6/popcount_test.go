package popcount

import (
	"example/gopl.io/ch11/work11.6/bitcount"
	"example/gopl.io/ch11/work11.6/popcount"
	"testing"
)

const bin = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(bin)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.BitCount(bin)
	}
}

func BenchmarkClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.Clearing(bin)
	}
}

func BenchmarkShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitcount.Shifting(bin)
	}
}
