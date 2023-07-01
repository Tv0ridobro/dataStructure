package graphalgo

import (
	"sort"
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/slices"
)

func TestFindArticulationPoints(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](10, false)
	g.AddEdgesDefault(0, 1, 2)
	points := FindArticulationPoints(g)
	if !slices.Equal(points, []int{0}) {
		t.Error("wrong answer")
	}

	g.AddEdgeDefault(2, 3)
	points = FindArticulationPoints(g)
	if !(slices.Equal(points, []int{2, 0}) || slices.Equal(points, []int{0, 2})) {
		t.Error("wrong answer")
	}
}

func TestFindArticulationPoints2(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](12, false)
	g.AddEdgesDefault(0, 1, 7, 2, 5, 3)
	g.AddEdgesDefault(1, 5, 9, 7, 6)
	g.AddEdgeDefault(5, 6)
	g.AddEdgesDefault(7, 8, 11)
	g.AddEdgesDefault(8, 9, 10)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(3, 4)

	points := FindArticulationPoints(g)
	sort.Ints(points)
	if !slices.Equal(points, []int{0, 3, 7, 8}) {
		t.Errorf("wrong asnwer")
	}
}

func TestFindArticulationPoints3(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](9, false)
	g.AddEdgesDefault(0, 1, 2)
	g.AddEdgesDefault(3, 4, 5)
	g.AddEdgesDefault(6, 7, 8)

	points := FindArticulationPoints(g)
	sort.Ints(points)
	if !slices.Equal(points, []int{0, 3, 6}) {
		t.Errorf("wrong asnwer %v", points)
	}
}

func TestFindArticulationPoints4(t *testing.T) {
	t.Parallel()

	g := graph.New[struct{}](10, false)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(1, 2)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(3, 4)
	g.AddEdgeDefault(4, 5)
	g.AddEdgeDefault(5, 6)
	g.AddEdgeDefault(6, 7)
	g.AddEdgeDefault(7, 8)
	g.AddEdgeDefault(8, 9)

	points := FindArticulationPoints(g)
	sort.Ints(points)
	if !slices.Equal(points, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("wrong asnwer %v", points)
	}
}
