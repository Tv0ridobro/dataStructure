package thinheap

import (
	"math"
	"math/rand"
	"testing"
)

func TestRightOrder(t *testing.T) {
	t.Parallel()
	heap := New[int]()
	for i := 0; i < 1000; i++ {
		heap.Insert(i)
	}

	for i := 0; i < 1000; i++ {
		if heap.DeleteMin() != i {
			t.Errorf("wrong answer %d", i)
		}
	}
}

func TestRandomOrder(t *testing.T) {
	t.Parallel()
	const size = 10000

	heap := New[int]()
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i
	}

	rand.Shuffle(len(ints), func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})

	for i := 0; i < size; i++ {
		heap.Insert(ints[i])
	}

	for i := 0; i < size; i++ {
		if value := heap.DeleteMin(); value != i {
			t.Errorf("wrong answer %d %d", value, i)
		}
	}
}

func TestRepeats(t *testing.T) {
	t.Parallel()
	heap := New[int]()

	heap.Insert(2)
	if value := heap.Min(); value != 2 {
		t.Errorf("wrong answer %d %d", value, 2)
	}

	if value := heap.DeleteMin(); value != 2 {
		t.Errorf("wrong answer %d %d", value, 2)
	}

	heap.Insert(math.MinInt)
	if value := heap.Min(); value != math.MinInt {
		t.Errorf("wrong answer %d %d", value, 2)
	}

	heap.Insert(math.MinInt)
	heap.Insert(math.MinInt)
	if value := heap.DeleteMin(); value != math.MinInt {
		t.Errorf("wrong answer %d %d", value, 2)
	}
	if value := heap.DeleteMin(); value != math.MinInt {
		t.Errorf("wrong answer %d %d", value, 2)
	}
	if value := heap.DeleteMin(); value != math.MinInt {
		t.Errorf("wrong answer %d %d", value, 2)
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	heap := New[int]()

	if value := heap.Min(); value != 0 {
		t.Errorf("wrong answer %d %d", value, 0)
	}

	if value := heap.Size(); value != 0 {
		t.Errorf("wrong answer %d %d", value, 0)
	}

	if value := heap.DeleteMin(); value != 0 {
		t.Errorf("wrong answer %d %d", value, 0)
	}

	if value := heap.Size(); value != 0 {
		t.Errorf("wrong answer %d %d", value, 0)
	}

	heap.Insert(1)
	if value := heap.Size(); value != 1 {
		t.Errorf("wrong answer %d %d", value, 1)
	}
	if value := heap.Min(); value != 1 {
		t.Errorf("wrong answer %d %d", value, 1)
	}

	if value := heap.DeleteMin(); value != 1 {
		t.Errorf("wrong answer %d %d", value, 1)
	}

	if value := heap.Size(); value != 0 {
		t.Errorf("wrong answer %d %d", value, 0)
	}
}
