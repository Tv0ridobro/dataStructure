package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/slices"
)

func TestFindCycle(t *testing.T) {
	t.Parallel()
	g := graph.New[int](4, false)
	g.AddEdge(0, 1, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(0, 2, 0)
	if !slices.Equal(FindCycle(g), []int{0, 1, 3, 2, 0}) {
		t.Error("wrong answer")
	}
}
