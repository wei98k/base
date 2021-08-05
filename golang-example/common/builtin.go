package common

import "fmt"

// builtin Golang 内置函数

const (
	true = 0 == 0
)

func constantDome() {
	fmt.Println(true)
}

// 常用内置函数

func appendDome()  {
	slice := append([]byte("hello"), "world"...)
	fmt.Println(slice)
}
