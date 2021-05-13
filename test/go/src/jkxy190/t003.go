package main


import "fmt"


func main() {
    //s := make([]int, 5)
    //s = append(s, 1, 2, 3)
    //fmt.Println(s)
    // 输出结果: [0 0 0 0 0 1 2 3]
    
    s := make([]int, 0)
    s = append(s, 1, 2, 3, 4)
    fmt.Println(s)
    // 输出结果: [1 2 3 4]
}
