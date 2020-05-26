package graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g := NewGraph()

	g.AddVertex(0, 1, 2, 3)

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	seq := g.TopoSort()
	t.Errorf("%v", seq)
}
