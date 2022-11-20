package main

import(
    "fmt"
)

func main() {
    // map的切片
    s1 := make([]map[int]int, 3)
    for i := range s1 {
        s1[i] = make(map[int]int, 1)
        s1[i][1] = 9
    }
    fmt.Println(s1)
    
    // 错误版本
    s2 := make([]map[int]int, 3)
    for _, v := range s2 {
        v = make(map[int]int, 1)
        v[1] = 8
    }
    fmt.Println(s2)
}
