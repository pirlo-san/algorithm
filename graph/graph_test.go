package graph

import (
	"testing"
)

func TestVertexOperations(t *testing.T) {
	g := NewGraph()
	g.AddVertex("aaa")
	g.AddVertex("bbb", "ccc", "ddd")

	if want, got := 4, g.GetVertexNum(); got != want {
		t.Errorf("vertex number assertion failed, want %v, got %v\n", want, got)
	}

	g.DelVertex(2)
	g.DelVertex(0)

	if want, got := 2, g.GetVertexNum(); got != want {
		t.Errorf("vertex number assertion failed, want %v, got %v\n", want, got)
	}

}

/*
0 --> 1 --> 3
      |
	  \/
	  2
*/
func TestSimpleGraph(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0, 1, 2, 3)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	if want, got := 0, g.GetVertexInDegree(0); want != got {
		t.Errorf("vertex indegree assertion failure, want %v, got %v\n", want, got)
	}

	if want, got := 2, g.GetVertexOutDegree(1); want != got {
		t.Errorf("vertex outdegree assertion failure, want %v, got %v\n", want, got)
	}

	if want, got := 1, g.GetVertexInDegree(2); want != got {
		t.Errorf("vertex indegree assertion failure, want %v, got %v\n", want, got)
	}

	if want, got := 1, g.GetVertexInDegree(3); want != got {
		t.Errorf("vertex indegree assertion failure, want %v, got %v\n", want, got)
	}
}
