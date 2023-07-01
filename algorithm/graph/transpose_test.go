package graphalgo

import (
	"reflect"
	"sort"
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestTranspose(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](5, true)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(1, 2)
	g.AddEdgeDefault(2, 3)
	g.AddEdgeDefault(3, 4)
	g.AddEdgeDefault(4, 0)

	ng := Transpose(g)
	if !reflect.DeepEqual(ng.Edges, [][]graph.Edge[struct{}]{
		{graph.Edge[struct{}]{To: 4}},
		{graph.Edge[struct{}]{To: 0}},
		{graph.Edge[struct{}]{To: 1}},
		{graph.Edge[struct{}]{To: 2}},
		{graph.Edge[struct{}]{To: 3}},
	}) {
		t.Errorf("wrong answer %v", ng.Edges)
	}
}

func TestTranspose2(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](2, true)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(1, 0)

	ng := Transpose(g)
	if !reflect.DeepEqual(ng.Edges, [][]graph.Edge[struct{}]{
		{graph.Edge[struct{}]{To: 1}},
		{graph.Edge[struct{}]{To: 0}},
	}) {
		t.Errorf("wrong answer %v", ng.Edges)
	}
}

func TestTranspose3(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](12, true)

	g.AddEdgesDefault(1, 2, 3, 7)
	g.AddEdgesDefault(2, 4)
	g.AddEdgesDefault(4, 3, 5)
	g.AddEdgesDefault(5, 4, 6, 9, 5, 5)
	g.AddEdgesDefault(7, 6, 8)
	g.AddEdgesDefault(8, 9)
	g.AddEdgesDefault(10, 8, 11, 11, 11, 11)
	g.AddEdgeDefault(11, 11)

	ng := Transpose(g)
	for i := range ng.Edges {
		sort.Slice(ng.Edges[i], func(u, v int) bool {
			return ng.Edges[i][u].To < ng.Edges[i][v].To
		})
	}

	arg := [][]graph.Edge[struct{}]{
		nil,
		nil,
		{graph.Edge[struct{}]{To: 1}},
		{graph.Edge[struct{}]{To: 1}, graph.Edge[struct{}]{To: 4}},
		{graph.Edge[struct{}]{To: 2}, graph.Edge[struct{}]{To: 5}},
		{graph.Edge[struct{}]{To: 4}, graph.Edge[struct{}]{To: 5}, graph.Edge[struct{}]{To: 5}},
		{graph.Edge[struct{}]{To: 5}, graph.Edge[struct{}]{To: 7}},
		{graph.Edge[struct{}]{To: 1}},
		{graph.Edge[struct{}]{To: 7}, graph.Edge[struct{}]{To: 10}},
		{graph.Edge[struct{}]{To: 5}, graph.Edge[struct{}]{To: 8}},
		nil,
		{
			graph.Edge[struct{}]{To: 10},
			graph.Edge[struct{}]{To: 10},
			graph.Edge[struct{}]{To: 10},
			graph.Edge[struct{}]{To: 10},
			graph.Edge[struct{}]{To: 11},
		},
	}

	if !reflect.DeepEqual(ng.Edges, arg) {
		t.Errorf("wrong answer \n%v\n%v", ng.Edges, arg)
	}
}

func TestTranspose4(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](12, false)

	g.AddEdgesDefault(1, 2, 3, 7)
	g.AddEdgesDefault(2, 4)
	g.AddEdgesDefault(4, 3, 5)
	g.AddEdgesDefault(5, 4, 6, 9, 5, 5)
	g.AddEdgesDefault(7, 6, 8)
	g.AddEdgesDefault(8, 9)
	g.AddEdgesDefault(10, 8, 11, 11, 11, 11)
	g.AddEdgeDefault(11, 11)

	ng := Transpose(g)
	if !reflect.DeepEqual(ng, g) {
		t.Error("now the same graph")
	}
}
