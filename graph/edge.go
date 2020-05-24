package graph

import (
	"fmt"
)

type Edges map[Edge]int

type Edge struct {
	from int
	to   int
}

func NewEdges() Edges {
	return make(map[Edge]int)
}

func (e *Edges) GetNum() (int, error) {
	if e == nil {
		return -1, fmt.Errorf("nil receiver")
	}

	result := 0
	for range *e {
		result += 1
	}
	return result, nil
}

func (e *Edges) AppendEdge(from, to, weight int) error {
	if e == nil {
		return fmt.Errorf("nil receiver")
	}

	(*e)[Edge{from, to}] = weight
	return nil
}

func (e *Edges) GetWeight(from, to int) (int, error) {
	if e == nil {
		return -1, fmt.Errorf("nil receiver")
	}

	if weight, ok := (*e)[Edge{from, to}]; !ok {
		return -1, fmt.Errorf("invalid vertex indexes: %v %v", from, to)
	} else {
		return weight, nil
	}
}
