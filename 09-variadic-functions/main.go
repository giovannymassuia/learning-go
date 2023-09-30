package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
