package main

import (
    "fmt"
)

// 切片重组
func main() {
    slice1 := make([]int, 0, 10)
    // 不能超过cap

    for i := 0; i < cap(slice1); i++ {
        slice1 = slice1[0:i+1]
        slice1[i] = i
        fmt.Printf("The length of slice is %d\n", len(slice1))
    }

    for i := 0; i < len(slice1); i++ {
        fmt.Printf("slice at %d is %d\n", i, slice1[i])
    }
}

