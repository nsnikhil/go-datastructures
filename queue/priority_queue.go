package queue

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/heap"
)

type PriorityQueue struct {
	h heap.Heap
}

func NewPriorityQueue(isMax bool, c comparator.Comparator) (*PriorityQueue, error) {
	var h heap.Heap
	var err error

	if isMax {
		h, err = heap.NewMaxHeap(c)
	} else {
		h, err = heap.NewMinHeap(c)
	}

	if err != nil {
		return nil, err
	}

	return &PriorityQueue{
		h: h,
	}, nil
}

func (pq *PriorityQueue) Add(e interface{}) error {
	return pq.h.Add(e)
}

func (pq *PriorityQueue) Remove() (interface{}, error) {
	return pq.h.Extract()
}

func (pq *PriorityQueue) Update(prev, new interface{}) error {
	return pq.h.Update(prev, new)
}

func (pq *PriorityQueue) UpdateFunc(prev interface{}, op func(interface{}) interface{}) error {
	return pq.h.UpdateFunc(prev, op)
}

func (pq *PriorityQueue) Peek() (interface{}, error) {
	return pq.h.Iterator().Next(), nil
}

func (pq *PriorityQueue) Empty() bool {
	return pq.h.IsEmpty()
}

func (pq *PriorityQueue) Count() int {
	return pq.h.Size()
}

func (pq *PriorityQueue) Clear() {
	pq.h.Clear()
}
