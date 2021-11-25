package main

import "fmt"

// 定义一个ByteCounter类型
type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	// Fprintf第一参数是Write方法 所以内部调用的时候就会执行 ByteCounter.Write
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

// 实现io.Writer方法
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
