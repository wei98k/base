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

func Newpulp(count int) *pulp{
    return &pulp{
        count: count,
    }
}

func main() {
    p := Newpulp(12)
    s := NewShip(p)

    fmt.Println(s.pulp.count)
}
