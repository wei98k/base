package main

import (
	"fmt"
	"reflect"
)

type a1 struct {
	name string
}

type a2 struct {
	title string
}

func main() {
	b1 := a1{name: "hello"}
	b2 := a2{title: "hello"}
	res := reflect.DeepEqual(b1, b2)
	fmt.Printf("result: %v\n", res)
}
