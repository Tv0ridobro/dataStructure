package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/slices"
)

func TestKruskalMST(t *testing.T) {
	t.Parallel()
	g := graph.New[int](7, true)
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 7)
	g.AddEdge(1, 2, 9)
	g.AddEdge(1, 3, 6)
	g.AddEdge(1, 5, 15)
	g.AddEdge(2, 4, 8)
	g.AddEdge(2, 5, 7)
	g.AddEdge(4, 5, 5)
	g.AddEdge(5, 6, 9)
	g.AddEdge(5, 3, 8)
	g.AddEdge(3, 6, 11)
	if !slices.Equal(KruskalMST(g), []Edge[int]{{0, 1, 5}, {4, 5, 5}, {1, 3, 6}, {0, 2, 7}, {2, 5, 7}, {5, 6, 9}}) {
		t.Error("wrong asnwer")
	}
}
