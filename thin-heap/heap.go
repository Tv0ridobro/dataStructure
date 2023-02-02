// Package thin_heap implements a thin heap.
// See http://www.cs.tau.ac.il/~haimk/papers/newthin1.pdf for more details.
package thin_heap

import (
	"golang.org/x/exp/constraints"

	mathx "github.com/Tv0ridobro/data-structure/math"
)

// ThinHeap represents a thin heap.
// Zero value of ThinHeap is empty ThinHeap.
type ThinHeap[T constraints.Ordered] struct {
	first *Node[T]
	last  *Node[T]
	size  int
}

// New returns empty ThinHeap.
func New[T constraints.Ordered]() *ThinHeap[T] {
	return &ThinHeap[T]{
		first: nil,
		last:  nil,
	}
}

// Min returns minimal element.
func (h *ThinHeap[T]) Min() T {
	if h.first != nil {
		return h.first.value
	}
	var zero T
	return zero
}

// DeleteMin delete minimal element from heap and return it.
func (h *ThinHeap[T]) DeleteMin() T {
	if h.first == nil {
		var zero T
		return zero
	}
	h.size--

	tmp := h.first
	h.first = h.first.right
	if h.first == nil {
		h.last = nil
	}

	x := tmp.child
	for x != nil {
		if x.isThin() {
			x.rank--
		}
		x.left = nil
		next := x.right
		x.right = nil
		h.insert(x)
		x = next
	}

	x = h.first
	rangs := make([]*Node[T], mathx.NearestPowerOf2(h.size))
	for x != nil {
		next := x.right
		for node := rangs[x.rank]; node != nil; node = rangs[x.rank] {
			if node.value < x.value {
				node, x = x, node
			}
			node.right = x.child
			if x.child != nil {
				x.child.left = node
			}
			node.left = x
			x.child = node

			rangs[x.rank] = nil
			x.rank++
		}
		rangs[x.rank] = x
		x = next
	}

	n := New[T]()
	for _, e := range rangs {
		if e != nil {
			e.left = nil
			e.right = nil
			n.insert(e)
		}
	}
	h.first, h.last = n.first, n.last

	return tmp.value
}

// Insert inserts elements into heap.
func (h *ThinHeap[T]) Insert(element T) {
	n := &Node[T]{value: element}
	h.size++
	h.insert(n)
}

// Size returns size of heap.
func (h *ThinHeap[T]) Size() int {
	return h.size
}

func (h *ThinHeap[T]) insert(n *Node[T]) {
	if h.first == nil {
		h.first = n
		h.last = n
		return
	}
	if n.value < h.first.value {
		n.right = h.first
		h.first.left = n
		h.first = n
		return
	}
	h.last.right = n
	n.left = h.last
	h.last = n
}
