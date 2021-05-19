package main

import (
    "fmt"
)

func main() {
    // 这个是切片还是数组?
    s := [3]int{1, 2, 3}
    //a := s[:2]
    c := s[1:2:cap(s)]
    c = append(c, 1)
    c = append(c, 2)
    c = append(c, 3)
    // 如何读取切片的长度和容量?
    fmt.Println(c, len(c), cap(s))
}

