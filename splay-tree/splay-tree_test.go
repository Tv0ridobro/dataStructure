package splaytree

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestAll(t *testing.T) {
	t.Parallel()
	tr := New[int]()
	for i := 0; i < 100; i++ {
		tr.Insert(i)
		if tr.Size() != i+1 {
			t.Errorf("wrong size expected %d  got %d", i+1, i)
		}
	}
	for i := 0; i < 100; i++ {
		if !tr.Contains(i) {
			t.Errorf("doesnt contain %d", i)
		}
	}
	for i := 0; i < 100; i++ {
		if i%7 == 4 {
			tr.Remove(i)
		}
		tr.Remove(i + 200)
	}
	for i := 0; i < 100; i++ {
		if i%7 == 4 {
			if tr.Contains(i) {
				t.Errorf("contain %d", i)
			}
		}
	}
}

func TestAll2(t *testing.T) {
	t.Parallel()
	tr := New[int]()
	permutation := rand.Perm(1000000)
	for i := range permutation {
		tr.Insert(i)
	}
	sort.Ints(permutation)
	if !slices.Equal(tr.All(), permutation) {
		t.Errorf("permutation doesn't equal All call")
	}
	if tr.Size() != 1000000 {
		t.Errorf("wrong size")
	}
	rand.Shuffle(len(permutation), func(i, j int) {
		permutation[i], permutation[j] = permutation[j], permutation[i]
	})
	for i := range permutation {
		tr.Remove(i)
	}
	if tr.root != nil {
		t.Errorf("root is not nil")
	}
}

func TestSplayTree_Kth(t *testing.T) {
	t.Parallel()
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Insert(i)
	}
	for i := 0; i < 1000; i++ {
		if v := s.Kth(i); v != i {
			t.Errorf("Kth != ans %d %d", v, i)
		}
	}
}

func TestSplayTree_Sub(t *testing.T) {
	t.Parallel()
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Insert(i)
	}
	for i := 0; i < 1000; i++ {
		s.Sub(-1, 0)
		s.Sub(0, 0)
		s.Sub(0, 1000)
		s.Sub(0, 2000)
	}
}

func TestSplayTree_Empty(t *testing.T) {
	t.Parallel()
	s := New[int]()
	s.Size()
	s.Contains(1)
	s.Remove(2)
	s.Kth(3)
	s.All()
}

func TestNewWithComparator(t *testing.T) {
	t.Parallel()
	s := NewWithComparator(func(a int8, b int8) int {
		switch {
		case a > b:
			return -1
		case a == b:
			return 0
		default:
			return 1
		}
	})
	for i := int8(0); i < 10; i++ {
		s.Insert(i)
	}
	if !slices.Equal(s.All(), []int8{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Errorf("slices are not equal")
	}
}
