package main

import (
   "fmt" 
)

func main() {
    p := *f()   
    fmt.Println(p.m)
}

type S struct {
    m string
}

func f() *S {
    return &S{"foo"}
}


