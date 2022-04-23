package main

import (
    "fmt"
)
// array_sum.go
// & 取地址 * 传指针 这样在函数调用的时候就不会发生变量拷贝
func main() {
    arr := [3]float64{7.0, 8.5, 9.1}

    x := Sum(&arr)
    fmt.Printf("The sum of the array is : %f\n", x)
}

func Sum(a *[3]float64) (sum float64) {
    a[1] = 1.1
    for _, v := range a {
        sum += v
    }
    return
}
