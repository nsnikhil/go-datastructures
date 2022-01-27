package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type heapSort[T comparable] struct {
}

func newHeapSort[T comparable]() *heapSort[T] {
	return &heapSort[T]{}
}

func (hs *heapSort[T]) sort(c comparator.Comparator[T], isMaxHeap bool, data *[]T) error {
	h, err := buildHeapUtil(c, isMaxHeap, data)
	if err != nil {
		return err
	}

	sz := h.Size()
	temp := make([]T, sz)

	for i := 0; i < sz; i++ {
		ele, err := h.Extract()
		if err != nil {
			return err
		}

		temp[i] = ele
	}

	*data = temp

	return nil
}

func buildHeapUtil[T comparable](c comparator.Comparator[T], isMaxHeap bool, data *[]T) (Heap[T], error) {
	if isMaxHeap {
		return NewMaxHeap[T](c, *data...)
	}

	return NewMinHeap[T](c, *data...)
}
