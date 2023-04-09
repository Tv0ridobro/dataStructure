package sparsetable

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/math"
)

func TestSparseTable_Query(t *testing.T) {
	t.Parallel()
	table := New(math.GCD, []int{2, 3, 5, 4, 6, 8, 16})
	if v := table.Query(0, 2); v != 1 {
		t.Errorf("wrong answer for query %d %d", v, 1)
	}
	if v := table.Query(3, 5); v != 2 {
		t.Errorf("wrong answer for query %d %d", v, 2)
	}
	if v := table.Query(2, 3); v != 1 {
		t.Errorf("wrong answer for query %d %d", v, 1)
	}
	if v := table.Query(5, 6); v != 8 {
		t.Errorf("wrong answer for query %d %d", v, 8)
	}
	if v := table.Query(2, 2); v != 5 {
		t.Errorf("wrong answer for query %d %d", v, 5)
	}
}
