package main

import (
    "fmt"
)

func main() {
    s := "hello"
    a , b := SplitString(s, 2)
    fmt.Printf("a: %s, b: %s\n", a, b)
}

func SplitString(s string, pos int) (string, string) {
    str := []byte(s)
    a := str[:pos]
    b := str[pos:]
    return string(a), string(b)
}
