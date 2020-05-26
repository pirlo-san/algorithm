package graph

import (
	"fmt"

	"github.com/pirlo-san/algorithm/queue"
)

func (g *Graph) TopoSort() []interface{} {
	if g == nil {
		return nil
	}

	result := []interface{}{}
	q := queue.NewQueue()
	for idx, _ := range g.vertexes {
		if g.vertexes[idx].indegree == 0 {
			q.Push(idx)
		}
	}

	for {
		idx := q.Pop()
		if idx == nil {
			break
		}
		fmt.Println(idx)
		result = append(result, g.vertexes[idx.(int)].data)
		next := g.vertexes[idx.(int)].next
		for next != nil {
			g.vertexes[next.idx].indegree--
			if g.vertexes[next.idx].indegree == 0 {
				q.Push(next.idx)
			}
			next = next.next
		}
	}
	return result
}
