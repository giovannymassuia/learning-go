package main

func main() {

	// if statement
	a := 1
	b := 2
	if a > b {
		println("a == b")
	} else if a <= b {
		println("a < b")
	} else {
		println("a > b")
	}

	if a == 1 && b == 2 {
		println("a == 1 && b == 2")
	} else if a == 1 || b == 2 {
		println("a == 1 || b == 2")
	} else {
		println("a != 1 && b != 2")
	}

	// switch statement
	switch a {
	case 1:
		println("a == 1")
	case 2:
		println("a == 2")
	default:
		println("a != 1 && a != 2")
	}

	// switch statement with struct
	type Person struct {
		name string
		age  int
	}
	p := Person{"John", 20}
	switch p {
	case Person{"John", 20}:
		println("John is 20")
	case Person{"John", 30}:
		println("John is 30")
	default:
		println("John is not 20 or 30")
	}
}
