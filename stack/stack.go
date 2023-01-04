// Package stack implements a stack.
// See https://en.wikipedia.org/wiki/Stack_(abstract_data_type) for more details.
package stack

// Stack represents a stack.
// Zero value of Stack is empty stack.
type Stack[T any] struct {
	array []T
}

// New returns an initialized stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{[]T{}}
}

// Push adds element.
func (s *Stack[T]) Push(value T) {
	s.array = append(s.array, value)
}

// Pop removes the most recently added element.
func (s *Stack[T]) Pop() T {
	if len(s.array) == 0 {
		var zero T
		return zero
	}
	v := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return v
}

// Peek returns element on top of stack.
func (s *Stack[T]) Peek() T {
	if len(s.array) == 0 {
		var zero T
		return zero
	}
	return s.array[len(s.array)-1]
}

// Size returns size of the stack.
func (s *Stack[T]) Size() int {
	return len(s.array)
}

// All returns all elements from the stack.
func (s *Stack[T]) All() []T {
	all := make([]T, len(s.array))
	copy(all, s.array)
	return all
}
