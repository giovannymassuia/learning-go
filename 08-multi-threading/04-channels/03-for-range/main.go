package main

import "fmt"

func main() {
	channel := make(chan int)

	go publish(channel)
	consume(channel)
}

func consume(channel chan int) {
	for i := range channel {
		fmt.Printf("Received %d\n", i)
	}
}

func publish(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel) // need to close channel, otherwise the program will hang
}
