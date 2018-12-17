package main

import (
	"fmt"
)

const c = "c"

var f int = 5

type T struct{}

func main() {
	var a int
	Func1()
	
	fmt.Println(a)
}

func (t T) Method1() {
	fmt.Println(t)
}

func Func1() {
	fmt.Println(11111)
}