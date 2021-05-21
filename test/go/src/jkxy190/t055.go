package main

import(
    "fmt"
)

func main() {
    s1 := []int{1, 2, 3}
    fmt.Println(s1)
    s2 := s1[1:]
    s2[1] = 4
    fmt.Println(s1, s2)
}
