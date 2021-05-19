package main

import (
    "fmt"
)

func main() {
    var s1 []int
    var s2 = []int{}
    
    fmt.Println(s1)
    fmt.Println(s2)

    if s2 == nil {
        fmt.Println("yes nil")
    } else {
        fmt.Println("no nil")
    }
}
