package graphalgo

// Edge struct represents edge of graph.
type Edge[T any] struct {
	From, To int
	Value    T
}
