package quadtree

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestQuadTree(t *testing.T) {
	t.Parallel()
	qt := New(Rect[int]{
		TopLeft:  Vec2[int]{-10, 10},
		BotRight: Vec2[int]{10, -10},
	}, 1)

	if !qt.Insert(Vec2[int]{0, 0}) {
		t.Fail()
	}
	if !qt.Contains(Vec2[int]{0, 0}) {
		t.Fail()
	}
	if qt.Contains(Vec2[int]{1, 0}) {
		t.Fail()
	}
	if !qt.Insert(Vec2[int]{0, 0}) {
		t.Fail()
	}
	qt.Delete(Vec2[int]{0, 0})
	if !qt.Contains(Vec2[int]{0, 0}) {
		t.Fail()
	}
	qt.Delete(Vec2[int]{0, 0})
	if qt.Contains(Vec2[int]{0, 0}) {
		t.Fail()
	}
}

func TestQuadTree_Points(t *testing.T) {
	t.Parallel()
	qt := New(Rect[int]{
		TopLeft:  Vec2[int]{-20, 20},
		BotRight: Vec2[int]{20, -20},
	}, 2)

	points := []Vec2[int]{
		//nolint:gofumpt // Too much points
		{-7, 10}, {3, 19}, {5, 6}, {5, -6}, {3, 2}, {7, 15}, {-10, -10},
		{-1, 0}, {4, 13}, {5, -4}, {9, 0}, {8, 9}, {-6, -6}, {-2, -1},
		{-19, 20}, {0, 0}, {-17, -11}, {7, 16}, {8, -18}, {0, 10}, {19, 0},
		{7, 16}, {7, 16}, {7, 16}, {5, -6}, {5, -6},
	}

	for i := range points {
		qt.Insert(points[i])
	}

	rectangles := []Rect[int]{
		{Vec2[int]{-20, 20}, Vec2[int]{20, -20}},
		{Vec2[int]{-7, 10}, Vec2[int]{17, 2}},
		{Vec2[int]{-7, 10}, Vec2[int]{17, 2}},
		{Vec2[int]{-20, 20}, Vec2[int]{0, 0}},
		{Vec2[int]{-6, 14}, Vec2[int]{10, -5}},
		{Vec2[int]{9, -1}, Vec2[int]{13, -14}},
		{Vec2[int]{100, 100}, Vec2[int]{100, -100}},
		{Vec2[int]{-50, 50}, Vec2[int]{-30, 30}},
	}

	for i := range rectangles {
		p := qt.Points(rectangles[i])
		for j := range p {
			if !(slices.Count(points, p[j]) == slices.Count(p, p[j])) {
				t.Errorf("dublicate points")
			}
		}

		for j := range points {
			if rectangles[i].Contains(points[j]) && !slices.Contains(p, points[j]) {
				t.Errorf("point %v missing", points[j])
			}
		}
	}
}

func TestQuadTree_Boundary(t *testing.T) {
	t.Parallel()
	r := Rect[int]{
		TopLeft:  Vec2[int]{-20, 20},
		BotRight: Vec2[int]{20, -20},
	}
	qt := New(r, 4)

	if qt.Boundary() != r {
		t.Errorf("wrong answer")
	}
}
