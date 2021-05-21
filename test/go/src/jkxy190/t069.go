package main

import (
    "fmt"
)

func main() {
    
}

func f(n int) (r int) {
    defer func() {
        r += n
        recover()
    }()
    
    var f func()
    
    defer f()
    f = func() {
        r += 2
    }
    
    return n + 1
}

