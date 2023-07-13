// Package graph provides a generic implementation of a Graph data structure.
package graph

// Edge represents an edge in a Graph.
type Edge[T any] struct {
	To    int
	Value T
}

// Graph is a generic graph structure where each edge holds a value of type T.
type Graph[T any] struct {
	// for optimization this is public
	Edges      [][]Edge[T]
	isDirected bool
}

// New creates and returns a new instance of Graph with the specified size.
func New[T any](size int, isDirected bool) *Graph[T] {
	return &Graph[T]{
		Edges:      make([][]Edge[T], size),
		isDirected: isDirected,
	}
}

// AddEdge adds a new edge to the graph.
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

// AddEdgeDefault adds a new edge to the graph with a default value of T.
func (g *Graph[T]) AddEdgeDefault(from, to int) {
	var zero T
	g.AddEdge(from, to, zero)
}

// AddEdgesDefault adds new edges to the graph with a default value of T.
func (g *Graph[T]) AddEdgesDefault(from int, to ...int) {
	var zero T
	for _, e := range to {
		g.AddEdge(from, e, zero)
	}
}

// AddEdges adds a slice of edges to the graph at the specified vertex int.
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

// IsDirected checks and returns whether the graph is directed or not.
func (g *Graph[T]) IsDirected() bool {
	return g.isDirected
}

// Size returns the size of the graph which is the number of edges it has.
func (g *Graph[T]) Size() int {
	return len(g.Edges)
}
