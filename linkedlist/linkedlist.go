// Package linkedlist provides functions to create/work with generic doubly linked lists.
package linkedlist

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	next  *node[T]
	prev  *node[T]
	value T
}

// LinkedList is a doubly linked list of T.
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// New creates a doubly linked list.
// It panics if the given comparing function is nil.
func New[T any](values ...T) *LinkedList[T] {
	ll := &LinkedList[T]{}

	for _, v := range values {
		ll.InsertEnd(v)
	}

	return ll
}

func (ll *LinkedList[T]) insertFirstNode(value T) {
	n := &node[T]{
		value: value,
	}

	ll.head = n
	ll.tail = n
	ll.size++
}

// Size returns the size of the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// PeekBeginning returns the first value of the linked list (its head).
func (ll *LinkedList[T]) PeekBeginning() (value T, hasValue bool) {
	if ll.head == nil {
		return *new(T), false
	}

	return ll.head.value, true
}

// InsertBeginning adds a value to the start of the linked list.
func (ll *LinkedList[T]) InsertBeginning(value T) {
	n := &node[T]{
		value: value,
	}

	if ll.Size() == 0 {
		ll.insertFirstNode(value)

		return
	}

	ll.head.prev = n
	n.next = ll.head
	ll.head = n

	ll.size++
}

// DeleteBeginning removes the first value of the linked list (its head).
func (ll *LinkedList[T]) DeleteBeginning() {
	if ll.Size() == 0 {
		return
	}

	if ll.Size() == 1 {
		ll.head = nil
		ll.tail = nil
		ll.size = 0

		return
	}

	ll.head.next.prev = nil
	ll.head = ll.head.next

	ll.size--
}

// PeekEnd returns the last value of the linked list (its tail).
func (ll *LinkedList[T]) PeekEnd() (value T, hasValue bool) {
	if ll.tail == nil {
		return *new(T), false
	}

	return ll.tail.value, true
}

// InsertEnd adds a value to the end of the linked list.
func (ll *LinkedList[T]) InsertEnd(value T) {
	n := &node[T]{
		value: value,
	}

	if ll.Size() == 0 {
		ll.insertFirstNode(value)

		return
	}

	ll.tail.next = n
	n.prev = ll.tail
	ll.tail = n

	ll.size++
}

// DeleteEnd removes the last value of the linked list (its tail).
func (ll *LinkedList[T]) DeleteEnd() {
	if ll.Size() == 0 {
		return
	}

	if ll.Size() == 1 {
		ll.tail = nil
		ll.head = nil
		ll.size = 0

		return
	}

	ll.tail.prev.next = nil
	ll.tail = ll.tail.prev
	ll.size--
}

// Traverse traverses the linked list by calling a callback at every encountered value.
// If the callback returns false, then the traversing is stopped.
func (ll *LinkedList[T]) Traverse(fromBeginning bool, cb func(v T) bool) {
	var n *node[T]

	if fromBeginning {
		n = ll.head
	} else {
		n = ll.tail
	}

	for n != nil {
		if !cb(n.value) {
			return
		}

		if fromBeginning {
			n = n.next
		} else {
			n = n.prev
		}
	}
}

// String returns the string representation of the linked list
func (ll *LinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("LinkedList{")

	switch {
	case ll.Size() == 0:
		sb.WriteString("}")
	default:
		first := true

		ll.Traverse(true, func(v T) bool {
			if !first {
				sb.WriteString(", ")
			}

			sb.WriteString(fmt.Sprint(v))

			first = false

			return true
		})

		sb.WriteString("}")
	}

	return sb.String()
}
