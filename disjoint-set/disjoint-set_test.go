package disjointset

import (
	"testing"
)

func TestAllConnected(t *testing.T) {
	t.Parallel()
	ds := New(100)
	for i := 0; i < 99; i++ {
		ds.Union(i, i+1)
		if ds.Size(0) != i+2 {
			t.Errorf("%d != %d", ds.Size(0), i+2)
		}
	}
}

func TestInit(t *testing.T) {
	t.Parallel()
	ds := New(100)
	for i := 0; i < 100; i++ {
		a := ds.Get(i)
		if a != i {
			t.Errorf("%d != %d", a, i)
		}
	}
}

func TestDisjointSet_Add(t *testing.T) {
	t.Parallel()
	ds := New(0)
	ds.Get(-1)
	ds.Get(3)
	ds.Add()
	ds.Get(0)
}
