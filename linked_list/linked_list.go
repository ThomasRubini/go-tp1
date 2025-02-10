package main

import (
	"fmt"
)

type LinkedListNode[T any] struct {
	data T
	next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	first *LinkedListNode[T]
}

func NewLinkedList[T any](data ...T) *LinkedList[T] {
	l := &LinkedList[T]{}
	for _, d := range data {
		l.Append(d)
	}
	return l
}

// Returns null if list is empty
func (l *LinkedList[T]) lastNode() *LinkedListNode[T] {
	if l.first == nil {
		return nil
	}

	currentNode := l.first
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l *LinkedList[T]) Append(data T) {
	last := l.lastNode()
	if last == nil {
		l.first = &LinkedListNode[T]{data: data}
	} else {
		last.next = &LinkedListNode[T]{data: data}
	}
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	i := 0
	currentNode := l.first
	for currentNode != nil {
		if i == index {
			// Found
			return currentNode.data, nil
		}
		i++
		currentNode = currentNode.next
	}
	// i is the length of the list at this point
	var zero T
	if i >= index {
		// Out of bounds
		return zero, fmt.Errorf("Index out of bounds")
	} else {
		// Not found
		return zero, fmt.Errorf("Element not found")
	}
}

func (l *LinkedList[T]) String() string {
	if l == nil {
		return ""
	}
	sb := ""
	currentNode := l.first
	for currentNode != nil {
		sb += fmt.Sprintf(", %v", currentNode.data)
		currentNode = currentNode.next
	}
	return fmt.Sprintf("LinkedList(%v)", sb[2:])
}

func (l *LinkedList[T]) Length() int {
	if l.first == nil {
		return 0
	}

	i := 0
	currentNode := l.first
	for currentNode != nil {
		i++
		currentNode = currentNode.next
	}
	return i
}

func main() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(l)

}
