package main

import "fmt"

func main() {
    a := []byte{1, 0}
    a = append(a, 1, 1, 1)
    fmt.Println("cap of a is ",cap(a))
    
    b := []int{23, 51}
    b = append(b, 4, 5, 6)
    fmt.Println("cap of b is ",cap(b))
    
    c := []int32{1, 23}
    c = append(c, 2, 5, 6)
    fmt.Println("cap of c is ",cap(c))

    type D struct{
        age byte
        name string

    }
    d := []D{
        {1,"123"},
        {2,"234"},
    }

    d = append(d,D{4,"456"},D{5,"567"},D{6,"678"})
    fmt.Println("cap of d is ",cap(d))
}