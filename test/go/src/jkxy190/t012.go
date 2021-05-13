package main

import "fmt"

func main() {
    a := []int{7, 8, 9}    

    b := append(a, 10)

    fmt.Println("xxxx: ", b)

    fmt.Printf("%+v\n", a)
    ap(a)
    fmt.Printf("%+v\n", a)
    app(a)
    fmt.Printf("%+v\n", a)
    
     
    
}


func ap(a []int) {
    a = append(a, 10)
    
    fmt.Printf("%+v\n", a)
}

func app(a []int) {
    a[0] = 1
}
