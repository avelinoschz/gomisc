package main

import (
	"fmt"
	"time"
)

// primitives: goroutines, channels, select
func main() {
	// unbuffered channel
	// this makes -in a way- the usage of channel synchronous
	// because the sender turns to a waiting state (blocked)
	// until the receiver, receives the data
	msgChan := make(chan string)

	// async goroutine
	// anonymous function
	go func() {
		time.Sleep(time.Second * 3)
		msgChan <- "data string"
	}()
	fmt.Println("receiving message...")

	// blocking operation
	// waiting for data in the channel
	msg := <-msgChan

	fmt.Println("message received:", msg)
}
