// This snippet fills a channel, closes it, and reads it with for-range.
// Closing the channel matters here because range stops only after the channel is closed.
// Source: kantan-coding https://www.youtube.com/watch?v=qyM8Pi1KiiM
package main

import "fmt"

func main() {
	letters := make(chan string, 3)
	values := []string{"a", "b", "c"}

	for _, value := range values {
		letters <- value
	}

	close(letters)

	for letter := range letters {
		fmt.Println(letter)
	}
}
