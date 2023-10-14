package main

// thread 1
func main() {

	channel := make(chan string)

	// thread 2
	go func() {
		channel <- "Hello World" // send message to channel, channel is full
	}()

	// thread 1
	message := <-channel // receive message from channel, channel is empty
	println(message)
}
