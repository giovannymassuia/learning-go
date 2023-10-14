package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id    int64
	value string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	// rabbitmq
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, value: "Hello from RabbitMQ"}
			c1 <- msg
			time.Sleep(time.Second)
			//c1 <- 1
		}
	}()

	// kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{id: i, value: "Hello from Kafka"}
			c2 <- msg
			time.Sleep(time.Second)
			//c2 <- 1
		}
	}()

	//for i := 0; i < 3; i++ {
	for {
		select {
		case msg := <-c1: // rabbitmq
			fmt.Printf("Msg from RabbitMQ: %v\n", msg)

		case msg := <-c2: // kafka
			fmt.Printf("Msg from Kafka: %v\n", msg)

		case <-time.After(3 * time.Second):
			println("timeout")

			//default:
			//	println("nothing ready")
		}
	}
}
