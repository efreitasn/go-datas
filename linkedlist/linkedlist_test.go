package linkedlist

import "testing"

func TestLast(t *testing.T) {
	ll := New()

	ll.AddLast("foo")
	ll.AddLast("bar")

	got, _ := ll.Last()
	want := "bar"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDeleteLast(t *testing.T) {
	ll := New()

	ll.AddLast("foo")
	ll.AddLast("bar")
	ll.AddLast("foobar")

	ll.DeleteLast("foobar")

	got, _ := ll.Last()
	want := "bar"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
