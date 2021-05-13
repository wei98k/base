package main

import "fmt"

// 1. 必须使用显示初始化
// 2. 不能提供数据类型, 编译器会自动推导
// 3. 只能在函数内部使用简短模式
var (
    size int = 100
    //max_size = size * 2
)

func main() {
   fmt.Println(size) 
   //fmt.Println(max_size) 
}
