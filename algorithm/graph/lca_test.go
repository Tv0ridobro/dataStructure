package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestLCA_LCA(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](9, true)
	g.AddEdgeDefault(0, 5)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(5, 2)
	g.AddEdgeDefault(5, 6)
	g.AddEdgeDefault(2, 7)
	g.AddEdgeDefault(2, 4)
	g.AddEdgeDefault(1, 3)
	g.AddEdgeDefault(1, 8)
	lc := NewLCA(g)
	if lc.LCA(0, 0) != 0 {
		t.Error("wrong answer")
	}
	if lc.LCA(6, 2) != 5 {
		t.Error("wrong answer")
	}
	if lc.LCA(6, 7) != 5 {
		t.Error("wrong answer")
	}
}

func TestLCA_LCA2(t *testing.T) {
	t.Parallel()
	g := graph.New[struct{}](16, true)
	g.AddEdgeDefault(0, 1)
	g.AddEdgeDefault(0, 2)
	g.AddEdgeDefault(0, 15)
	g.AddEdgeDefault(0, 4)
	g.AddEdgeDefault(2, 6)
	g.AddEdgeDefault(2, 11)
	g.AddEdgeDefault(15, 3)
	g.AddEdgeDefault(15, 5)
	g.AddEdgeDefault(3, 9)
	g.AddEdgeDefault(3, 8)
	g.AddEdgeDefault(3, 10)
	g.AddEdgeDefault(10, 12)
	g.AddEdgeDefault(12, 13)
	g.AddEdgeDefault(12, 14)
	lca := NewLCA(g)
	if value := lca.LCA(0, 0); value != 0 {
		t.Errorf("wrong asnwer %d != %d", value, 0)
	}
	if value := lca.LCA(3, 11); value != 0 {
		t.Errorf("wrong asnwer %d != %d", value, 0)
	}
	if value := lca.LCA(6, 11); value != 2 {
		t.Errorf("wrong asnwer %d != %d", value, 2)
	}
	if value := lca.LCA(1, 2); value != 0 {
		t.Errorf("wrong asnwer %d != %d", value, 0)
	}
	if value := lca.LCA(15, 8); value != 15 {
		t.Errorf("wrong asnwer %d != %d", value, 15)
	}
	if value := lca.LCA(9, 5); value != 15 {
		t.Errorf("wrong asnwer %d != %d", value, 15)
	}
	if value := lca.LCA(11, 2); value != 2 {
		t.Errorf("wrong asnwer %d != %d", value, 2)
	}
	if value := lca.LCA(8, 9); value != 3 {
		t.Errorf("wrong asnwer %d != %d", value, 3)
	}
	if value := lca.LCA(3, 5); value != 15 {
		t.Errorf("wrong asnwer %d != %d", value, 15)
	}
	if value := lca.LCA(0, 4); value != 0 {
		t.Errorf("wrong asnwer %d != %d", value, 0)
	}
}
