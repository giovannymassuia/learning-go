package main

func main() {

	// regular for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	// range for loop
	numbers := []int{1, 2, 3}
	for i, v := range numbers {
		println(i, v)
	}

	// while loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// infinite loop
	//for {
	//	println("infinite")
	//}

}
