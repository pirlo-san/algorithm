package graph

import (
	"fmt"
)

type Graph struct {
	vertexes []Vertex
	edges    []Edge
}

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

type Edge struct {
	from   int
	to     int
	weight int
}

func (g *Graph) AppendVertex(data interface{}) error {
	if g == nil || data == nil {
		return fmt.Errorf("nil input: graph(%v), data(%v)", g, data)
	}

	g.vertexes = append(g.vertexes, Vertex{data: data})
	return nil
}

func (g *Graph) DeleteVertex(index int) error {
	if g == nil {
		return fmt.Errorf("nil receiver")
	}

	if index < 0 || index >= g.GetVertexNum() {
		return fmt.Errorf("invalid vertex index: %v", index)
	}
	g.vertexes = append(g.vertexes[0:index], g.vertexes[index+1:]...)
	return nil
}

func (g *Graph) GetVertexNum() int {
	if g == nil {
		return 0
	}
	return len(g.vertexes)
}
