package main

import (
    "fmt"
)

func main() {
   var arrF = [4]float32{2.1, 3.1, 1.4, 1.1} 
   sum := Sum(arrF)
   fmt.Printf("数组运算: %v\n", sum)

   sum1 := Sum1(arrF[:])
   fmt.Printf("数组运算: %v\n", sum1)

   var arrI = []int{1, 2, 3, 4}
   a, b := SumAndAverage(arrI)
   fmt.Printf("sum-int: %d, average-float32: %f\n", a, b)
}

// 对数组运算
func Sum(arrF [4]float32) (sum float32){
    for _, v := range arrF {
        sum += v
    }
    return 
}

// 对切片运算
func Sum1(arrF []float32) (sum float32) {
    for _, v := range arrF {
        sum += v
    }
    return
}

func SumAndAverage(arrI []int) (int, float32) {
    sum := 0
    for _, v := range arrI {
        sum += v
    }
    return sum, float32(sum / len(arrI))
}

