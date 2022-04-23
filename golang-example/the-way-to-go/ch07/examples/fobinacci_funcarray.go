package main

import (
    "fmt"
)
// 主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片
var term = 15

func main() {
    result := fibarray(term)
    for ix, fib := range result {
        fmt.Printf("The %d-th Fibonacci number is: %d\n", ix, fib)
    }
}

func fibarray(term int) []int {
    farr := make([]int, term)
    farr[0], farr[1] = 1, 1

    for i := 2; i < term; i++ {
        farr[i] = farr[i-1] + farr[i-2]
    }
    return farr
}
