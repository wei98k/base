package main

import(
    "fmt"
)

func main() {
    // 声明
    var m map[string]int
    // 初始化
    m = map[string]int{"a": 1}
    m["a"] = 99
    m["b"] = 88
    fmt.Println(m)
    // # command-line-arguments
    // ./main.go:9:13: type map[string]int is not an expression
}
