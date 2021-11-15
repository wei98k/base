package main

import "fmt"

// 单个整型最加slice
func appendInt(x []int, y int) []int {
	var z []int
	// 新的slice长度
	zlen := len(x) + 1
	// 判断新的长度小于等于久的长度
	if zlen <= cap(x) {
		z = x[:zlen]
	} else { // 当大于久的长度时候
		// 把新的长度赋值给新变量
		zcap := zlen
		// 如果追加新的长度小于旧的两倍的话
		if zcap < 2*len(x) {
			// 扩展长度2倍
			zcap = 2 * len(x)
		}
		// 如果是大于两倍的长度 直接扩展成对应的容量
		z = make([]int, zlen, zcap)
		// 把x复制到z
		copy(z, x) // 目标, 源头
	}
	z[len(x)] = y
	return z
}

func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
