package main

import (
    "fmt"
)

func main() {
    var arr = [5]int{0, 1, 2, 3, 4}
    //TODO:: 中括号+冒号可以把数组转成切片, 还有其他的方式吗？
    s := sum(arr[:])
    //s := sum(arr) // 报告错误 函数需要接收一个引用值
    //s := sum(&arr)
    fmt.Printf("sum result: %d\n", s)
}

func sum(a []int) int {
    s := 0
    for i := 0; i < len(a); i++ {
        s += a[i]
    }
    return s
}

