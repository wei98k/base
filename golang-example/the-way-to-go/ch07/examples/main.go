package main

import (
    "fmt"
)

func main() {
    mp := func(v int) int {
        return v * 10
    }

   a := mapFunc(mp, []int{1,3,9}) 
   fmt.Println(a)
}

func mapFunc(m func(a int) int, list []int) []int {
    res := make([]int, len(list))
    for k, v := range list {
        res[k] = m(v)
    } 
    return res
}
