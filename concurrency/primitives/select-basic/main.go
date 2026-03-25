// This snippet shows how to use select to wait on multiple channels.
// It continues with whichever channel becomes ready first.
package main

import (
	"fmt"
	"time"
)

func main() {
	slowChan := make(chan string)
	fastChan := make(chan string)

	go func() {
		time.Sleep(400 * time.Millisecond)
		slowChan <- "slow message"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		fastChan <- "fast message"
	}()

	fmt.Println("waiting for the first message...")

	select {
	case msg := <-slowChan:
		fmt.Println(msg)
	case msg := <-fastChan:
		fmt.Println(msg)
	}
}
