package main

import (
    "fmt"
)

func main() {
    const cap = 5
    ch := make(chan string, cap)

    go func() {
        // range 用于接收消息
        for p := range ch {
            fmt.Println("接收: ", p)
        }
    }()

    const work = 20

    for w := 0; w < work; w++ {
        select {
            case ch <- "paper":
                    fmt.Println("manager: send ack")
            default:
                    fmt.Println("manager: drop")
        }
    }

    close(ch)
}
