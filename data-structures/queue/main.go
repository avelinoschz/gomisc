// This snippet implements a simple queue using slices.
// It is a quick reminder of FIFO behavior.
package main

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue() int {
	first := q.Items[0]
	q.Items = q.Items[1:]
	return first
}

func main() {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("queue:", q.Items)
	fmt.Println("dequeued:", q.Dequeue())
	fmt.Println("queue:", q.Items)
}
