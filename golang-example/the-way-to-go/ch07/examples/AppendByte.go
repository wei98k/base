package main

import (
    "fmt"
)

func main() {
    // var slice1 = []int{3,4,5,5}
    var b = []byte{1, 3, 4}
    slice2 := AppendByte(b, 2, 8, 9, 3, 88)
    fmt.Printf("new slice2: %v\n", slice2)
}

// 完成控制追加流程
func AppendByte(slice []byte, data ...byte) []byte {
    // 1. 获取原切片的长度
    m := len(slice)
    n := m + len(slice)
    // 2. 当切片长度的2倍小于c容量的值，开始创建新的切片
    if n > cap(slice) {
        // 3. 创建一个n+1两倍长度的新切片
        newSlice := make([]byte, (n+1)*2)
        // 4. 把旧切片的值复制到新切片中
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    // 5. 复制到指定位置(复制到原切片末尾-不然的会覆盖掉原来的切片的数据)
    copy(slice[m:n], data)
    return slice
}
