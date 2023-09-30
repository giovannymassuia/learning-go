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
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 10

	fmt.Println(array[1])
	fmt.Println(len(array))

	for i, v := range array {
		fmt.Printf("\nThe value of index [%d] is [%d]", i, v)
	}
}
