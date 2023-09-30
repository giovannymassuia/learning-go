package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	// `:` is a slice of the slice
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// ignores after the 4th element
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// ignores first 2  elements
	fmt.Printf("len=%d cap=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// every time the `append` is called, and it has exceeded its current capacity, it creates a new array with double the size
	slice = append(slice, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
