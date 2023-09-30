package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))

	value, err := sub(1, 2)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("all good: %d\n", value)
	}

}

func sum(a int, b int) (int, bool) {
	if a+b >= 50 {
		return a - b, false
	}
	return a + b, true
}

func sub(a, b int) (int, error) {
	if a < b {
		return 0, errors.New("a must not be less than b")
	}
	return a - b, nil
}
