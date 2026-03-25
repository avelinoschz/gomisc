// This snippet implements a simple linked list with prepend, append, and traversal.
// It keeps only the basic operations needed to remember the shape of the structure.
package main

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func (ll *LinkedList) Prepend(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		return
	}

	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) Append(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}

func (ll *LinkedList) List() []int {
	var list []int
	current := ll.Head
	for current != nil {
		list = append(list, current.Data)
		current = current.Next
	}

	return list
}

func main() {
	ll := &LinkedList{}
	ll.Append(&Node{Data: 2})
	ll.Prepend(&Node{Data: 1})
	ll.Append(&Node{Data: 3})

	fmt.Println(ll.List())
}
