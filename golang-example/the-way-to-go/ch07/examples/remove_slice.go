package main

import (
    "fmt"
)

func main() {
    // 7.12 写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
    // 思路: start end 截取 0 到start 部分 然后在截取 end 到末尾部分 重新组合新的切片
    a := []int{1, 2, 3, 4, 5, 6, 7}
    r := removeSlice(a, 2, 4)
    fmt.Printf("new slice: %v\n", r)
}

func removeSlice(s []int, start, end int) (r []int) {
    len := len(s) - (end - start)
    r = make([]int, len)
    n := copy(r, s[:start])
    copy(r[n:], s[end:])
    return
}
