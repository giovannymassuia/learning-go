package main

import "fmt"

type Account struct {
	balance float64
}

func (a Account) simulate(amount float64) float64 {
	a.balance += amount
	return a.balance
}

func (a *Account) apply(amount float64) float64 {
	a.balance += amount
	return a.balance
}

func NewAccount() *Account {
	return &Account{0}
}

func main() {
	c1 := Account{100}

	simulationResult := c1.simulate(10)

	fmt.Printf("Simulation result: %f\n", simulationResult)
	fmt.Printf("Balance: %f\n", c1.balance)

	applicationResult := c1.apply(10)
	fmt.Printf("Application result: %f\n", applicationResult)
	fmt.Printf("Balance: %f\n", c1.balance)

	newAccount := NewAccount()
	fmt.Printf("New account balance: %f\n", newAccount.balance)
	newAccount.balance = 100
	fmt.Printf("New account balance: %f\n", newAccount.balance)
}
