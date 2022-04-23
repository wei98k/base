package main

import (
    "fmt"
)

func main() {
    // a := []byte("a") 
    // fmt.Println([]byte("b"))
    r := Compare([]byte("da"), []byte("b"))
    fmt.Println(r)
}

func Compare(a, b[]byte) int {
    for i:=0; i < len(a) && i < len(b); i++ {
        switch {
            case a[i] > b[i]:
                return 1
            case a[i] < b[i]:
                return -1
        }
    }

    // 数组的长度不同
    switch {
    case len(a) < len(b):
        return -1
    case len(a) > len(b):
        return 1
    }
    return 0
}
