package main

import "fmt"

func main() {
    //list := new([]int)

    //list = append(list, 1)
    //fmt.Println(list)
    // build ./t006.go:8:18: first argument to append must be slice; have *[]int
    
    list := make([]int, 0)
    
    list = append(list, 1)
    
    fmt.Println(list)
}
