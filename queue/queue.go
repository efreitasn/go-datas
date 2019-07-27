// Package queue provides functions to create/work with queues of ints.
package queue

import (
	"github.com/efreitasn/go-datas/linkedlist"
)

// Queue is a queue of ints.
type Queue struct {
	items *linkedlist.LinkedList
}

// New creates a queue of ints.
func New() *Queue {
	items := linkedlist.New()

	return &Queue{
		items: items,
	}
}

// Size returns the size of the queue.
func (q *Queue) Size() int {
	return q.items.Size()
}

// Peek returns the value at the start of the queue.
func (q *Queue) Peek() (value int, hasValue bool) {
	return q.items.PeekBeginning()
}

// Enqueue inserts an value at the end of the queue.
func (q *Queue) Enqueue(value int) {
	q.items.InsertEnd(value)
}

// Dequeue removes the value at the start of the queue.
func (q *Queue) Dequeue() {
	q.items.DeleteBeginning()
}
