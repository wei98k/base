package main

import "fmt"

type ship struct {
    pulp *pulp
}

func NewShip(pulp *pulp) *ship {
    return &ship{
        pulp: pulp,
    }
}

type pulp struct {
    count int
}

func Newpulp() *pulp{
    return &pulp{
    }
}

func (c *pulp) set (count int) {
    c.count = count
}

func (c *pulp) get() int {
    return c.count
}

func main() {
    p := Newpulp()
    s := NewShip(p)
    s.pulp.set(12)

    fmt.Println(s.pulp.get())

    s.pulp.set(33)

    fmt.Println(s.pulp.get())
}
