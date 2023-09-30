package main

import "fmt"

func main() {

	salaries := map[string]int{"John": 1000, "Ana": 2000, "Paul": 3000}

	fmt.Println(salaries["John"])

	delete(salaries, "John")

	fmt.Println(salaries)

	salaries["Jay"] = 5000

	fmt.Println(salaries)

	// empty map
	map1 := make(map[string]int)
	map2 := map[string]int{}

	fmt.Println(map1, map2)

	// loop through map
	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d\n", name, salary)
	}

	// loop through map ignoring key
	for _, salary := range salaries {
		fmt.Printf("salary is %d\n", salary)
	}
}
