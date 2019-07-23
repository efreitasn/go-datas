package linkedlist

import "testing"

func TestInsertEnd(t *testing.T) {
	ll := New()

	ll.InsertEnd("thing1")
	ll.InsertEnd("thing2")
	ll.InsertEnd("thing3")

	tail, _ := ll.PeekEnd()
	expectedTail := "thing3"

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	size := ll.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestDeleteEnd(t *testing.T) {
	ll := New()

	ll.InsertEnd("foo")
	ll.InsertEnd("bar")
	ll.InsertEnd("foobar")
	ll.DeleteEnd()

	tail, _ := ll.PeekEnd()
	expectedTail := "bar"

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}
}
