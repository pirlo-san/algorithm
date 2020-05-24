package graph

import (
	"fmt"
)

type Vertex struct {
	data      interface{}
	indegree  int
	outdegree int
	nexts     *VertexIndex
}

type VertexIndex struct {
	idx  int
	next *VertexIndex
}

func NewVertex(data interface{}) Vertex {
	return Vertex{data: data}
}

func (v *Vertex) AddNext(nextIndex int) error {
	if v == nil {
		return fmt.Errorf("nil receiver")
	}
	idx := VertexIndex{nextIndex, v.nexts}
	v.nexts = &idx
	v.outdegree += 1
	return nil
}

func (v *Vertex) IncInDegree() error {
	if v == nil {
		return fmt.Errorf("nil receiver")
	}
	v.indegree += 1
	return nil
}
