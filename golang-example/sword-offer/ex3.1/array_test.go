package array

import (
	"testing"
)

// 对比两个函数哪个效率更高
// 空间比较、时间比较、执行次数比较、安全性比较
// 同样次数 用时间更少意味着效率高
// 同样次数 用空间更少意味着效率高
func BenchmarkDuplicate(b *testing.B) {
	var numbers = [7]int{2, 6, 1, 0, 4, 5, 3}
	for i := 0; i < b.N; i++ {
		duplicate(numbers)
	}
}
