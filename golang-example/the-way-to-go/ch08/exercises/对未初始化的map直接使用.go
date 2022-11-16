package main

import(
    "fmt"
)

func main() {
    // 声明
    var m map[string]int
    m["a"] = 99
    fmt.Println(m)
    // # command-line-arguments
    // ./main.go:9:13: type map[string]int is not an expression
}
