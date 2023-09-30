package main

import "fmt"

const a = "hello"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = " giovanny"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("E type is: %T", e)
	fmt.Printf("E type is: %v", e)
}
