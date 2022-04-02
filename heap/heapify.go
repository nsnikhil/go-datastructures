package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/internal"
)

func heapify[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) {
	heapUtil(curr, c, maxHeapify, data)
}

func heapUtil[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) {
	if curr == len(data)-1 {
		shiftUp(curr, c, maxHeapify, data)
	} else {
		shiftDown(curr, c, maxHeapify, data)
	}
}

func shiftUp[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) {
	if curr == 0 {
		return
	}

	shouldSwap, parent := shouldSwapWithParent(curr, c, maxHeapify, data)

	for curr > 0 && shouldSwap {
		data[curr], data[parent] = data[parent], data[curr]

		curr = parent

		if curr <= 0 {
			break
		}

		shouldSwap, parent = shouldSwapWithParent(curr, c, maxHeapify, data)
	}

}

func shiftDown[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) {
	if curr >= len(data)/2 {
		return
	}

	shouldSwap, child := shouldSwapWithChild(curr, c, maxHeapify, data)

	for curr < len(data)/2 && shouldSwap {
		data[curr], data[child] = data[child], data[curr]

		curr = child

		if curr >= len(data)/2 {
			break
		}

		shouldSwap, child = shouldSwapWithChild(curr, c, maxHeapify, data)
	}
}

func shouldSwapWithParent[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) (bool, int) {
	if curr == 0 {
		return false, internal.InvalidIndex
	}

	parent := parentIndex(curr)

	diff := c.Compare(data[parent], data[curr])

	if maxHeapify {
		return diff < 0, parent
	}

	return diff > 0, parent
}

func shouldSwapWithChild[T any](curr int, c comparator.Comparator[T], maxHeapify bool, data []T) (bool, int) {
	lcIndex := leftChildIndex(curr)
	leftDiff := c.Compare(data[curr], data[lcIndex])

	hasRC := hasRightChild(curr, len(data))
	var rcIndex, rightDiff int
	if hasRC {
		rcIndex = rightChildIndex(curr)
		rd := c.Compare(data[curr], data[rcIndex])
		rightDiff = rd
	}

	if maxHeapify {
		return shouldSwapWithChildMaxUtil(hasRC, leftDiff, rightDiff, lcIndex, rcIndex)
	}

	return shouldSwapWithChildMinUtil(hasRC, leftDiff, rightDiff, lcIndex, rcIndex)
}

func shouldSwapWithChildMaxUtil(hasRC bool, leftDiff, rightDiff, lcIndex, rcIndex int) (bool, int) {
	if hasRC {
		if leftDiff > 0 && rightDiff > 0 {
			return false, internal.InvalidIndex
		}

		if leftDiff < rightDiff {
			return true, lcIndex
		}

		return true, rcIndex
	}

	if leftDiff > 0 {
		return false, internal.InvalidIndex
	}

	return true, lcIndex
}

//TODO MERGE WITH shouldSwapWithChildMaxUtil
func shouldSwapWithChildMinUtil(hasRC bool, leftDiff, rightDiff, lcIndex, rcIndex int) (bool, int) {
	if hasRC {
		if leftDiff < 0 && rightDiff < 0 {
			return false, internal.InvalidIndex
		}

		if leftDiff > rightDiff {
			return true, lcIndex
		}

		return true, rcIndex
	}

	if leftDiff < 0 {
		return false, internal.InvalidIndex
	}

	return true, lcIndex
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
