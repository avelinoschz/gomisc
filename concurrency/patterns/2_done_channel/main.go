package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 2)

	//this esentially sends a zero value to the chan
	close(done)

	// this ends the execution of the goroutine too
	// done <- true
}

// here is receiving a read only channel
// that's why it has the receiving data syntax
// the for-select pattern works here because we are waiting for a signal
// while doing the default job of the goroutine
func doWork(done <-chan bool) {
	for {
		select {
		case <-done: // whenever it receives any value, finishes
			return
		default:
			fmt.Println("doing work infinitely")
		}
	}
}
