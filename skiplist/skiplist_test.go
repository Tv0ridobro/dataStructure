package skiplist

import (
	"math/rand"
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestSkipList_Insert(t *testing.T) {
	t.Parallel()
	sl := NewWithProbability[int](0.9)
	for i := 0; i < 100; i++ {
		sl.Insert(rand.Intn(1000))
	}
}

func TestSkipList_All(t *testing.T) {
	t.Parallel()
	sl := New[int]()
	if !slices.Equal(sl.All(), []int{}) {
		t.Errorf("wrong all %v", []int{})
	}
	sl.Insert(0)
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(3)
	sl.Insert(4)
	if !slices.Equal(sl.All(), []int{0, 1, 2, 3, 4}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 3, 4})
	}
	sl.Remove(3)
	if !slices.Equal(sl.All(), []int{0, 1, 2, 4}) {
		t.Errorf("wrong all %v", []int{0, 1, 2, 4})
	}
}
