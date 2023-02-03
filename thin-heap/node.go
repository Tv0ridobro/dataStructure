package thinheap

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	value T
	rank  int
	child *Node[T]
	right *Node[T]
	left  *Node[T]
}

func (n Node[T]) isThin() bool {
	if n.rank == 1 {
		return n.child == nil
	}
	if n.child == nil {
		return false
	}
	return n.child.rank+1 != n.rank
}
