package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    ch := make(chan string)
    
    go func() {
        p := <-ch
        fmt.Println(p)
    }()
    
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    
    ch <- "paper"
}
