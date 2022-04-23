package main

import (
    "fmt"
)

func main() {
    var slice1 []int = make([]int, 10)

    for i := 0; i < len(slice1); i++ {
        slice1[i] = 5 * i
    }

    // 打印这个切片
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("key: %d, val: %d\n", i, slice1[i])
    }

    fmt.Printf("len: %d, cap: %d\n", len(slice1), cap(slice1))

}
