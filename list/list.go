// Package list implements a doubly linked list.
// See https://en.wikipedia.org/wiki/Linked_list for more details.
package list

// List represents a doubly linked list.
// Zero value of List is empty list.
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// New returns an initialized list.
func New[T any]() *List[T] {
	return &List[T]{}
}

// PushFront adds value to the beginning of the list.
func (l *List[T]) PushFront(value T) {
	l.len++
	nn := &Node[T]{
		Value: value,
	}
	if l.head == nil {
		l.head = nn
		return
	}
	if l.tail == nil { // 1 element in list
		l.tail = l.head
		l.head = nn
		l.tail.prev = l.head
		l.head.next = l.tail
		return
	}
	nn.next = l.head
	l.head.prev = nn
	l.head = nn
}

// PushBack adds value to the end of the list.
func (l *List[T]) PushBack(value T) {
	l.len++
	nn := &Node[T]{
		Value: value,
	}
	if l.head == nil {
		l.head = nn
		return
	}
	if l.tail == nil { // 1 element in list
		l.tail = nn
		nn.prev = l.head
		l.head.next = l.tail
		return
	}
	l.tail.next = nn
	nn.prev = l.tail
	l.tail = nn
}

// Len returns length of the list.
func (l *List[T]) Len() int {
	return l.len
}

// All returns all elements from the list.
func (l *List[T]) All() []T {
	data := make([]T, 0, l.len)
	it := l.head
	for it != nil {
		data = append(data, it.Value)
		it = it.next
	}
	return data
}

// PopBack removes last element of the list.
// Returns zero value if list is empty.
func (l *List[T]) PopBack() T {
	if l.len == 0 {
		var empty T
		return empty
	}
	l.len--
	if l.tail == nil {
		v := l.head.Value
		l.head = nil
		return v
	}
	v := l.tail.Value
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	if l.tail == l.head { // if 1 element in list
		l.tail = nil
	}
	return v
}

// PopFront removes first element of the list.
// Returns zero value if list is empty.
func (l *List[T]) PopFront() T {
	if l.len == 0 {
		var empty T
		return empty
	}
	l.len--
	v := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		return v
	}
	l.head.prev = nil
	if l.head == l.tail {
		l.tail = nil
	}
	return v
}

// ChangeAt changes value at given index.
// Returns zero value if list is empty.
func (l *List[T]) ChangeAt(i int, value T) {
	l.Node(i).Value = value
}

// Peek returns element at the given index.
// Returns zero value if index is less or equal to len of the list.
func (l *List[T]) Peek(i int) T {
	v := l.Node(i)
	if v == nil {
		var a T
		return a
	}
	return v.Value
}

// Node returns Node at the given index.
// Returns zero value if index is less or equal to len of the list.
func (l *List[T]) Node(i int) *Node[T] {
	if i < 0 || i >= l.Len() {
		return nil
	}
	if i <= l.len/2 {
		it := l.head
		counter := 0
		for counter != i {
			it = it.next
			counter++
		}
		return it
	}
	it := l.tail
	counter := 0
	for counter != l.len-1-i {
		it = it.prev
		counter++
	}
	return it
}

// Merge merges two lists.
func (l *List[T]) Merge(other *List[T]) {
	if other.head == nil {
		return
	}
	if l.head == nil {
		l.len = other.len
		l.head = other.head
		l.tail = other.tail
		return
	}
	l.len += other.len
	if l.tail == nil {
		l.head.next = other.head
		other.head.prev = l.head
	} else {
		l.tail.next = other.head
		other.head.prev = l.tail
	}
	if other.tail == nil {
		l.tail = other.head
	} else {
		l.tail = other.tail
	}
}

// Clear clears list.
func (l *List[T]) Clear() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

// Reverse reverses list.
func (l *List[T]) Reverse() {
	if l.head == nil || l.tail == nil {
		return
	}
	var prev *Node[T]
	it, f := l.head, l.head
	next := it.next
	for next != nil {
		it.next = prev
		it.prev = next
		prev, it, next = it, next, next.next
	}
	it.next = prev
	l.head = it
	l.tail = f
}

// InsertAfter adds new Node with given value after given Node.
func (l *List[T]) InsertAfter(node *Node[T], value T) {
	if node == l.tail || l.len == 1 {
		l.PushBack(value)
		return
	}
	l.len++
	nn := &Node[T]{
		Value: value,
	}
	next := node.next
	node.next = nn
	nn.next = next
	nn.prev = node
	next.prev = nn
}

// Remove removes given node from list.
func (l *List[T]) Remove(n *Node[T]) {
	if n == l.head {
		l.PopFront()
		return
	}
	if n == l.tail {
		l.PopBack()
		return
	}
	l.len--
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

// InsertBefore adds new Node with given value before given Node.
func (l *List[T]) InsertBefore(n *Node[T], value T) {
	if n == l.head {
		l.PushFront(value)
		return
	}
	l.len--
	nn := &Node[T]{Value: value}
	prev := n.prev
	prev.next = nn
	nn.prev = prev
	nn.next = n
	n.prev = nn
}

// Cut cuts list into two list.
// i is last element in first list, for example
// cut(1) for list {1,2,3,4,5} returns {1,2}, {3,4,5}.
func (l *List[T]) Cut(i int) (*List[T], *List[T]) {
	n := l.Node(i)
	if n == nil {
		return l, New[T]()
	}
	other := New[T]()
	other.len = l.len - i - 1
	l.len = i + 1
	other.tail = l.tail
	other.head = n.next
	if other.len == 1 {
		other.tail = nil
	}
	if other.head != nil {
		other.head.prev = nil
	}
	l.tail = n
	l.tail.next = nil
	if l.len == 1 {
		l.tail = nil
	}
	return l, other
}

// Back returns last element.
func (l *List[T]) Back() T {
	if l.head == nil {
		var a T
		return a
	}
	if l.tail != nil {
		return l.tail.Value
	}
	return l.head.Value
}

// Front returns first element.
func (l *List[T]) Front() T {
	if l.head == nil {
		var a T
		return a
	}
	return l.head.Value
}
