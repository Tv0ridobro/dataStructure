package graph

type Edge[T any] struct {
	To    int
	Value T
}

type Graph[T any] struct {
	// for optimization this is public
	Edges      [][]Edge[T]
	isDirected bool
}

func New[T any](size int, isDirected bool) *Graph[T] {
	return &Graph[T]{
		Edges:      make([][]Edge[T], size),
		isDirected: isDirected,
	}
}

func (g *Graph[T]) AddEdge(from, to int, value T) {
	g.Edges[from] = append(g.Edges[from], Edge[T]{
		Value: value,
		To:    to,
	})
	if !g.isDirected && from != to {
		g.Edges[to] = append(g.Edges[to], Edge[T]{
			Value: value,
			To:    from,
		})
	}
}

func (g *Graph[T]) AddEdgeDefault(from, to int) {
	var zero T
	g.AddEdge(from, to, zero)
}

func (g *Graph[T]) AddEdgesDefault(from int, to ...int) {
	var zero T
	for _, e := range to {
		g.AddEdge(from, e, zero)
	}
}

func (g *Graph[T]) AddEdges(vertex int, edges []Edge[T]) {
	g.Edges[vertex] = append(g.Edges[vertex], edges...)
	if !g.isDirected {
		for _, e := range edges {
			if e.To == vertex {
				continue
			}
			g.Edges[e.To] = append(g.Edges[e.To], Edge[T]{
				Value: e.Value,
				To:    vertex,
			})
		}
	}
}

func (g *Graph[T]) IsDirected() bool {
	return g.isDirected
}

func (g *Graph[T]) Size() int {
	return len(g.Edges)
}
