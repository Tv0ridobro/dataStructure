// Package queue implements a queue.
// See https://en.wikipedia.org/wiki/Queue_(abstract_data_type) for more details.
package queue

import "github.com/Tv0ridobro/data-structure/list"

// Queue represents a queue.
// Zero value of Stack is invalid stack, should be used only with New().
type Queue[T any] struct {
	list *list.List[T]
}

// New returns an initialized queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{list: list.New[T]()}
}

// Enqueue adds element to the end of queue.
func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

// Dequeue removes element from head of queue.
func (q *Queue[T]) Dequeue() T {
	return q.list.PopFront()
}

// Back returns last element of queue.
func (q *Queue[T]) Back() T {
	return q.list.Back()
}

// Front returns first element of queue.
func (q *Queue[T]) Front() T {
	return q.list.Front()
}

// Size returns size of queue.
func (q *Queue[T]) Size() int {
	return q.list.Len()
}

// All returns all elements from the queue.
func (q *Queue[T]) All() []T {
	return q.list.All()
}
