package queue

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/heap"
	"github.com/nsnikhil/go-datastructures/internal"
)

type PriorityQueue[T comparable] struct {
	h heap.Heap[T]
}

func NewPriorityQueue[T comparable](isMax bool, c comparator.Comparator[T]) *PriorityQueue[T] {
	var h heap.Heap[T]

	if isMax {
		h = heap.NewMaxHeap[T](c)
	} else {
		h = heap.NewMinHeap[T](c)
	}

	return &PriorityQueue[T]{
		h: h,
	}
}

func (pq *PriorityQueue[T]) Add(e T) {
	pq.h.Add(e)
}

func (pq *PriorityQueue[T]) Remove() (T, error) {
	v, err := pq.h.Extract()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("PriorityQueue.Remove"), err)
	}

	return v, nil
}

//TODO: IMPLEMENT PEEK IN HEAP AND CHANGE IT HERE
func (pq *PriorityQueue[T]) Peek() (T, error) {
	v, err := pq.h.Iterator().Next()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("PriorityQueue.Peek"), err)
	}

	return v, nil
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.h.IsEmpty()
}

func (pq *PriorityQueue[T]) Size() int64 {
	return pq.h.Size()
}

func (pq *PriorityQueue[T]) Clear() {
	pq.h.Clear()
}
