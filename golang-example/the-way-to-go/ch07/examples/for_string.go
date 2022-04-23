package main

import (
    "fmt"
)

func main() {
    // 使用for range遍历一个字符串(字符串本质是一个数组)
    s := "\u00ff\u754c"
    for k, v := range s {
        fmt.Printf("%d, %c\n", k, v)
    }
}
