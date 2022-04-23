package main

import (
    "fmt"
)
// pointer_array2.go
// 取任意数组常量的地址来作为指向新实例的指针

// 接收一个数组指针, 打印出指针的值
func fp(a *[3]int) { fmt.Println(a) }

func main() {
    for i := 0; i < 3; i++ {
        arr := [3]int{i, i*i, i*i*i}
        fp(&arr)
    }
}
