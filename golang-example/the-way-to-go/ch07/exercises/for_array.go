package main

import "fmt"

// 写一个循环并用下标给数组赋值(从0到15) 并打印数组
func main() {
    var arr1 [15]int

    for i := 0; i < 15; i++ {
        arr1[i] = i
    }
    fmt.Println(arr1)
}
