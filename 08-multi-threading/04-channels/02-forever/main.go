package main

// thread 1
func main() {
	forever := make(chan bool) // create channel, channel is empty

	// without this goroutine, the program will exit with deadlock error
	go func() {
		println("Hello World")
		forever <- true // send message to channel, channel is full
	}()

	<-forever // receive message from channel, channel is empty
}
