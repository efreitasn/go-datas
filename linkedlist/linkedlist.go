// Package linkedlist provides functions to create/work with doubly linked lists of integers
package linkedlist

type node struct {
	next  *node
	prev  *node
	value int
}

// LinkedList is a doubly linked list of strings
type LinkedList struct {
	head *node
	tail *node
	size int
}

// New creates a doubly linked list
func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) insertFirstNode(value int) {
	n := &node{
		value: value,
	}

	ll.head = n
	ll.tail = n
	ll.size++
}

// Size returns the size of the linked list
func (ll *LinkedList) Size() int {
	return ll.size
}

// PeekBeginning returns the first element of the linked list (its head)
func (ll *LinkedList) PeekBeginning() (value int, ok bool) {
	if ll.head == nil {
		return 0, false
	}

	return ll.head.value, true
}

// InsertBeginning adds a value to the start of the linked list
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

// DeleteBeginning removes the first element of the linked list (its head)
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

// PeekEnd returns the last element of the linked list (its tail)
func (ll *LinkedList) PeekEnd() (value int, ok bool) {
	if ll.tail == nil {
		return 0, false
	}

	return ll.tail.value, true
}

// InsertEnd adds a value to the end of the linked list
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

// DeleteEnd removes the last element of the linked list (its tail)
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

// Contains returns whether the list contains the specified element
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
