package main

import "fmt"

func main() {
   // 证明当数组赋值时, 发生了数组内存拷贝 
   // 当arr1 赋值给 arr2 而当arr2改变后arr1不变, 就可以证明数组发生了内存拷贝
   
   var arr1 [5]int 
   for i := 0; i < len(arr1); i++ {
        arr1[i] = i * 2
   }

   arr2 := arr1
   arr2[2] = 99
   
   for i := 0; i < len(arr1); i++ {
        fmt.Printf("Array arr1 at index %d is %d\n", i, arr1[i])
   }
   fmt.Println()
   for i :=0; i < len(arr2); i++ {
        fmt.Printf("Array arr2 at index %d is %d\n", i, arr2[i])
   }
}
