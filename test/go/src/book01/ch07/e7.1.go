package main

import (
    "fmt"
    "bufio"
    "bytes"
)

func main() {
    s := "Hello, World!\nHello, 世界！"
    var wc WordCounter
    fmt.Fprintf(&wc, s)
    fmt.Println(wc)

    var lc LineCounter
    fmt.Fprintf(&lc, s)
    fmt.Println(lc)
}


type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {

    scanner := bufio.NewScanner(bytes.NewReader(p))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        *c++
    }
    fmt.Println("call func write")
    //return len(p), nil
    return 222, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(bytes.NewReader(p))   
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        *c++
    }
    return 0, nil
}

