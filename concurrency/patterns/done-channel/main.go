// This snippet uses a done channel to stop a goroutine explicitly.
// Closing the channel broadcasts the stop signal to every receiver.
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go doWork(done)

	time.Sleep(350 * time.Millisecond)
	close(done)

	// Give the worker a moment to print its stop message before main exits.
	time.Sleep(50 * time.Millisecond)
}

func doWork(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("stopping worker")
			return
		default:
			fmt.Println("working...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
