// Package splaytree implements a splay tree.
// See https://en.wikipedia.org/wiki/Splay_tree for more details.
package splaytree

import (
	"golang.org/x/exp/constraints"

	"github.com/Tv0ridobro/data-structure/math"
)

// SplayTree represents a splay tree.
// Zero value of SplayTree is empty splay tree.
type SplayTree[T any] struct {
	comp func(T, T) int
	root *node[T]
}

// New returns an initialized splay tree.
func New[T constraints.Ordered]() *SplayTree[T] {
	return &SplayTree[T]{
		comp: math.Comparator[T](),
	}
}

// NewWithComparator returns an initialized splay tree using given comparator.
func NewWithComparator[T any](comp func(T, T) int) *SplayTree[T] {
	return &SplayTree[T]{
		comp: comp,
	}
}

// Insert inserts value in a tree.
func (s *SplayTree[T]) Insert(value T) {
	n := &node[T]{
		value: value,
		size:  1,
	}
	if s.root == nil {
		s.root = n
		return
	}
	nn := s.root.find(value, s.comp)
	l, r := split(nn, value, s.comp)
	setLeft(n, l)
	setRight(n, r)
	s.root = n
	s.root.recalculateSize()
}

// Contains returns true if tree contains given value, false otherwise.
func (s *SplayTree[T]) Contains(value T) bool {
	if s.root == nil {
		return false
	}
	s.root = s.root.find(value, s.comp)
	return s.comp(s.root.value, value) == 0
}

// Size returns size of the tree.
func (s *SplayTree[T]) Size() int {
	if s.root == nil {
		return 0
	}
	return s.root.size
}

// Remove removes value from tree.
// returns true if tree contained given value, false otherwise.
func (s *SplayTree[T]) Remove(value T) bool {
	if s.root == nil {
		return false
	}
	s.root = s.root.find(value, s.comp)
	if s.comp(s.root.value, value) != 0 {
		return false
	}
	l, r := s.root.left, s.root.right
	s.root.left, s.root.right = nil, nil
	if l != nil {
		l.parent = nil
	}
	if r != nil {
		r.parent = nil
	}
	s.root = merge(l, r)
	return true
}

// Kth returns kth greatest element.
func (s *SplayTree[T]) Kth(i int) T {
	if i > s.Size() || i < 0 {
		var empty T
		return empty
	}
	n := s.root.kth(i)
	return n.value
}

// Sub returns elements [l, r) in ascending order.
func (s *SplayTree[T]) Sub(l, r int) []T {
	if l >= r || l < 0 {
		return nil
	}
	if r > s.Size() {
		r = s.Size()
	}
	sl := make([]T, r-l)
	n := s.root.kth(l)
	sl[0] = n.value
	for i := 1; i < r-l; i++ {
		n = n.next()
		sl[i] = n.value
	}
	return sl
}

// All returns all elements from tree.
// Returned slice is sorted.
func (s *SplayTree[T]) All() []T {
	if s.root == nil {
		return nil
	}
	d := make([]T, s.Size())
	s.root.getAll(d)
	return d
}
