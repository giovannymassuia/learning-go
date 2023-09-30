package main

func main() {

	a := 10

	// Value of a
	println(a)

	// Memory address of a
	println(&a)

	// Pointer to a
	var pointer *int = &a
	println(pointer)
	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)

}
