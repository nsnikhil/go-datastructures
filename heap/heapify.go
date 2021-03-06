package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

func heapify(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}, indexes map[interface{}]int) error {
	return heapUtil(curr, c, maxHeapify, data, indexes)
}

func heapUtil(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}, indexes map[interface{}]int) error {
	if curr == len(data)-1 {
		return shiftUp(curr, c, maxHeapify, data, indexes)
	}
	return shiftDown(curr, c, maxHeapify, data, indexes)
}

func shiftUp(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}, indexes map[interface{}]int) error {
	if curr == 0 {
		return nil
	}

	shouldSwap, parent, err := shouldSwapWithParent(curr, c, maxHeapify, data)
	if err != nil {
		return err
	}

	for curr > 0 && shouldSwap {
		data[curr], data[parent] = data[parent], data[curr]
		indexes[data[curr]], indexes[data[parent]] = indexes[data[parent]], indexes[data[curr]]

		curr = parent

		if curr <= 0 {
			break
		}

		shouldSwap, parent, err = shouldSwapWithParent(curr, c, maxHeapify, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func shiftDown(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}, indexes map[interface{}]int) error {
	if curr >= len(data)/2 {
		return nil
	}

	shouldSwap, child, err := shouldSwapWithChild(curr, c, maxHeapify, data)
	if err != nil {
		return err
	}

	for curr < len(data)/2 && shouldSwap {
		data[curr], data[child] = data[child], data[curr]
		indexes[data[curr]], indexes[data[child]] = indexes[data[child]], indexes[data[curr]]

		curr = child

		if curr >= len(data)/2 {
			break
		}

		shouldSwap, child, err = shouldSwapWithChild(curr, c, maxHeapify, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func shouldSwapWithParent(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}) (bool, int, error) {
	if curr == 0 {
		return false, invalidIndex, nil
	}

	parent := parentIndex(curr)

	diff, err := c.Compare(data[parent], data[curr])
	if err != nil {
		return false, invalidIndex, err
	}

	if maxHeapify {
		return diff < 0, parent, nil
	}

	return diff > 0, parent, nil
}

func shouldSwapWithChild(curr int, c comparator.Comparator, maxHeapify bool, data []interface{}) (bool, int, error) {
	lcIndex := leftChildIndex(curr)
	leftDiff, err := c.Compare(data[curr], data[lcIndex])
	if err != nil {
		return false, invalidIndex, err
	}

	hasRC := hasRightChild(curr, len(data))
	var rcIndex, rightDiff int
	if hasRC {
		rcIndex = rightChildIndex(curr)
		rd, err := c.Compare(data[curr], data[rcIndex])
		if err != nil {
			return false, invalidIndex, err
		}
		rightDiff = rd
	}

	if maxHeapify {
		return shouldSwapWithChildMaxUtil(hasRC, leftDiff, rightDiff, lcIndex, rcIndex)
	}

	return shouldSwapWithChildMinUtil(hasRC, leftDiff, rightDiff, lcIndex, rcIndex)
}

func shouldSwapWithChildMaxUtil(hasRC bool, leftDiff, rightDiff, lcIndex, rcIndex int) (bool, int, error) {
	if hasRC {
		if leftDiff > 0 && rightDiff > 0 {
			return false, invalidIndex, nil
		}

		if leftDiff < rightDiff {
			return true, lcIndex, nil
		}

		return true, rcIndex, nil
	}

	if leftDiff > 0 {
		return false, invalidIndex, nil
	}

	return true, lcIndex, nil
}

//TODO MERGE WITH shouldSwapWithChildMaxUtil
func shouldSwapWithChildMinUtil(hasRC bool, leftDiff, rightDiff, lcIndex, rcIndex int) (bool, int, error) {
	if hasRC {
		if leftDiff < 0 && rightDiff < 0 {
			return false, invalidIndex, nil
		}

		if leftDiff > rightDiff {
			return true, lcIndex, nil
		}

		return true, rcIndex, nil
	}

	if leftDiff < 0 {
		return false, invalidIndex, nil
	}

	return true, lcIndex, nil
}

func hasRightChild(curr, sz int) bool {
	return rightChildIndex(curr) < sz
}

func parentIndex(curr int) int {
	if curr%2 == 0 {
		return (curr - 1) / 2
	}
	return curr / 2
}

func leftChildIndex(curr int) int {
	return (curr * 2) + 1
}

func rightChildIndex(curr int) int {
	return (curr * 2) + 2
}
