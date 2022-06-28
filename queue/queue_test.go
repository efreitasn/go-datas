package queue

import "testing"

func TestSize(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()

	size := q.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()

	value, hasValue := q.Peek()
	expectedValue := 2
	expectedHasValue := true

	if value != expectedValue {
		t.Errorf("got %v, want %v", value, expectedValue)
	}

	if hasValue != expectedHasValue {
		t.Errorf("got %v, want %v", hasValue, expectedHasValue)
	}

	// When the queue is empty
	q2 := New[int]()

	value, hasValue = q2.Peek()
	expectedValue = 0
	expectedHasValue = false

	if value != expectedValue {
		t.Errorf("got %v, want %v", value, expectedValue)
	}

	if hasValue != expectedHasValue {
		t.Errorf("got %v, want %v", hasValue, expectedHasValue)
	}
}
