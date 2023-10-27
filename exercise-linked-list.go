package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.

type Node[T any] struct {
	next *Node[T]
	val  T
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (ls *List[T]) Append(val T) *Node[T] {
	new := Node[T]{val: val}
	if ls.head == nil {
		ls.head = &new
	} else {
		ls.tail.next = &new
	}

	ls.tail = &new
	ls.size++
	return &new
}

func (ls List[T]) PrintAll() {
	if ls.head != nil {
		nextNode := *ls.head
		for {
			fmt.Print(nextNode.val," -> ")
			if nextNode.next == nil {
				break
			}
			nextNode = *nextNode.next
		}
	}

}

func main() {
	ls := List[int]{}
	
	ls.Append(3)
	ls.Append(4)
	ls.Append(5)
	ls.Append(10)

	fmt.Printf("Size : %v, Head: %v, Tail: %v\n" , ls.size,ls.head,ls.tail)
	ls.PrintAll()
}
