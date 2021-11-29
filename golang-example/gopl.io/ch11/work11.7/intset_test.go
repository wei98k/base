package intset

import (
	"math/rand"
	"testing"
	"time"
)

var (
	s1 []int
	s2 []int
)

const (
	n     = 100000
	scale = 100
)

func initSet() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < n; i++ {
		s1 = append(s1, rng.Intn(scale*n))
		s2 = append(s2, rng.Intn(scale*n))
	}
}

func BenchmarkIntSetAdd(b *testing.B) {
	initSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := &IntSet{}
		for _, k := range s1 {
			s.Add(k)
		}
	}
}

func BenchmarkIntSetHas(b *testing.B) {
	initSet()
	s := &IntSet{}
	for _, k := range s1 {
		s.Add(k)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range s1 {
			s.Has(s1[i])
		}
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {
	initSet()
	is1 := &IntSet{}
	for _, k := range s1 {
		is1.Add(k)
	}
	is2 := &IntSet{}
	for _, k := range s2 {
		is2.Add(k)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is1.UnionWith(is2)
	}
}

func BenchmarkMapAdd(b *testing.B) {
	initSet()
	s := make(map[int]bool)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			s[k] = true
		}
	}
}

func BenchmarkMapHas(b *testing.B) {
	initSet()
	s := make(map[int]bool)
	for _, k := range s1 {
		s[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range s1 {
			_ = s[k]
		}
	}
}

func benchmarkMapUnionWith(b *testing.B) {
	initSet()
	ms1 := make(map[int]bool)
	for _, k := range s1 {
		ms1[k] = true
	}
	ms2 := make(map[int]bool)
	for _, k := range s2 {
		ms2[k] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := range ms2 {
			ms1[k] = true
		}
	}
}
