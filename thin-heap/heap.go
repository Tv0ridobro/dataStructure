package thin_heap

import (
	"golang.org/x/exp/constraints"
)

type ThinHeap[T constraints.Ordered] struct {
	first *Node[T]
	last  *Node[T]
}

func (h *ThinHeap[T]) Min() T {
	if h.first != nil {
		return h.first.value
	}
	var zero T
	return zero
}

func (h *ThinHeap[T]) DeleteMin() T {
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

	max := -1
	x = h.first
	rangs := make(map[int]*Node[T])
	for x != nil {
		next := x.right
		for node, ok := rangs[x.rank]; ok; node, ok = rangs[x.rank] {
			if node.value < x.value {
				node, x = x, node
			}
			node.right = x.child
			if x.child != nil {
				x.child.left = node
			}
			node.left = x
			x.child = node

			delete(rangs, x.rank)
			x.rank++
		}
		rangs[x.rank] = x
		if x.rank > max {
			max = x.rank
		}
		x = next
	}

	n := New[T]()
	for i := 0; i <= max; i++ {
		value := rangs[i]
		if value != nil {
			value.left = nil
			value.right = nil
			n.insert(value)
		}
	}
	h.first, h.last = n.first, n.last

	return tmp.value
}

func (h *ThinHeap[T]) Insert(element T) {
	n := &Node[T]{value: element}
	h.insert(n)
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

func New[T constraints.Ordered]() *ThinHeap[T] {
	return &ThinHeap[T]{
		first: nil,
		last:  nil,
	}
}
