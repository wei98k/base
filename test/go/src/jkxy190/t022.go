package main

import (
    "fmt"
)

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    t := a[2:3:4]
    //  第三个参数为容量
    //  第三个参数不能比第二个参数小
    fmt.Println(t[0])
}
