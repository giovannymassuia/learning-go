package main

import "fmt"

func main() {

	channel := make(chan string)

	go send("Hello World", channel)
	read(channel)
}

// the `ch <-` syntax means that the channel is only allowed to send data
func send(value string, ch chan<- string) {
	ch <- value
}

// the `<-ch` syntax means that the channel is only allowed to receive data
func read(ch <-chan string) {
	fmt.Println(<-ch)
}
