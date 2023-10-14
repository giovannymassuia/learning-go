package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	data := make(chan int)
	workersQty := 1_000_000

	// init workers
	for i := 0; i < workersQty; i++ {
		go worker(i, data)
	}

	// send data to workers
	for i := 0; i < 10_000_000; i++ {
		data <- i
	}

	fmt.Printf("Total time: %s\n", time.Since(startTime))
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
