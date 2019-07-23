package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()

	stack.Push("foobar")
	stack.Push("foo")
	stack.Push("bar")
	stack.Pop()

	r := stack.Peek()
	want := "foo"

	if r != want {
		t.Errorf("got %v, want %v", r, want)
	}
}

func TestPop(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("got a panic, want no panic")
		}
	}()

	stack := New()

	stack.Pop()
}
