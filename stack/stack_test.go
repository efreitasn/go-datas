package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	s.Push(9)
	s.Push(599)
	s.Push(3939)
	s.Pop()

	got := s.Peek()
	want := 599

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

	s.Push(93030)
	s.Push(1000)

	got := s.Size()
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
