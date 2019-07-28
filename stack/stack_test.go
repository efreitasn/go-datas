package stack

import "testing"

func TestStack(t *testing.T) {
	s := New(9, 599)

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
	s := New()

	s.Push(30)
	s.Push(50)

	value, hasValue := s.Pop()
	expectedValue := 50
	expectedHasValue := true

	if value != expectedValue {
		t.Errorf("got %v, want %v", value, expectedValue)
	}

	if hasValue != expectedHasValue {
		t.Errorf("got %v, want %v", hasValue, expectedHasValue)
	}

	// Empty stack
	s = New()

	value, hasValue = s.Pop()
	expectedValue = 0
	expectedHasValue = false

	if value != expectedValue {
		t.Errorf("got %v, want %v", value, expectedValue)
	}

	if hasValue != expectedHasValue {
		t.Errorf("got %v, want %v", hasValue, expectedHasValue)
	}
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
