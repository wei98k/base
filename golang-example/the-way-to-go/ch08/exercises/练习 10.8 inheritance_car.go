package main

import(
    "fmt"
)
type Engine interface {
    Start()
    Stop()
}

type Car struct {
    Engine // Car 必须实现接口中的方法
    wheelCount int // Car 类中的属性
}

// 返回Car中的属性
func (car Car) numberOfWheels() int {
    return car.wheelCount
}

type Mercedes struct {
    Car
}

func (m *Mercedes) sayHiToMerkel() {
    fmt.Println("Hi Angela!")
}

// 必须实现接口方法
func (c *Car) Start() {
    fmt.Println("Car is started")
}

func (c *Car) Stop() {
    fmt.Println("car is stopped")
}

// Car中定义的新方法
func (c *Car) GoToWorkIn() {
    c.Start()

    c.Stop()
}

func main() {
    m := Mercedes{Car{nil, 4}}
    m.GoToWorkIn()
    m.sayHiToMerkel()
}
