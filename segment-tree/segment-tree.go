// Package segmenttree implements a segment tree.
// See https://en.wikipedia.org/wiki/Segment_tree for more details.
package segmenttree

import (
	"github.com/Tv0ridobro/data-structure/math"
)

// SegmentTree represents a segment tree.
// Zero value of SegmentTree is invalid segment tree, should be used only with New().
type SegmentTree[T any] struct {
	op       func(T, T) T
	neutral  T
	elements []T
}

// New returns new SegmentTree.
// T should be monoid.
// op(op(a, b), c) = op(a, op(b, c)) and op(neutral, a) = op(a, neutral) = a.
func New[T any](elements []T, op func(T, T) T, neutral T) *SegmentTree[T] {
	d := math.NearestPowerOf2(len(elements))
	c := make([]T, d*2)
	copied := copy(c[d:], elements)
	for i := d + copied; i < 2*d; i++ {
		c[i] = neutral
	}
	st := SegmentTree[T]{
		op:       op,
		neutral:  neutral,
		elements: c,
	}
	st.build(1, 0, d-1)
	return &st
}

// build recursively builds segment tree.
// i is current index in tree,
// left and right are its bound.
func (s *SegmentTree[T]) build(i, left, right int) {
	if right == left {
		return
	}
	middle := (left + right) / 2
	s.build(2*i, left, middle)
	s.build(2*i+1, middle+1, right)
	s.elements[i] = s.op(s.elements[2*i], s.elements[2*i+1])
}

// Query returns result of operation on given segment.
func (s *SegmentTree[T]) Query(l, r int) T {
	return s.query(1, 0, len(s.elements)/2-1, l, r)
}

// query returns result of operation on given segment.
// i is current index in tree,
// l and r are its bound,
// left and right are initial values of query.
func (s *SegmentTree[T]) query(i, left, right, l, r int) T {
	if l > r {
		return s.neutral
	}
	if l == left && r == right {
		return s.elements[i]
	}
	middle := (left + right) / 2
	return s.op(
		s.query(i*2, left, middle, l, math.Min(r, middle)),
		s.query(i*2+1, middle+1, right, math.Max(l, middle+1), r),
	)
}

// Modify modifies value at given index.
func (s *SegmentTree[T]) Modify(i int, value T) {
	s.modify(1, i, 0, len(s.elements)/2-1, value)
}

// modify modifies value at given index
// i is current index in tree,
// left and right are its bound,
// ind is initial index.
func (s *SegmentTree[T]) modify(i, ind, left, right int, value T) {
	if right == left {
		s.elements[i] = value
		return
	}
	if middle := (left + right) / 2; ind <= middle {
		s.modify(i*2, ind, left, middle, value)
	} else {
		s.modify(i*2+1, ind, middle+1, right, value)
	}
	s.elements[i] = s.op(s.elements[i*2], s.elements[i*2+1])
}
