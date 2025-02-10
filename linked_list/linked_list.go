package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	data T
	next *LinkedList[T]
}

func NewLinkedList[T any](data ...T) *LinkedList[T] {
	l := &LinkedList[T]{}
	for _, d := range data {
		l.Append(d)
	}
	return l
}

func (l *LinkedList[T]) Append(data T) {
	if l.next == nil {
		l.next = &LinkedList[T]{data: data}
	} else {
		l.next.Append(data)
	}
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	i := 0
	for i != index {
		// Check if there is a next node
		if l.next == nil {
			var zero T
			return zero, fmt.Errorf("Index out of range")
		}
		l = l.next
		i++
	}

	return l.data, nil
}

func (l *LinkedList[T]) String() string {
	if l == nil {
		return ""
	}
	sb := ""
	currentNode := l
	for currentNode.next != nil {
		sb += fmt.Sprintf(", %v", currentNode.data)
		currentNode = currentNode.next
	}
	return fmt.Sprintf("LinkedList(%v)", sb[2:])
}

func (l *LinkedList[T]) Length() int {
	if l == nil {
		return 0
	}
	return 1 + l.next.Length()
}

func main() {
	l := NewLinkedList(1, 2, 3, 4, 5)
	fmt.Println(l)

}
