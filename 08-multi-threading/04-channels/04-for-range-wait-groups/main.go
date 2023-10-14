package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(5)

	go publish(channel)
	go consume(channel, &waitGroup)

	waitGroup.Wait()

	// Using wait groups, the publisher and consumer threads are synchronized,
	// and can run in any order, and each in a different go routine.
}

func consume(channel chan int, waitGroup *sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Received %d\n", i)
		waitGroup.Done()
	}
}

func publish(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel) // need to close channel, otherwise the program will hang
}
