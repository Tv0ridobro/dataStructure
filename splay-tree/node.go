package splaytree

// node represents node of a splay tree.
type node[T any] struct {
	parent *node[T]
	right  *node[T]
	left   *node[T]
	value  T
	size   int
}

// isRight returns true if node is right child of his parent, false otherwise.
func (n *node[T]) isRight() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.right == n
}

// isRight returns true if node is left child of his parent, false otherwise.
func (n *node[T]) isLeft() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.left == n
}

// isRoot returns true if given node is root.
func (n *node[T]) isRoot() bool {
	return n.parent == nil
}

// setRight sets c as right child of p.
func setRight[T any](p, c *node[T]) {
	if p == nil {
		return
	}
	p.right = c
	if c != nil {
		c.parent = p
	}
}

// setRight sets c as left child of p.
func setLeft[T any](p, c *node[T]) {
	if p == nil {
		return
	}
	p.left = c
	if c != nil {
		c.parent = p
	}
}

// find returns node containing given value
// or last node reached.
func (n *node[T]) find(value T, comp func(T, T) int) *node[T] {
	if n == nil {
		return nil
	}
	if comp(n.value, value) == 0 {
		n.splay()
		return n
	}
	if comp(value, n.value) < 0 {
		if n.left == nil {
			n.splay()
			return n
		}
		return n.left.find(value, comp)
	}
	if n.right == nil {
		n.splay()
		return n
	}
	return n.right.find(value, comp)
}

// splay uses splay.
func (n *node[T]) splay() {
	for !n.isRoot() && !n.parent.isRoot() {
		if (n.isLeft() && n.parent.isLeft()) || (n.isRight() && n.parent.isRight()) {
			n.zigZig()
		} else {
			n.zigZag()
		}
	}
	if !n.isRoot() {
		n.zig()
	}
}

// zig uses left or right rotation.
func (n *node[T]) zig() {
	grandParent := n.parent.parent
	isLeft := n.parent.isLeft()
	if n.isLeft() {
		right := n.right
		setRight(n, n.parent)
		setLeft(n.right, right)
		n.right.recalculateSize()
	} else {
		left := n.left
		setLeft(n, n.parent)
		setRight(n.left, left)
		n.left.recalculateSize()
	}
	n.recalculateSize()
	n.parent = grandParent
	if isLeft {
		setLeft(grandParent, n)
	} else {
		setRight(grandParent, n)
	}
}

// zigZig uses zig zig rotation.
func (n *node[T]) zigZig() {
	n.parent.zig()
	n.zig()
}

// zigZag uses zig zag rotation.
func (n *node[T]) zigZag() {
	n.zig()
	n.zig()
}

// max returns node with max element.
func (n *node[T]) max() *node[T] {
	for n.right != nil {
		n = n.right
	}
	return n
}

// min returns node with min element.
func (n *node[T]) min() *node[T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

// split splits given node by given key into two nodes.
func split[T any](n *node[T], key T, comp func(T, T) int) (*node[T], *node[T]) {
	if n == nil {
		return nil, nil
	}
	nn := n.find(key, comp)
	if comp(key, n.value) <= 0 {
		left := nn.left
		if left != nil {
			left.parent = nil
		}
		nn.left = nil
		nn.recalculateSize()
		return left, n
	}
	right := nn.right
	if right != nil {
		right.parent = nil
	}
	nn.right = nil
	nn.recalculateSize()
	return nn, right
}

// merge merges two nodes, all elements of left node should be less than any of right node.
func merge[T any](left, right *node[T]) *node[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	max := left.max()
	max.splay()
	setRight(max, right)
	max.recalculateSize()
	return max
}

// recalculateSize recalculates size of given node.
func (n *node[T]) recalculateSize() {
	if n == nil {
		return
	}
	n.size = 0
	if n.left != nil {
		n.size += n.left.size
	}
	if n.right != nil {
		n.size += n.right.size
	}
	n.size++
}

// kth returns kth greatest element.
func (n *node[T]) kth(i int) *node[T] {
	if n.size-1 == i {
		return n
	}
	if n.size > i {
		return n.left.kth(i)
	}
	return n.right.kth(i - n.left.size - 1)
}

// next returns node A in a way that this.value < A.value
// and there are no node B that this.value < B.value < A.value.
func (n *node[T]) next() *node[T] {
	if n.right != nil {
		return n.right.min()
	}
	for n.isRight() {
		n = n.parent
	}
	n = n.parent
	return n
}

// getAll returns all elements in node.
// len of elements should be same as size of node.
func (n *node[T]) getAll(elements []T) {
	lSize := 0
	if n.left != nil {
		lSize = n.left.size
		n.left.getAll(elements[:lSize])
	}
	elements[lSize] = n.value
	if n.right != nil {
		n.right.getAll(elements[lSize+1:])
	}
}
