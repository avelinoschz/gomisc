package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func (ll *LinkedList) Prepend(n *Node) {
	if ll.Head == nil {
		ll.Head = n
		return
	}

	next := ll.Head
	ll.Head = n

	ll.Head.Next = next
}

func (ll *LinkedList) Append(n *Node) {
	if ll.Head == nil {
		ll.Head = n
		return
	}

	currentNode := ll.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = n
}

// List() turns the values into a slice
// only for displaying purposes
func (ll *LinkedList) List() []int {
	if ll == nil {
		return nil
	}

	var list []int
	currentNode := ll.Head
	for currentNode != nil {
		list = append(list, currentNode.Data)
		currentNode = currentNode.Next
	}

	return list
}

func main() {
	ll := &LinkedList{
		Head: &Node{
			Data: 0,
		},
	}
	ll.Prepend(&Node{
		Data: 1,
	})
	ll.Prepend(&Node{
		Data: 2,
	})
	ll.Prepend(&Node{
		Data: 3,
	})
	ll.Append(&Node{
		Data: 4,
	})

	fmt.Printf("%+v\n", ll)
	fmt.Println(ll.List())
}
