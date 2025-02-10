package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListLength(t *testing.T) {
	l1 := NewLinkedList[int]()
	assert.Equal(t, 0, l1.Length())

	l2 := NewLinkedList(1, 2, 3, 4, 5)
	assert.Equal(t, 5, l2.Length())
}
