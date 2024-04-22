// kantan-coding https://www.youtube.com/watch?v=qyM8Pi1KiiM

package main

import (
	"fmt"
)

func main() {
	// using buffered channels, makes the usage of goroutines and channels async
	// this is because the sending goroutine can place multiple data to the channel
	// and continue it's work. if the channel is full, then the sending goroutine
	// will be block until there is more space in the channel
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, c := range chars {
		select {
		case charChannel <- c:
		}
		// in this specific case, a simple send/receive can be used
		// charChannel <- c
	}

	// because there is a range below, we need to explicitly `close` the channel
	// if not a deadlock will be caused, because will range forever
	close(charChannel)

	// we can iterate a channel with `for range` to retrieve its data
	for data := range charChannel {
		fmt.Println(data)
	}
}
