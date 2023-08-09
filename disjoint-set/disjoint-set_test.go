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
	if v := ds.Get(-1); v != -1 {
		t.Errorf("%d != %d", v, -1)
	}
	if v := ds.Get(3); v != -1 {
		t.Errorf("%d != %d", v, -1)
	}
	if v := ds.Get(0); v != -1 {
		t.Errorf("%d != %d", v, -1)
	}
	ds.Add()
	if v := ds.Get(0); v != 0 {
		t.Errorf("%d != %d", v, 0)
	}
}

func TestDisjointSet_Components(t *testing.T) {
	t.Parallel()
	ds := New(10)
	if v := ds.Components(); v != 10 {
		t.Errorf("%d != %d", v, 10)
	}
	ds.Add()
	if v := ds.Components(); v != 11 {
		t.Errorf("%d != %d", v, 11)
	}
	ds.Union(1, 2)
	if v := ds.Components(); v != 10 {
		t.Errorf("%d != %d", v, 10)
	}
	ds.Union(3, 4)
	if v := ds.Components(); v != 9 {
		t.Errorf("%d != %d", v, 9)
	}
	ds.Union(1, 3)
	if v := ds.Components(); v != 8 {
		t.Errorf("%d != %d", v, 8)
	}
	ds.Union(2, 4)
	if v := ds.Components(); v != 8 {
		t.Errorf("%d != %d", v, 8)
	}
}
