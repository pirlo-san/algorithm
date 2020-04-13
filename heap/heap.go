/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodeHeap struct {
	nodes []*ListNode
}

func (h *ListNodeHeap) Len() int {
	return len(h.nodes)
}

func (h *ListNodeHeap) Less(i, j int) bool {
	return h.nodes[i].Val < h.nodes[j].Val
}

func (h *ListNodeHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	h.nodes = append(h.nodes, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := h.nodes
	n := len(old)
	x := old[n-1]
	h.nodes = old[0 : n-1]
	return x
}

// LeetCode 23
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) <= 0 {
		return nil
	}
	listNodeHeap := &ListNodeHeap{}
	heap.Init(listNodeHeap)
	k := len(lists)
	for i := 0; i < k; i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(listNodeHeap, lists[i])
	}

	res := ListNode{}
	p := &res
	for len(listNodeHeap.nodes) > 0 {
		node := heap.Pop(listNodeHeap).(*ListNode)
		p.Next = &ListNode{node.Val, nil}
		p = p.Next
		node = node.Next
		if node != nil {
			heap.Push(listNodeHeap, node)
		}
	}

	return res.Next
}