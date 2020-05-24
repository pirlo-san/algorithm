package graph

import (
	"testing"
)

func TestVertexOperations(t *testing.T) {
	g := NewGraph(false)
	g.AppendVertex("aaa")
	g.AppendVertex("bbb")
	g.AppendVertex("ccc")
	g.AppendVertex("aaa")

	if want, got := 4, g.GetVertexNum(); got != want {
		t.Errorf("vertex num assert failed, want %v, got %v\n", want, got)
	}

	g.DeleteVertex(2)
	g.DeleteVertex(0)

	if want, got := 2, g.GetVertexNum(); got != want {
		t.Errorf("vertex num assert failed, want %v, got %v\n", want, got)
	}

}

func TestAppendEdge(t *testing.T) {
	g := NewGraph(false)

	g.AppendVertex("A")
	g.AppendVertex("B")

	g.AppendEdge(0, 1, 100)
	want := 100
	if got, _ := g.GetEdgeWeight(0, 1); want != got {
		t.Errorf("incorrect edge weight, want %v, got %v", want, got)
	}

	want = 1
	if got, _ := g.GetEdgeNum(); want != got {
		t.Errorf("incorrect edge number, want %v, got %v\n", want, got)
	}
}
