package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/list"
	"github.com/Tv0ridobro/data-structure/math"
)

// FindBridges returns all bridges of given graph.
func FindBridges[T any](g *graph.Graph[T]) []Edge[T] {
	enter := make([]int, g.Size())
	ret := make([]int, g.Size())
	ans := list.New[Edge[T]]()

	for i := 0; i < g.Size(); i++ {
		if enter[i] == 0 {
			dfsBridge(g, i, 0, new(int), enter, ret, ans)
		}
	}

	return ans.All()
}

func dfsBridge[T any](g *graph.Graph[T], vertex, from int, time *int, enter, ret []int, answer *list.List[Edge[T]]) {
	*time++
	enter[vertex] = *time
	ret[vertex] = *time
	for _, e := range g.Edges[vertex] {
		if e.To == from {
			continue
		}
		if enter[e.To] != 0 {
			ret[vertex] = math.Min(enter[e.To], ret[vertex])
		} else {
			dfsBridge(g, e.To, vertex, time, enter, ret, answer)
			ret[vertex] = math.Min(ret[e.To], ret[vertex])

			if enter[vertex] < ret[e.To] {
				answer.PushBack(Edge[T]{
					From:  vertex,
					To:    e.To,
					Value: e.Value,
				})
			}
		}
	}
}
