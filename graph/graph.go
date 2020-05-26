package graph

type Graph struct {
	vertexes []Vertex
}

func NewGraph() Graph {
	return Graph{}
}

func (g *Graph) GetVertexNum() int {
	if g == nil {
		return 0
	}
	return len(g.vertexes)
}

func (g *Graph) GetVertexInDegree(idx int) int {
	if g == nil || !g.IsValidVertexIndex(idx) {
		return 0
	}

	return g.vertexes[idx].GetInDegree()
}

func (g *Graph) GetVertexOutDegree(idx int) int {
	if g == nil || !g.IsValidVertexIndex(idx) {
		return 0
	}

	return g.vertexes[idx].GetOutDegree()
}

func (g *Graph) IsValidVertexIndex(index int) bool {
	return 0 <= index && index < g.GetVertexNum()
}

func (g *Graph) AddVertex(data ...interface{}) {
	if g == nil {
		return
	}

	for _, d := range data {
		g.vertexes = append(g.vertexes, NewVertex(d))
	}
}

func (g *Graph) DelVertex(index int) {
	if g == nil || !g.IsValidVertexIndex(index) {
		return
	}

	for neighbor := g.vertexes[index].next; neighbor != nil; neighbor = neighbor.next {
		g.vertexes[neighbor.idx].indegree -= 1
	}
	g.vertexes = append(g.vertexes[0:index], g.vertexes[index+1:]...)
}

func (g *Graph) AddEdge(from, to int) {
	if g == nil || !g.IsValidVertexIndex(from) || !g.IsValidVertexIndex(to) {
		return
	}

	g.vertexes[from].AddNext(to)
	g.vertexes[to].IncInDegree()
}
