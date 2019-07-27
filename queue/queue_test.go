package queue

import "testing"

func TestSize(t *testing.T) {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()

	size := q.Size()
	expectedSize := 4

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestPeek(t *testing.T) {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()

	head, hasHead := q.Peek()
	expectedHead := 2
	expectedHasHead := true

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}

	// When the queue is empty
	q2 := New()

	head, hasHead = q2.Peek()
	expectedHead = 0
	expectedHasHead = false

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}
