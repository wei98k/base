package main

import (
    "fmt"
)

func main() {
    // 定义Buffer的方式?
    var buffer bytes.Buffer // 通过var定义
    var r *bytes.Buffer = new(bytes.Buffer) // 通过new获得一个指针
    // func NewBuffer(buf []byte) *Buffer // 通过函数参数创建并且用buf初始化好
}
