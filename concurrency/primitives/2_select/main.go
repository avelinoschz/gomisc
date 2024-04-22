package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)
	anotherChan := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		msgChan <- "data string 1"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		anotherChan <- "data string 2"
	}()

	fmt.Println("receiving message...")

	// select let the goroutine wait on multiple communication operations
	// in this case, let the main function wait on messages from multiple channels
	// if messages are at the same time, select will choose one randomly
	select {
	case msg := <-msgChan:
		fmt.Println(msg)
	case anotherMsg := <-anotherChan:
		fmt.Println(anotherMsg)
	}
}
