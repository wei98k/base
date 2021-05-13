package main

import (
    "fmt"
    "strings"
)

func main() {
    //str := 'abc' + '123'
    //str := "abc" + '123'

    //str := "abc" + "123"
    //str := fmt.Sprintf("abc%d", 2233)
    a := []string{"a", "b", "c"}

    str := strings.Join(a, "a")
    fmt.Println(str)
}
