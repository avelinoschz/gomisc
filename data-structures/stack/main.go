package main

import "fmt"

// Stack structure implementation
// LIFO (last in first out) behavior
// items can be anything, this example uses
// `int` for simplicity
type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.Items) - 1
	lastItem := s.Items[lastIndex]
	// this syntax is not including last elem
	s.Items = s.Items[:lastIndex]

	return lastItem
}

func main() {
	// stacky := &Stack{
	// 	items: nil,
	// }
	// fmt.Println(stacky)
	// &{[]}

	// var stacky Stack
	// fmt.Println(stacky)
	// &{[]}

	// s := &Stack{
	// 	items: []int{},
	// }
	// fmt.Println(stacky)
	// &{[]}

	// all above declarations work
	// even if you would think they should
	// have a nil pointer error
	s := &Stack{}
	fmt.Println(s)

	s.Push(100)
	s.Push(200)
	s.Push(300)
	fmt.Println(s)

	popped := s.Pop()
	fmt.Println(s)
	fmt.Println(popped)
}
