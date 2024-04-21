package main

import "fmt"

// Queue structure implementation
// FIFO (first in first out) behavior
// items can be anything, this example uses
// `int` for simplicity
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
	fmt.Println(q)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)

	dequeued := q.Dequeue()
	fmt.Println(q)
	fmt.Println(dequeued)
}
