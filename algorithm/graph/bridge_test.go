package graphalgo

import (
	"testing"

	"golang.org/x/exp/slices"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestFindBridges(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](12, false)
	g.AddEdgesDefault(0, 1, 7, 2, 5, 3)
	g.AddEdgesDefault(1, 5, 9, 7, 6)
	g.AddEdgeDefault(5, 6)
	g.AddEdgesDefault(7, 8, 11)
	g.AddEdgesDefault(8, 9, 10)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(3, 4)
	bridges := FindBridges(g)

	if len(bridges) != 3 || !slices.Contains(bridges, Edge[struct{}]{
		From:  3,
		To:    4,
		Value: struct{}{},
	}) || !slices.Contains(bridges, Edge[struct{}]{
		From:  7,
		To:    11,
		Value: struct{}{},
	}) || !slices.Contains(bridges, Edge[struct{}]{
		From:  8,
		To:    10,
		Value: struct{}{},
	}) {
		t.Errorf("wrong answer")
	}
}

func TestFindBridges2(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](4, false)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(2, 3)

	bridges := FindBridges(g)

	if len(bridges) != 2 || !slices.Contains(bridges, Edge[struct{}]{
		From:  0,
		To:    1,
		Value: struct{}{},
	}) || !slices.Contains(bridges, Edge[struct{}]{
		From:  2,
		To:    3,
		Value: struct{}{},
	}) {
		t.Errorf("wrong answer")
	}
}
