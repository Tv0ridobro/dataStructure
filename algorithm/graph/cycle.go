package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/stack"
)

// FindCycle returns cycle if there is any
// otherwise empty slice returned.
func FindCycle[T any](g *graph.Graph[T]) []int {
	visited := make([]byte, g.Size())
	for i := 0; i < g.Size(); i++ {
		s := stack.New[int]()
		s.Push(i)
		if visited[i] == 0 && dfsCycle(i, -1, visited, g, s) {
			last := s.Peek()
			order := s.All()
			for j := len(order) - 2; j >= 0; j-- {
				if order[j] == last {
					return order[j:]
				}
			}
		}
	}
	return nil
}

// dfsCycle helper function to find cycle using dfs.
func dfsCycle[T any](vertex, from int, visited []byte, g *graph.Graph[T], order *stack.Stack[int]) bool {
	visited[vertex] = 1
	for _, e := range g.Edges[vertex] {
		if !g.IsDirected() && e.To == from {
			continue
		}
		order.Push(e.To)
		if visited[e.To] == 1 {
			return true
		}
		if visited[e.To] == 0 && dfsCycle(e.To, vertex, visited, g, order) {
			return true
		}
	}
	visited[vertex] = 2
	order.Pop()
	return false
}
