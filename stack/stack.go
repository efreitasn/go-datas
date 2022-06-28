// Package stack provides functions to create/work with generic stacks.
package stack

// Stack is a stack of T.
type Stack[T any] struct {
	items []T
}

// New creates a stack of ints.
func New[T any](values ...T) *Stack[T] {
	s := &Stack[T]{}

	for _, v := range values {
		s.Push(v)
	}

	return s
}

// Peek returns the last value of the stack.
func (s *Stack[T]) Peek() (value T, hasValue bool) {
	if len(s.items) == 0 {
		return *new(T), false
	}

	return s.items[len(s.items)-1], true
}

// Push adds an value to the end of the stack.
func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

// Pop removes and returns the last value of the stack.
func (s *Stack[T]) Pop() (value T, hasValue bool) {
	value, hasValue = s.Peek()

	if hasValue {
		s.items = s.items[0:(len(s.items) - 1)]
	}

	return value, hasValue
}

// Size returns the size of the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}
