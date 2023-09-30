package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

func (c Customer) Deactivate() {
	c.Active = false
	fmt.Println(c)
}

func main() {
	customer1 := Customer{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", customer1.Name, customer1.Age, customer1.Active)

	customer1.Deactivate()

	fmt.Println(customer1)
}
