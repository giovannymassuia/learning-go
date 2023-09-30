package main

func sum(a, b int) int {
	a = 50
	return a + b
}

func sumWithPointer(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	var1 := 1
	var2 := 2

	result := sum(var1, var2)

	println(var1, var2)
	println(result)

	println("--- with pointers")
	result = sumWithPointer(&var1, &var2)

	println(var1, var2)
	println(result)

}
