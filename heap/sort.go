package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type heapSort struct {
}

func newHeapSort() *heapSort {
	return &heapSort{}
}

func (hs *heapSort) sort(c comparator.Comparator, isMaxHeap bool, data *[]interface{}) error {
	h, err := buildHeapUtil(c, isMaxHeap, data)
	if err != nil {
		return err
	}

	sz := h.Size()
	temp := make([]interface{}, sz)

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

func buildHeapUtil(c comparator.Comparator, isMaxHeap bool, data *[]interface{}) (Heap, error) {
	if isMaxHeap {
		return NewMaxHeap(c, *data...)
	}

	return NewMinHeap(c, *data...)
}
