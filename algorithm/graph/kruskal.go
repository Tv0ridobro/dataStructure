package graphalgo

import (
	"sort"

	"golang.org/x/exp/constraints"

	disjointset "github.com/Tv0ridobro/data-structure/disjoint-set"
	"github.com/Tv0ridobro/data-structure/graph"
)

// KruskalMST finds minimum spanning tree using Kruskal’s algorithm.
func KruskalMST[T constraints.Integer](graph *graph.Graph[T]) []Edge[T] {
	edges := make([]Edge[T], 0)

	for i, e := range graph.Edges {
		for _, ed := range e {
			edges = append(edges, Edge[T]{
				From:  i,
				To:    ed.To,
				Value: ed.Value,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Value < edges[j].Value
	})

	ds := disjointset.New(graph.Size())
	answer := make([]Edge[T], 0, graph.Size()-1)
	for _, e := range edges {
		if ds.Get(e.From) != ds.Get(e.To) {
			answer = append(answer, e)
			ds.Union(e.From, e.To)
		}
	}

	return answer
}
