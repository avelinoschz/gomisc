package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("sorted array:", arr)

	shuffled := shuffle(arr)

	fmt.Println("shuffled array:", shuffled)
}

func shuffle(arr []int) []int {
	// deprecated Go 1.20
	// rand.Seed(time.Now().UnixNano())

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	length := len(arr)

	for i := length - 1; i > 0; i-- {
		randIdx := r.Intn(i + 1)

		fmt.Printf("idx to swap: %d and %d\n", i, randIdx)
		fmt.Printf("vals to swap: %d and %d\n", arr[i], arr[randIdx])

		arr[randIdx], arr[i] = arr[i], arr[randIdx]

		fmt.Println(arr)
		fmt.Println("")
	}

	return arr
}
