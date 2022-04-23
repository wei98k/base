package main

import (
    "fmt"
)

func main() {
    // 切片的复制与追加    
    // 必须创建一个新的更大的切片并把原分片的内容都拷贝过来
    // 如果创建一个小于原切片的大小呢？ 答：不会报错误，只是复制之后容量等于切片的长度
    // 如果是目标切片是有值的呢? 答：会覆盖掉前面的值 
    
    slFrom := []int{1, 2, 3}
    //slTo := make([]int, 2)
    slTo2 := []int{4, 5, 6, 7, 8}

    n := copy(slTo2, slFrom)

    fmt.Println(slTo2)
    fmt.Printf("Copied %d elements\n", n)

    sl3 := []int{1, 2, 3}
    sl3 = append(sl3, 4, 5, 6)
    fmt.Println(sl3)
}
