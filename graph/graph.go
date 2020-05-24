package graph

import (
	"fmt"
)

type Graph struct {
	vertexes []Vertex
	edges    Edges
	directed bool
}

func NewGraph(directed bool) *Graph {
	return &Graph{directed: directed, edges: NewEdges()}
}

func (g *Graph) GetVertexNum() int {
	if g == nil {
		return 0
	}
	return len(g.vertexes)
}

func (g *Graph) GetEdgeNum() (int, error) {
	if g == nil {
		return -1, fmt.Errorf("nil receiver")
	}

	result, err := g.edges.GetNum()
	if err != nil {
		return result, err
	}
	if !g.directed {
		result /= 2
	}
	return result, nil
}

func (g *Graph) IsValidVertexIndex(index int) bool {
	return 0 <= index && index < g.GetVertexNum()
}

func (g *Graph) AppendVertex(data interface{}) error {
	if g == nil || data == nil {
		return fmt.Errorf("nil input: graph(%v), data(%v)", g, data)
	}

	g.vertexes = append(g.vertexes, NewVertex(data))
	return nil
}

func (g *Graph) DeleteVertex(index int) error {
	if g == nil {
		return fmt.Errorf("nil receiver")
	}

	if !g.IsValidVertexIndex(index) {
		return fmt.Errorf("invalid vertex index: %v", index)
	}
	g.vertexes = append(g.vertexes[0:index], g.vertexes[index+1:]...)
	return nil
}

func (g *Graph) AppendEdge(from int, to int, weight int) error {
	if g == nil {
		return fmt.Errorf("nil receiver")
	}

	if !g.IsValidVertexIndex(from) || !g.IsValidVertexIndex(to) {
		return fmt.Errorf("invalid vertex index: %v %v", from, to)
	}

	g.edges.AppendEdge(from, to, weight)
	fromVertex, toVertex := g.vertexes[from], g.vertexes[to]
	if err := fromVertex.AddNext(to); err != nil {
		return err
	}
	if err := toVertex.IncInDegree(); err != nil {
		return err
	}

	if g.directed {
		return nil
	}

	g.edges.AppendEdge(to, from, weight)
	if err := toVertex.AddNext(from); err != nil {
		return err
	}
	return fromVertex.IncInDegree()
}

func (g *Graph) GetEdgeWeight(from, to int) (int, error) {
	if g == nil {
		return -1, fmt.Errorf("nil receiver")
	}

	return g.edges.GetWeight(from, to)
}
