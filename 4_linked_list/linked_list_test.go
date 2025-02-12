package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func E1[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func TestListLength(t *testing.T) {
	l1 := NewLinkedList[int]()
	assert.Equal(t, 0, l1.Length())

	l2 := NewLinkedList(1, 2, 3, 4, 5)
	assert.Equal(t, 5, l2.Length())
}

func TestGet(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4, 5)
	assert.Equal(t, 1, E1(l.Get(0)))
	assert.Equal(t, 2, E1(l.Get(1)))
	assert.Equal(t, 3, E1(l.Get(2)))
	assert.Equal(t, 4, E1(l.Get(3)))
	assert.Equal(t, 5, E1(l.Get(4)))
}
