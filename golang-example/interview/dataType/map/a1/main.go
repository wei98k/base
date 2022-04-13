package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2

	fmt.Println(m1)

	m1["a"] = 3

	fmt.Println(m1)

}
