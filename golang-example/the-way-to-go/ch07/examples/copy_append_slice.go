package main

import (
    "fmt"
)

func main() {
    // 切片的复制与追加    
    // 必须创建一个新的更大的切片并把原分片的内容都拷贝过来
    // 如何创建一个小于原切片的大小呢？ 如果是目标切片是有值的呢？
    slFrom := []int{1, 2, 3}
    slTo := make([]int, 10)

    n := copy(slTo, slFrom)

    fmt.Println(slTo)
    fmt.Printf("Copied %d elements\n", n)

    sl3 := []int{1, 2, 3}
    sl3 = append(sl3, 4, 5, 6)
    fmt.Println(sl3)
}
