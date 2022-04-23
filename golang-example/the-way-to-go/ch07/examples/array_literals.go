package main

import (
    "fmt"
)
// array_iterals.go
// 关于数组初始化
func main() {
    // 固定长度
    // var arr = [5]int{18, 13, 14, 22, 19} 
    // 用省略号不具体指定长度
    // var arr = [...]int{18, 13, 14, 22, 19} 
    // 长度为空 切片
    // var arr = []int{3, 2, 4, 3}
    // 指定 值在数组的key位置
    // var arr = [5]string{3: "Chris", 4: "Ron"}
    var arr = []string{3: "Chris", 4: "Ron"}

    for i := 0; i < len(arr); i++ {
        fmt.Printf("key: %d val: %v\n", i, arr[i])
    }
}
