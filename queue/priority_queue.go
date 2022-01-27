package queue

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/heap"
)

type PriorityQueue[T comparable] struct {
	h heap.Heap[T]
}

func NewPriorityQueue[T comparable](isMax bool, c comparator.Comparator[T]) (*PriorityQueue[T], error) {
	var h heap.Heap[T]
	var err error

	if isMax {
		h, err = heap.NewMaxHeap[T](c)
	} else {
		h, err = heap.NewMinHeap[T](c)
	}

	if err != nil {
		return nil, err
	}

	return &PriorityQueue[T]{
		h: h,
	}, nil
}

func (pq *PriorityQueue[T]) Add(e T) error {
	return pq.h.Add(e)
}

func (pq *PriorityQueue[T]) Remove() (T, error) {
	return pq.h.Extract()
}

func (pq *PriorityQueue[T]) Update(prev, new T) error {
	return pq.h.Update(prev, new)
}

func (pq *PriorityQueue[T]) UpdateFunc(prev T, op func(T) T) error {
	return pq.h.UpdateFunc(prev, op)
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	return pq.h.Iterator().Next()
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.h.IsEmpty()
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.h.Size()
}

func (pq *PriorityQueue[T]) Clear() {
	pq.h.Clear()
}
