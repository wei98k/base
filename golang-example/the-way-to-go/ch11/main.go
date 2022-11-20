package main

import (
    "fmt"
)

// 定义接口和方法
type Simpler interface {
    Get() int
    Put(int)
}
// 定义结构实现接口方法
type Simple struct {
    i int
}

func (p *Simple) Get() int {
    return p.i
}

func (p *Simple) Put(u int) {
    p.i = u
}

//测试接口方法

func fI(it Simpler) int {
    it.Put(5)
    return it.Get()
}

func main() {
    var s Simple
    fmt.Println(fI(&s))
}

