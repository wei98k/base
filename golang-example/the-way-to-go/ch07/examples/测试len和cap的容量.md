package main

import (
    "fmt"
)

func main() {
    // 生成相同的切片
   //make([]int, 50) 
   //new([100]int)[0:50]
   // 切片内的元素不能超过len, 那cap的意义是什么？
   var slice1 []int = make([]int, 2, 4)
   // slice1[-1] = 2 // 不可以是负数
   slice1[0] = 9
   slice1[1] = 8
   // slice1[2] = 7 // 手动赋值超过len就会报告错误

   // 通过append函数会产生新的切片 当超过原来的cap后就x2
   slice1 = append(slice1, 7)
   slice1 = append(slice1, 6)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)
   slice1 = append(slice1, 5)

   fmt.Printf("len: %d, cap: %d, val: %v\n", len(slice1), cap(slice1), slice1)
}
