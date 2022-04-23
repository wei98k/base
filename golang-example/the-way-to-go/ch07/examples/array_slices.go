package main

import (
    "fmt"
)

func main() {
    var arr1 [8]int
    // slice1 引用数组的地址 引用赋值
    // slice1 的值改变会影响到arr1
    // arr1的长度会影响到slice1的cap容量
    var slice1 []int = arr1[2:5]

    for i := 0; i < len(arr1); i++ {
        arr1[i] = i
    }

    // 打印切片
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("slice key: %d, val: %v\n", i, slice1[i])
    }
    
    //slice1[1] = 99

       
    fmt.Printf("arr1: %v\n", arr1)
    fmt.Printf("slice1: %v\n", slice1)
    fmt.Printf("arr1 len: %d\n", len(arr1))
    fmt.Printf("arr1 cap: %d\n", cap(arr1))
    fmt.Printf("slice1 len: %v\n", len(slice1))
    fmt.Printf("lice1 cap: %v\n", cap(slice1))

    // 切片赋值改切片
    slice1 = slice1[0:4]
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("slice at %d is %d\n", i, slice1[i])
    }

    fmt.Printf("slice len: %d\n", len(slice1));
    fmt.Printf("slice cap: %d\n", cap(slice1));

}
