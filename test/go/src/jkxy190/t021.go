package main

import "fmt"

// 可变参数和引用传值是一个样的?
func hello(num ...int) {
    num[0] = 18
    num[1] = 33
}

func main() {
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
    fmt.Println(i[1])
}

