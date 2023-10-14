package main

import (
	"fmt"
	"sync"
	"time"
)

// Thread main
func main() {
	startTime := time.Now()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10) // 10 because we have 2 tasks and each task has 5 iterations

	go task("A", &waitGroup)
	go task("B", &waitGroup)

	waitGroup.Wait() // wait until all tasks are done

	elapsed := time.Since(startTime)
	fmt.Printf(">> Total time taken: %s\n", elapsed)
}

func task(name string, waitGroup *sync.WaitGroup) {
	// log start time
	start := time.Now()
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)

		waitGroup.Done()
	}
	// log elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Task %s took %s\n", name, elapsed)
}
