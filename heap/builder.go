package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

func buildHeap(c comparator.Comparator, isMaxHeap bool, data []interface{}) error {
	sz := len(data)

	for i := sz / 2; i >= 0; i-- {
		if err := heapify(i, c, isMaxHeap, data); err != nil {
			return err
		}
	}

	return nil
}
