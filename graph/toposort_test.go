package graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g := NewGraph()

	g.AddVertex(0, 1, 2, 3, 4, 5, 6, 7)

	g.AddEdge(0, 2)
	g.AddEdge(2, 4)
	g.AddEdge(1, 3)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(4, 6)
	g.AddEdge(5, 7)
	g.AddEdge(6, 7)

	wants := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7},
	}
	seq := g.TopoSort()
	match := false
	for _, want := range wants {
		if equalIntSlice(want, seq) {
			match = true
			break
		}
	}

	if !match {
		t.Errorf("incorrect toposort: %v\n", seq)
	}

}

func equalIntSlice(s1 []int, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i].(int) {
			return false
		}
	}

	return true
}
