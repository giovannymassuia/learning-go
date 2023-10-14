package main

import (
	"fmt"
	"time"
)

// Thread main
func main() {
	startTime := time.Now()

	go task("A")
	go task("B")
	go func() {
		fmt.Println("Task C is running")
		time.Sleep(1 * time.Second)
		fmt.Println("Task C is done")
	}()

	time.Sleep(15 * time.Second)

	elapsed := time.Since(startTime)
	fmt.Printf(">> Total time taken: %s\n", elapsed)
}

func task(name string) {
	// log start time
	start := time.Now()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
	// log elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Task %s took %s\n", name, elapsed)
}
