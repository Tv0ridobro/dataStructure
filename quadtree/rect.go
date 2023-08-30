package quadtree

import mathx "github.com/Tv0ridobro/data-structure/math"

// Vec2 represents point.
type Vec2[T mathx.Numeric] struct {
	X, Y T
}

// Rect is axis-aligned rectangle.
type Rect[T mathx.Numeric] struct {
	TopLeft  Vec2[T]
	BotRight Vec2[T]
}

// Overlap returns true if given rectangles overlap, false otherwise.
func (r Rect[T]) Overlap(o Rect[T]) bool {
	if r.Contains(o.TopLeft) || r.Contains(o.BotRight) ||
		r.Contains(Vec2[T]{o.BotRight.X, o.TopLeft.Y}) || r.Contains(Vec2[T]{o.TopLeft.X, o.BotRight.Y}) {
		return true
	}

	if o.Contains(r.TopLeft) || o.Contains(r.BotRight) ||
		o.Contains(Vec2[T]{r.BotRight.X, r.TopLeft.Y}) || o.Contains(Vec2[T]{r.TopLeft.X, r.BotRight.Y}) {
		return true
	}

	return false
}

// Contains returns true if rectangle contains given point, false otherwise.
func (r Rect[T]) Contains(v Vec2[T]) bool {
	x := v.X
	y := v.Y

	return r.TopLeft.X <= x && x <= r.BotRight.X &&
		r.BotRight.Y <= y && y <= r.TopLeft.Y
}
