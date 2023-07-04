package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/list"
	"github.com/Tv0ridobro/data-structure/math"
)

// FindArticulationPoints finds all articulation points in given graph.
func FindArticulationPoints[T any](g *graph.Graph[T]) []int {
	enter := make([]int, g.Size())
	ret := make([]int, g.Size())
	ans := list.New[int]()

	for i := 0; i < g.Size(); i++ {
		if enter[i] == 0 {
			dfsArticulationPoints(g, i, i, new(int), enter, ret, ans)
		}
	}

	return ans.All()
}

func dfsArticulationPoints[T any](g *graph.Graph[T], v, from int, time *int, enter, ret []int, answer *list.List[int]) {
	*time++
	enter[v] = *time
	ret[v] = *time

	counter := 0
	for _, e := range g.Edges[v] {
		if e.To == from {
			continue
		}
		if enter[e.To] != 0 {
			ret[v] = math.Min(enter[e.To], ret[v])
		} else {
			counter++
			dfsArticulationPoints(g, e.To, v, time, enter, ret, answer)
			ret[v] = math.Min(ret[e.To], ret[v])

			if enter[v] <= ret[e.To] && v != from {
				answer.PushBack(v)
			}
		}
	}

	if v == from && counter > 1 {
		answer.PushBack(v)
	}
}
