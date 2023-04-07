package graphalgo

import (
	"testing"

	"golang.org/x/exp/slices"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestTopologicalSort(t *testing.T) {
	t.Parallel()
	g := graph.New[int](7, true)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(0, 2)
	g.AddEdgeDefault(1, 5)
	g.AddEdgeDefault(1, 2)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(5, 3)
	g.AddEdgeDefault(5, 4)
	g.AddEdgeDefault(6, 1)
	g.AddEdgeDefault(6, 5)

	answer := TopologicalSort(g)

	if !checkCorrectness(g, answer) {
		t.Error("incorrect topological order")
	}
}

func checkCorrectness[T comparable](g *graph.Graph[T], order []int) bool {
	for i, e := range g.Edges {
		for _, edge := range e {
			first := slices.Index(order, i)
			second := slices.Index(order, edge.To)

			if first > second {
				return false
			}
		}
	}

	return true
}
