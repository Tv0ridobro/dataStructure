package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/list"
)

// TopologicalSort returns vertices in topological order.
func TopologicalSort[T any](g *graph.Graph[T]) []int {
	if !g.IsDirected() {
		ans := make([]int, g.Size())
		for i := 0; i < g.Size(); i++ {
			ans[i] = i
		}
		return ans
	}
	l := list.New[int]()
	visited := make([]byte, g.Size())
	for i := 0; i < g.Size(); i++ {
		if visited[i] == 0 {
			dfsTopSort(i, visited, g, l)
		}
	}
	return l.All()
}

// dfsTopSort helper function to find topological sort using dfs.
func dfsTopSort[T any](vertex int, visited []byte, g *graph.Graph[T], l *list.List[int]) {
	visited[vertex] = 1
	for _, e := range g.Edges[vertex] {
		if visited[e.To] == 0 {
			dfsTopSort(e.To, visited, g, l)
		}
	}
	visited[vertex] = 2
	l.PushFront(vertex)
}
