package main

import (
    "fmt"
)

func main() {
    var x int
    
    x, _ := f() 
    fmt.Println(x)
}

func f() (int, string) {
   return 1, "str" 
}
