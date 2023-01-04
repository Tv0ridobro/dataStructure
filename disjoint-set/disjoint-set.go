// Package disjointset implements a disjoint set.
// See https://en.wikipedia.org/wiki/Disjoint-set_data_structure for more details.
package disjointset

// DisjointSet represents a disjoint set.
// Zero value of DisjointSet is disjoint set of 0 elements.
type DisjointSet struct {
	ranks   []int
	parents []int
	sizes   []int
}

// New returns an initialized disjoint set of given size.
func New(size int) *DisjointSet {
	d := &DisjointSet{
		ranks:   make([]int, size),
		parents: make([]int, size),
		sizes:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		d.parents[i] = i
		d.sizes[i] = 1
	}
	return d
}

// Add adds new element to DisjointSet.
func (d *DisjointSet) Add() {
	d.ranks = append(d.ranks, 0)
	d.sizes = append(d.sizes, 1)
	d.parents = append(d.parents, len(d.parents))
}

// Get returns root element of set containing given element.
func (d *DisjointSet) Get(x int) int {
	if x < 0 || x >= len(d.parents) {
		return -1
	}
	if d.parents[x] != x {
		d.parents[x] = d.Get(d.parents[x])
	}
	return d.parents[x]
}

// Size returns size of set containing given element.
func (d *DisjointSet) Size(x int) int {
	if x < 0 || x >= len(d.parents) {
		return 0
	}
	return d.sizes[d.Get(x)]
}

// Union unions two sets containing given elements.
func (d *DisjointSet) Union(x int, y int) {
	x = d.Get(x)
	y = d.Get(y)
	if x == y {
		return
	}
	if d.ranks[x] == d.ranks[y] {
		d.ranks[x]++
	}
	if d.ranks[x] > d.ranks[y] {
		d.parents[y] = x
		d.sizes[x] += d.sizes[y]
	} else {
		d.parents[x] = y
		d.sizes[y] += d.sizes[x]
	}
}
