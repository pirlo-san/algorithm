package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil || q.head == nil || q.head != q.tail {
		t.Fatalf("NewQueue error: %v %v %v", q, q.head, q.tail)
	}
}

func TestPushInt(t *testing.T) {
	q := NewQueue()
	q.Push(1, 2, 3, 4)
	want := []int{1, 2, 3, 4}
	got := []int{}
	for {
		d := q.Pop()
		if d == nil {
			break
		}
		got = append(got, d.(int))
	}
	if len(got) != len(want) {
		t.Fatalf("Push/Pop assert failure, want %v, got %v\n", want, got)
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("Push/Pop assert failure, want %v, got %v\n", want, got)
		}
	}
}

func TestPushString(t *testing.T) {
	q := NewQueue()
	q.Push("1", "22", "3", "AA")
	q.Push("BBB")
	want := []string{"1", "22", "3", "AA", "BBB"}
	got := []string{}
	for {
		d := q.Pop()
		if d == nil {
			break
		}
		got = append(got, d.(string))
	}
	if len(got) != len(want) {
		t.Fatalf("Push/Pop assert failure, want %v, got %v\n", want, got)
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("Push/Pop assert failure, want %v, got %v\n", want, got)
		}
	}
}
