package queue

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

type QueueNode struct {
	data interface{}
	next *QueueNode
}

func NewQueue() (q *Queue) {
	head := QueueNode{}
	return &Queue{&head, &head}
}

func (q *Queue) Push(data ...interface{}) {
	if q == nil || q.tail == nil {
		return
	}

	for _, d := range data {
		node := &QueueNode{data: d, next: nil}
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) Pop() interface{} {
	if q == nil || q.head == nil || q.head == q.tail {
		return nil
	}

	if q.head.next == q.tail {
		q.tail = q.head
	}
	data := q.head.next.data
	q.head.next = q.head.next.next
	return data
}
