package queue

import "testing"

func TestSize(t *testing.T) {
	q := New()

	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("foobar")
	q.Enqueue("thing")
	q.Dequeue()

	size := q.Size()
	expectedSize := 4

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestPeek(t *testing.T) {
	q := New()

	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("foobar")
	q.Enqueue("thing")
	q.Dequeue()

	head, hasHead := q.Peek()
	expectedHead := "bar"
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
	expectedHead = ""
	expectedHasHead = false

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}
