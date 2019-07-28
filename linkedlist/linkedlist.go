// Package linkedlist provides functions to create/work with doubly linked lists of ints.
package linkedlist

import "strconv"

type node struct {
	next  *node
	prev  *node
	value int
}

// LinkedList is a doubly linked list of ints.
type LinkedList struct {
	head *node
	tail *node
	size int
}

// New creates a doubly linked list.
func New(values ...int) *LinkedList {
	ll := &LinkedList{}

	for _, v := range values {
		ll.InsertEnd(v)
	}

	return ll
}

func (ll *LinkedList) insertFirstNode(value int) {
	n := &node{
		value: value,
	}

	ll.head = n
	ll.tail = n
	ll.size++
}

// Size returns the size of the linked list.
func (ll *LinkedList) Size() int {
	return ll.size
}

// PeekBeginning returns the first value of the linked list (its head).
func (ll *LinkedList) PeekBeginning() (value int, hasValue bool) {
	if ll.head == nil {
		return 0, false
	}

	return ll.head.value, true
}

// InsertBeginning adds a value to the start of the linked list.
func (ll *LinkedList) InsertBeginning(value int) {
	n := &node{
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
func (ll *LinkedList) DeleteBeginning() {
	if ll.Size() == 0 {
		return
	}

	if ll.Size() == 1 {
		ll.head = nil
		ll.tail = nil

		return
	}

	ll.head.next.prev = nil
	ll.head = ll.head.next
}

// PeekEnd returns the last value of the linked list (its tail).
func (ll *LinkedList) PeekEnd() (value int, hasValue bool) {
	if ll.tail == nil {
		return 0, false
	}

	return ll.tail.value, true
}

// InsertEnd adds a value to the end of the linked list.
func (ll *LinkedList) InsertEnd(value int) {
	n := &node{
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
func (ll *LinkedList) DeleteEnd() {
	if ll.Size() == 0 {
		return
	}

	if ll.Size() == 1 {
		ll.tail = nil
		ll.head = nil

		return
	}

	ll.tail.prev.next = nil
	ll.tail = ll.tail.prev
}

// Contains returns whether the list contains the specified value.
func (ll *LinkedList) Contains(value int) bool {
	n := ll.head

	for n != nil {
		if n.value == value {
			return true
		}

		n = n.next
	}

	return false
}

// Traverse traverses the linked list by calling a callback at every encountered value.
func (ll *LinkedList) Traverse(fromBeginning bool, cb func(v int)) {
	var n *node

	if fromBeginning {
		n = ll.head
	} else {
		n = ll.tail
	}

	for n != nil {
		cb(n.value)

		if fromBeginning {
			n = n.next
		} else {
			n = n.prev
		}
	}
}

// String returns the string representation of the linked list
func (ll *LinkedList) String() string {
	str := "LinkedList{"

	if ll.Size() == 0 {
		return str + "}"
	}

	ll.Traverse(true, func(v int) {
		str += strconv.Itoa(v) + ", "
	})

	return str[:len(str)-2] + "}"
}
