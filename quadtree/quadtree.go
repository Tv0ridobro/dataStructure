// Package quadtree implements a quadtree.
// See https://en.wikipedia.org/wiki/Quadtree for more details.
package quadtree

import (
	mathx "github.com/Tv0ridobro/data-structure/math"
)

// QuadTree represents a QuadTree.
// Zero value of QuadTree is invalid segment tree, should be used only with New().
type QuadTree[T mathx.Numeric] struct {
	boundary Rect[T]

	points []Vec2[T]

	northWest *QuadTree[T]
	northEast *QuadTree[T]
	southWest *QuadTree[T]
	southEast *QuadTree[T]
}

// New returns new QuadTree.
func New[T mathx.Numeric](boundary Rect[T], capacity int) *QuadTree[T] {
	return &QuadTree[T]{
		points:    make([]Vec2[T], 0, capacity),
		boundary:  boundary,
		northWest: nil,
		northEast: nil,
		southWest: nil,
		southEast: nil,
	}
}

// Insert inserts point into quadtree.
func (qt *QuadTree[T]) Insert(point Vec2[T]) bool {
	if !qt.boundary.Contains(point) {
		return false
	}

	if len(qt.points) < cap(qt.points) {
		qt.points = append(qt.points, point)
		return true
	}

	if qt.northWest == nil {
		qt.subdivide()
	}

	switch {
	case qt.northWest.Insert(point):
	case qt.northEast.Insert(point):
	case qt.southWest.Insert(point):
	case qt.southEast.Insert(point):
	default:
		return false
	}
	return true
}

// Delete deletes point from quadtree.
func (qt *QuadTree[T]) Delete(point Vec2[T]) bool {
	if !qt.boundary.Contains(point) {
		return false
	}

	for i := range qt.points {
		if qt.points[i] == point {
			qt.points = append(qt.points[:i], qt.points[i+1:]...)
			return true
		}
	}

	if qt.northWest == nil {
		return false
	}

	switch {
	case qt.northWest.Delete(point):
	case qt.northEast.Delete(point):
	case qt.southWest.Delete(point):
	case qt.southEast.Delete(point):
	default:
		return false
	}
	return true
}

// Contains returns true if quadtree contains given element, false otherwise.
func (qt *QuadTree[T]) Contains(point Vec2[T]) bool {
	if !qt.boundary.Contains(point) {
		return false
	}

	for i := range qt.points {
		if qt.points[i] == point {
			return true
		}
	}

	if qt.northWest == nil {
		return false
	}

	switch {
	case qt.northWest.Contains(point):
	case qt.northEast.Contains(point):
	case qt.southWest.Contains(point):
	case qt.southEast.Contains(point):
	default:
		return false
	}
	return true
}

// Boundary returns boundary of quadtree.
func (qt *QuadTree[T]) Boundary() Rect[T] {
	return qt.boundary
}

// Points returns all points inside given rectangle.
func (qt *QuadTree[T]) Points(r Rect[T]) []Vec2[T] {
	if !qt.boundary.Overlap(r) {
		return nil
	}

	points := make([]Vec2[T], 0)
	for i := range qt.points {
		if r.Contains(qt.points[i]) {
			points = append(points, qt.points[i])
		}
	}

	if qt.northWest == nil {
		return points
	}

	points = append(points, qt.northWest.Points(r)...)
	points = append(points, qt.northEast.Points(r)...)
	points = append(points, qt.southWest.Points(r)...)
	points = append(points, qt.southEast.Points(r)...)

	return points
}

func (qt *QuadTree[T]) subdivide() {
	x1 := qt.boundary.TopLeft.X
	y1 := qt.boundary.TopLeft.Y

	x3 := qt.boundary.BotRight.X
	y3 := qt.boundary.BotRight.Y

	x2 := (x1 + x3) / 2
	y2 := (y1 + y3) / 2

	qt.northWest = New(Rect[T]{
		TopLeft:  Vec2[T]{x1, y1},
		BotRight: Vec2[T]{x2, y2},
	}, cap(qt.points))
	qt.northEast = New(Rect[T]{
		TopLeft:  Vec2[T]{x2, y1},
		BotRight: Vec2[T]{x3, y2},
	}, cap(qt.points))
	qt.southWest = New(Rect[T]{
		TopLeft:  Vec2[T]{x1, y2},
		BotRight: Vec2[T]{x2, y3},
	}, cap(qt.points))
	qt.southEast = New(Rect[T]{
		TopLeft:  Vec2[T]{x2, y2},
		BotRight: Vec2[T]{x3, y3},
	}, cap(qt.points))
}
