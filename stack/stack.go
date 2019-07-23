// Package stack provides functions to create/work with stacks of strings.
package stack

// Stack is a stack of strings
type Stack struct {
	items []string
}

// New creates a Stack of the specified size
func New() *Stack {
	var items []string

	return &Stack{
		items: items,
	}
}

// Peek returns the last element of the stack
func (s *Stack) Peek() string {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}

	return ""
}

// Push adds an element to the end of the stack
func (s *Stack) Push(val string) {
	s.items = append(s.items, val)
}

// Pop removes the last element of the stack
func (s *Stack) Pop() {
	if len(s.items) > 0 {
		s.items = s.items[0:(len(s.items) - 1)]
	}
}
