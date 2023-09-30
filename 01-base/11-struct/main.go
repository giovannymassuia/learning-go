package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	customer1 := Customer{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", customer1.Name, customer1.Age, customer1.Active)

	customer1.Active = false
	fmt.Printf("Name: %s, Age: %d, Active: %t\n", customer1.Name, customer1.Age, customer1.Active)
}
