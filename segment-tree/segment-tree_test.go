package segmenttree

import (
	"math"
	"testing"

	mathx "github.com/Tv0ridobro/data-structure/math"
)

type Matrix struct {
	data [2][2]int
}

func TestSum(t *testing.T) {
	t.Parallel()
	tree := New[int]([]int{17, 2, 3, 4}, func(a, b int) int { return a + b }, 0)
	tests := []struct {
		l, r   int
		answer int
	}{
		{0, 2, 22},
		{0, 3, 26},
		{1, 2, 5},
		{1, 1, 2},
		{2, 3, 7},
	}
	for _, e := range tests {
		if val := tree.Query(e.l, e.r); val != e.answer {
			t.Errorf("%d != %d", val, e.answer)
		}
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	tree := New[int]([]int{17, 2, 3, 4, 5}, mathx.Max[int], math.MinInt64)
	tests := []struct {
		l, r   int
		answer int
	}{
		{0, 2, 17},
		{0, 3, 17},
		{1, 2, 3},
		{1, 1, 2},
		{2, 3, 4},
	}
	for _, e := range tests {
		if val := tree.Query(e.l, e.r); val != e.answer {
			t.Errorf("%d != %d", val, e.answer)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tree := New[int]([]int{17, 2, 3, 4, 5}, func(a, b int) int { return a * b }, 1)
	tests := []struct {
		l, r   int
		answer int
	}{
		{0, 2, 17 * 2 * 3},
		{0, 3, 17 * 2 * 3 * 4},
		{1, 2, 6},
		{1, 1, 2},
		{2, 3, 12},
	}
	for _, e := range tests {
		if val := tree.Query(e.l, e.r); val != e.answer {
			t.Errorf("%d != %d", val, e.answer)
		}
	}
}

func TestMatrixMul(t *testing.T) {
	t.Parallel()
	mul := func(a, b Matrix) Matrix {
		return Matrix{data: [2][2]int{
			{a.data[0][0]*b.data[0][0] + a.data[0][1]*b.data[1][0], a.data[0][0]*b.data[0][1] + a.data[0][1]*b.data[1][1]},
			{a.data[1][0]*b.data[0][0] + a.data[1][1]*b.data[1][0], a.data[1][0]*b.data[0][1] + a.data[1][1]*b.data[1][1]},
		}}
	}
	neutral := Matrix{[2][2]int{{1, 0}, {0, 1}}}
	tree := New[Matrix]([]Matrix{
		{[2][2]int{{1, 0}, {0, 1}}},
		{[2][2]int{{1, 0}, {0, 1}}},
		{[2][2]int{{1, 0}, {0, 1}}},
		{[2][2]int{{1, 0}, {0, 1}}},
	}, mul, neutral)
	if val := tree.Query(0, 3); val != neutral {
		t.Errorf("val != neutral")
	}
}

func TestSegmentTree_Modify(t *testing.T) {
	t.Parallel()
	tree := New[int]([]int{8, 0, 10, 500}, mathx.Max[int], math.MinInt)
	if val := tree.Query(0, 3); val != 500 {
		t.Errorf("val != neutral %d", val)
	}
	tree.Modify(3, -1)
	if val := tree.Query(3, 3); val != -1 {
		t.Errorf("val != neutral %d", val)
	}
	if val := tree.Query(2, 3); val != 10 {
		t.Errorf("val != neutral %d", val)
	}
	if val := tree.Query(0, 1); val != 8 {
		t.Errorf("val != neutral %d", val)
	}
	tree.Modify(0, -2)
	if val := tree.Query(0, 0); val != -2 {
		t.Errorf("val != neutral %d", val)
	}
	if val := tree.Query(0, 1); val != 0 {
		t.Errorf("val != neutral %d", val)
	}
}
