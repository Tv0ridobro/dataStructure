package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/slices"
)

func TestShortedPath(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](6, true)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(0, 2)
	g.AddEdgeDefault(1, 2)
	g.AddEdgeDefault(1, 3)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(3, 4)
	g.AddEdgeDefault(4, 1)
	g.AddEdgeDefault(4, 0)
	g.AddEdgeDefault(4, 5)

	if !(slices.Equal(FindShortestPath(g, 0, 5), []int{0, 1, 3, 4, 5}) ||
		slices.Equal(FindShortestPath(g, 0, 5), []int{0, 2, 3, 4, 5})) {
		t.Errorf("wrong answer")
	}
}
