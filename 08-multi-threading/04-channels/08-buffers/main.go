package main

func main() {
	ch := make(chan string, 2) // buffer size 2

	ch <- "Hello" // send message to channel, channel is not full yet
	ch <- "World" // send message to channel, channel is full

	println(<-ch) // receive message from channel, channel is not empty
	println(<-ch) // receive message from channel, channel is empty
}
