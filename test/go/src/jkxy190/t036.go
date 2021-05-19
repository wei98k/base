package main

import (
    "fmt"
)

func main() {
   //sum := add(1, 3, 8) 
   sum := add([]int{1,3,8}...) 
   fmt.Println(sum)
}

func add(args ...int) int {
    sum := 0 
    for _, arg := range args {
        sum += arg
    }
    return sum
}
