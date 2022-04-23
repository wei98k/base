package main

import (
    "fmt"
)

func main() {
    // 7.11 写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
    var s1 = []string{"a", "b", "c"}
    var s2 = []string{"d", "e", "f", "g"}
    r := insertSlice(s1, s2, 0)
    fmt.Printf("new slice: %v\n", r)
}

func insertSlice(s1, s2 []string, pos int) (s []string) {
    s = make([]string, len(s1)+len(s2))
    n := copy(s, s1[:pos])
    n += copy(s[pos:], s2)
    fmt.Printf("copy return len: %d\n", n)
    copy(s[n:], s1[pos:])
    return
}
