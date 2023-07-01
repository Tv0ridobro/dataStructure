package graphalgo

import "github.com/Tv0ridobro/data-structure/graph"

// Transpose return graph that is transposes of given graph.
// Returns given graph if it is not directed.
func Transpose[T any](g *graph.Graph[T]) *graph.Graph[T] {
	if !g.IsDirected() {
		return g
	}
	ng := graph.New[T](g.Size(), true)

	for i, edge := range g.Edges {
		for _, e := range edge {
			ng.AddEdge(e.To, i, e.Value)
		}
	}

	return ng
}
