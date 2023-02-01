package thin_heap

import (
	"math"
	"math/rand"
	"testing"
)

func TestName(t *testing.T) {
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

func TestName2(t *testing.T) {
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
		//fmt.Println(heap.Min())
		if value := heap.DeleteMin(); value != i {
			t.Errorf("wrong answer %d %d", value, i)
		}
	}
}

func TestName3(t *testing.T) {
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
