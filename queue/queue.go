// Package queue provides functions to create/work with queues of strings
package queue

import (
	"github.com/efreitasn/go-datas/linkedlist"
)

// Queue is a queue of strings
type Queue struct {
	items *linkedlist.LinkedList
}

// New creates a queue
func New() *Queue {
	items := linkedlist.New()

	return &Queue{
		items: items,
	}
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return q.items.Size()
}

// Peek returns the element at the start of the queue
func (q *Queue) Peek() (value string, ok bool) {
	return q.items.PeekBeginning()
}

// Enqueue inserts an element at the end of the queue
func (q *Queue) Enqueue(value string) {
	q.items.InsertEnd(value)
}

// Dequeue removes the element at the start of the queue
func (q *Queue) Dequeue() {
	q.items.DeleteBeginning()
}
