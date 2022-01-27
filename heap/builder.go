package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

func buildHeap[T comparable](c comparator.Comparator[T], isMaxHeap bool, data []T, indexes map[T]int) error {
	sz := len(data)

	for i, d := range data {
		indexes[d] = i
	}

	for i := sz / 2; i >= 0; i-- {
		if err := heapify(i, c, isMaxHeap, data, indexes); err != nil {

			//TODO: change deletion implementation, New Assignment did not work.
			for k := range indexes {
				delete(indexes, k)
			}

			return err
		}
	}

	return nil
}
