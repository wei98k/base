package main

import(
    "fmt"
)
type employee struct {
    salary float32
}

func (e *employee) giveRaise(scale float32) float32 {
    return e.salary * scale
}
func main() {
    e1 := &employee{3200.00}
    r := e1.giveRaise(0.2)
    fmt.Println(r)
}
