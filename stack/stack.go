// Package stack provides functions to create/work with stacks of ints.
package stack

// Stack is a stack of ints.
type Stack struct {
	items []int
}

// New creates a stack of ints.
func New() *Stack {
	return &Stack{}
}

// Peek returns the last value of the stack.
func (s *Stack) Peek() (value int, hasValue bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// Push adds an value to the end of the stack.
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop removes the last value of the stack.
func (s *Stack) Pop() {
	if s.Size() > 0 {
		s.items = s.items[0:(len(s.items) - 1)]
	}
}

// Size returns the size of the stack.
func (s *Stack) Size() int {
	return len(s.items)
}
