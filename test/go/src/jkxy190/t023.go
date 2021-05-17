package main

import (
    "fmt"
)

func main() {    
    a := [2]int{5, 8}
    b := [2]int{5, 6}
    
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
