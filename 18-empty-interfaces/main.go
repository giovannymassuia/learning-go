package main

import "fmt"

// empty interfaces

func main() {

	var x interface{} = 10
	var y interface{} = "Hello"

	fmt.Println(x)
	fmt.Println(y)

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("%T and %v\n", t, t)

	switch t.(type) {
	case int:
		fmt.Println("-> int")
	case string:
		fmt.Println("-> string")
	default:
		fmt.Println("-> unknown")
	}
}
