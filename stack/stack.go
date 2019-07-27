// Package stack provides functions to create/work with stacks of ints.
package stack

// Stack is a stack of ints
type Stack struct {
	items []int
}

// New creates a stack
func New() *Stack {
	return &Stack{}
}

// Peek returns the last element of the stack
func (s *Stack) Peek() int {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}

	return 0
}

// Push adds an element to the end of the stack
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop removes the last element of the stack
func (s *Stack) Pop() {
	if s.Size() > 0 {
		s.items = s.items[0:(len(s.items) - 1)]
	}
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}
