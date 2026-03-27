// This snippet implements the Fisher-Yates algorithm to shuffle a slice.
// It swaps each element with a random earlier position.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := shuffle(append([]int(nil), values...))

	fmt.Println("original:", values)
	fmt.Println("shuffled:", shuffled)
}

func shuffle(values []int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(values) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		values[i], values[j] = values[j], values[i]
	}

	return values
}
