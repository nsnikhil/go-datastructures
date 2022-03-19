package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

func buildHeap[T comparable](c comparator.Comparator[T], isMaxHeap bool, data []T) {
	sz := len(data)

	if sz == 0 {
		return
	}

	for i := sz / 2; i >= 0; i-- {
		heapify(i, c, isMaxHeap, data)
	}
}
