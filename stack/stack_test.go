package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	s.Push(9)
	s.Push(599)
	s.Push(3939)
	s.Pop()

	value, hasValue := s.Peek()
	expectedValue := 599
	expectedHasValue := true

	if value != expectedValue {
		t.Errorf("got %v, want %v", value, expectedValue)
	}

	if hasValue != expectedHasValue {
		t.Errorf("got %v, want %v", hasValue, expectedHasValue)
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
