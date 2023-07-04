package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	mathx "github.com/Tv0ridobro/data-structure/math"
	sparsetable "github.com/Tv0ridobro/data-structure/sparse-table"
	"github.com/Tv0ridobro/data-structure/stack"
)

type LCA struct {
	m   []int
	o   []int
	pos []int
	st  *sparsetable.SparseTable[int]
}

// NewLCA returns new LCA struct for computing LCA of given vertexes.
func NewLCA[T any](g *graph.Graph[T]) LCA {
	m := make([]int, g.Size())
	o := make([]int, g.Size())
	pos := make([]int, g.Size())
	counter := new(int)
	s := stack.New[int]()
	dfsLCA(0, g, m, o, pos, counter, s)
	st := sparsetable.New(mathx.Min[int], s.All())
	return LCA{
		m:   m,
		o:   o,
		pos: pos,
		st:  st,
	}
}

// LCA returns LCA of given vertexes.
func (a LCA) LCA(u, v int) int {
	first, second := a.pos[u], a.pos[v]
	return a.o[a.st.Query(mathx.Min(first, second), mathx.Max(first, second))]
}

func dfsLCA[T any](vertex int, g *graph.Graph[T], m, o, pos []int, counter *int, order *stack.Stack[int]) {
	a := *counter
	pos[vertex] = order.Size()
	order.Push(a)
	o[a] = vertex
	m[vertex] = *counter
	*counter = a + 1
	for _, e := range g.Edges[vertex] {
		dfsLCA(e.To, g, m, o, pos, counter, order)
		order.Push(a)
	}
}
