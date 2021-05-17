// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitShip() *Ship {
	pulp := NewPulp()
	ship := NewShip(pulp)
	return ship
}

// wire.go:

type Ship struct {
	pulp *Pulp
}

func NewShip(pulp *Pulp) *Ship {
	return &Ship{
		pulp: pulp,
	}
}

type Pulp struct {
	count int
}

func NewPulp() *Pulp {
	return &Pulp{}
}

func (c *Pulp) set(count int) {
	c.count = count
}

func (c *Pulp) get() int {
	return c.count
}

func main() {
}