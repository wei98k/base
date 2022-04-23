package main

import (
    "fmt"
)

func main () {
    items := [...]int{10, 20, 30, 40, 50} // 数组
    // 无法正常工作, 写for循环让值可以double?
    for _, item := range items[:] {
        item *= 2
    }
    // for 循环后的items 值是多少？
    fmt.Printf("%v\n", items)
}
