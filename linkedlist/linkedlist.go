// Package linkedlist provides functions to create/work with doubly linked lists of strings
package linkedlist

type node struct {
	next  *node
	prev  *node
	value string
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

// Last returns the last element of the linked list (its tail)
func (ll *LinkedList) Last() (value string, ok bool) {
	if ll.tail == nil {
		return "", false
	}

	return ll.tail.value, true
}

// AddLast adds a value to the end of the linked list
func (ll *LinkedList) AddLast(value string) {
	n := &node{
		value: value,
	}

	if ll.tail != nil {
		ll.tail.next = n
		n.prev = ll.tail
	}

	ll.tail = n
	ll.size++
}

// DeleteLast removes a value starting from the linked list's tail
func (ll *LinkedList) DeleteLast(value string) {
	n := ll.tail

	for n != nil {
		if n.value == value {
			if n.prev != nil {
				n.prev.next = n.next
			}

			if n.next != nil {
				n.next.prev = n.prev
			}

			if n == ll.tail {
				ll.tail = n.prev
			}

			return
		}

		n = n.prev
	}
}

// Size returns the size of the linked list
func (ll *LinkedList) Size() int {
	return ll.size
}
