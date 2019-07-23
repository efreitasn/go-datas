package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	s.Push("foobar")
	s.Push("foo")
	s.Push("bar")
	s.Pop()

	got := s.Peek()
	want := "foo"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPop(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("got a panic, want no panic")
		}
	}()

	s := New()

	s.Pop()
}

func TestSize(t *testing.T) {
	s := New()

	s.Push("foo")
	s.Push("bar")

	got := s.Size()
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
