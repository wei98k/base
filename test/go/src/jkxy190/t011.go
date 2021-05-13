package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
    var i int = 0
    var i1 MyInt1 = MyInt1(i)
    
    fmt.Println(i1)
}
