package main

import (
    "fmt"
)

func main() {
    s := make([]string, 10)
    a := []string{1, 2, 3}
    b := []int{4,5}
    copy(s, a)
    copy(s[5:], b)
    fmt.Printf("slice val: %v\n", s)
}
