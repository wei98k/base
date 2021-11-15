package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	//fmt.Println("c1: ", c1)
	//fmt.Println("c2: ", c2)
	fmt.Println(sha256DiffBitCount(c1, c2))
}

func sha256DiffBitCount(c1, c2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += diffBitCount(c1[i], c2[i])
	}
	return count
}

func diffBitCount(b1, b2 byte) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		mask := byte(1 << i)
		// 基础的8个bit
		// aa := b1 & mask
		// fmt.Printf("与后结果: %b\n", aa)
		if b1&mask != b2&mask {
			count++
		}
	}
	return count
}
