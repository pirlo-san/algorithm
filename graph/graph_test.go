package graph

import (
	"testing"
)

func TestVertexOperations(t *testing.T) {
	g := &Graph{}
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
