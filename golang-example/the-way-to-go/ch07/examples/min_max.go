package main

import (
    "fmt"
    "math"
)

func main() {
    sl1 := []int{78, 34, 634, 12, 90, 492, 13, 2}    
   max := maxSlice(sl1)
   fmt.Printf("max: %d\n", max)

   //fmt.Printf("maxint32: %d\n", math.MaxInt32)
   min := minSlice(sl1)
   fmt.Printf("min: %d\n", min)
}

func maxSlice(sl []int) (max int) {
    for _, v := range sl {
        if v > max {
            max = v
        } 
    }
    return
}

func minSlice(sl []int) (min int) {
    // int32 最大的值 
    min = math.MaxInt32
    for _, v := range sl {
        if v < min {
            min = v
        }
    }
    return
}
