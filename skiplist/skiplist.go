// Package skiplist implements a skiplist.
// See https://en.wikipedia.org/wiki/Skip_list for more details.
package skiplist

import (
	"math/rand"

	"golang.org/x/exp/constraints"

	"github.com/Tv0ridobro/data-structure/math"
	"github.com/Tv0ridobro/data-structure/stack"
)

// SkipList represents a skiplist.
// Zero value of SkipList is invalid skiplist, should be used only with New() or NewWithProbability().
type SkipList[T constraints.Ordered] struct {
	p    float64
	head *Node[T]
}

// NewWithProbability returns an initialized skiplist with given probability.
func NewWithProbability[T constraints.Ordered](p float64) *SkipList[T] {
	n := &Node[T]{
		value: math.SmallestValue[T](),
	}
	return &SkipList[T]{
		p:    p,
		head: n,
	}
}

// New returns an initialized skiplist with probability = 0.5.
func New[T constraints.Ordered]() *SkipList[T] {
	return NewWithProbability[T](0.5)
}

// Insert inserts value in a skiplist.
func (s *SkipList[T]) Insert(value T) {
	it, nodes := s.find(value)
	if it.value == value {
		return
	}
	below := &Node[T]{ // insert ones
		next:  it.next,
		value: value,
	}
	it.next = below
	for rand.Float64() <= s.p {
		below = &Node[T]{
			below: below,
			value: value,
		}
		if nodes.Size() == 0 {
			it = &Node[T]{
				next:  below,
				below: s.head,
				value: s.head.value,
			}
			s.head = it
			continue
		}
		it = nodes.Pop()
		below.next = it.next
		it.next = below
	}
}

// Remove removes value from skiplist.
// Returns true if skiplist contained given value, false otherwise.
func (s *SkipList[T]) Remove(value T) bool {
	it, nodes := s.find(value)
	if it.value != value {
		return false
	}
	var last *Node[T]
	if nodes.Size() == 0 {
		last = s.head
	} else {
		last = nodes.Pop()
		last = last.below
	}
	for it != nil {
		for last.next != it {
			last = last.next
		}
		last.next = it.next
		last = last.below
		it = it.below
	}
	s.normalize()
	return true
}

// normalize removes all layers without elements.
func (s *SkipList[T]) normalize() {
	for s.head.HasBelow() && !s.head.HasNext() {
		s.head = s.head.below
	}
}

// Find returns Node that contains value if its exist,
// closest otherwise.
func (s *SkipList[T]) Find(value T) *Node[T] {
	it, _ := s.find(value)
	return it
}

// find returns Node that contains value if its exist,
// closest one otherwise.
// Returns stack of nodes on path to returned node.
func (s *SkipList[T]) find(value T) (*Node[T], *stack.Stack[*Node[T]]) {
	nodes := stack.New[*Node[T]]()
	it := s.head
Loop:
	for { // find where to place it
		if it.value == value {
			return it, nodes
		}
		switch {
		case it.HasNext() && it.next.value <= value:
			it = it.next
		case it.HasBelow():
			nodes.Push(it)
			it = it.below
		default:
			break Loop
		}
	}
	return it, nodes
}

// Contains returns true if skiplist contains given value, false otherwise.
func (s *SkipList[T]) Contains(value T) bool {
	return s.Find(value).value == value
}

// All returns all elements from the skiplist.
func (s *SkipList[T]) All() []T {
	it := s.head
	if it == nil {
		return []T{}
	}
	for it.HasBelow() {
		it = it.Below()
	}
	a := make([]T, 0)
	it = it.Next()
	for it != nil {
		a = append(a, it.Value())
		it = it.Next()
	}
	return a
}
