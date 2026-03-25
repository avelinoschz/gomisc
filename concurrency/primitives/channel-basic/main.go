// This snippet shows the basic usage of an unbuffered channel.
// The send blocks until another goroutine receives the value.
package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		msgChan <- "hello from another goroutine"
	}()

	fmt.Println("receiving message...")
	msg := <-msgChan
	fmt.Println("message received:", msg)
}
