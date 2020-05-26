package graph

type Vertex struct {
	data      interface{}
	indegree  int
	outdegree int
	next      *VertexIndex
}

type VertexIndex struct {
	idx  int
	next *VertexIndex
}

func NewVertex(data interface{}) Vertex {
	return Vertex{data: data}
}

func (v *Vertex) AddNext(nextIndex int) {
	if v == nil {
		return
	}
	idx := VertexIndex{nextIndex, v.next}
	v.next = &idx
	v.outdegree += 1
}

func (v *Vertex) IncInDegree() {
	if v == nil {
		return
	}
	v.indegree += 1
}

func (v *Vertex) GetInDegree() int {
	if v == nil {
		return 0
	}

	return v.indegree
}

func (v *Vertex) GetOutDegree() int {
	if v == nil {
		return 0
	}

	return v.outdegree
}
