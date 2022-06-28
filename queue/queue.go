// Package queue provides functions to create/work with generic queues.
package queue

import (
	"github.com/efreitasn/go-datas/linkedlist"
)

// Queue is a queue of T.
type Queue[T any] struct {
	items *linkedlist.LinkedList[T]
}

// New creates a queue of ints.
func New[T any]() *Queue[T] {
	items := linkedlist.New[T]()

	return &Queue[T]{
		items: items,
	}
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() int {
	return q.items.Size()
}

// Peek returns the value at the start of the queue.
func (q *Queue[T]) Peek() (value T, hasValue bool) {
	return q.items.PeekBeginning()
}

// Enqueue inserts an value at the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.items.InsertEnd(value)
}

// Dequeue removes the value at the start of the queue.
func (q *Queue[T]) Dequeue() {
	q.items.DeleteBeginning()
}
