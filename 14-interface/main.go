package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

// - only methods can be declared for interfaces
// - interfaces can't have fields

type Person interface {
	Deactivate()
}

func (c Customer) Deactivate() {
	c.Active = false
	fmt.Println(c)
}

func Deactivation(p Person) {
	p.Deactivate()
}

func main() {
	customer1 := Customer{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	fmt.Println(customer1)

	Deactivation(customer1)
}
