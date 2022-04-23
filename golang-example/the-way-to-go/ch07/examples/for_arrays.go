package main
import "fmt"

func main() {
    // 定义arr1 数组 长度为5
    var arr1 [5]int
    // 循环数组 并压入值
    for i:=0; i < len(arr1); i++ {
        arr1[i] = i * 2
    }

    // 遍历打印数组
    for i:=0; i < len(arr1); i++ {
        fmt.Printf("Array at index %d is %d\n", i, arr1[i])
    }
}
