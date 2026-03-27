// This snippet implements a minimal stack using slices.
// It is a quick reference for Push and Pop with LIFO behavior.
package main

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.Items) - 1
	lastItem := s.Items[lastIndex]
	s.Items = s.Items[:lastIndex]
	return lastItem
}

func main() {
	s := &Stack{}
	s.Push(100)
	s.Push(200)
	s.Push(300)

	fmt.Println("stack:", s.Items)
	fmt.Println("popped:", s.Pop())
	fmt.Println("stack:", s.Items)
}
