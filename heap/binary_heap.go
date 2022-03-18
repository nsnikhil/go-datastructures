package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
)

type binaryHeap[T comparable] struct {
	c         comparator.Comparator[T]
	isMaxHeap bool
	data      []T
}

type MaxHeap[T comparable] struct {
	*binaryHeap[T]
}

func NewMaxHeap[T comparable](c comparator.Comparator[T], data ...T) *MaxHeap[T] {
	heap := newBinaryHeap(c, true, data...)

	return &MaxHeap[T]{binaryHeap: heap}
}

type MinHeap[T comparable] struct {
	*binaryHeap[T]
}

func NewMinHeap[T comparable](c comparator.Comparator[T], data ...T) *MinHeap[T] {
	heap := newBinaryHeap(c, false, data...)

	return &MinHeap[T]{binaryHeap: heap}
}

func newBinaryHeap[T comparable](c comparator.Comparator[T], isMaxHeap bool, data ...T) *binaryHeap[T] {
	if len(data) == 0 {
		return &binaryHeap[T]{
			c:         c,
			isMaxHeap: isMaxHeap,
		}
	}

	buildHeap(c, isMaxHeap, data)

	return &binaryHeap[T]{
		c:         c,
		isMaxHeap: isMaxHeap,
		data:      data,
	}
}

func (bh *binaryHeap[T]) Add(data ...T) {
	for _, d := range data {
		bh.data = append(bh.data, d)

		heapify(len(bh.data)-1, bh.c, bh.isMaxHeap, bh.data)
	}

}

func (bh *binaryHeap[T]) Extract() (T, error) {
	if bh.IsEmpty() {
		return *new(T), emptyHeapError("binaryHeap.Extract")
	}

	ele := bh.data[0]

	bh.data[0] = bh.data[bh.Size()-1]

	bh.data = bh.data[:bh.Size()-1]

	heapify(0, bh.c, bh.isMaxHeap, bh.data)

	return ele, nil
}

func (bh *binaryHeap[T]) Delete() error {
	if _, err := bh.Extract(); err != nil {
		return err
	}

	return nil
}

func (bh *binaryHeap[T]) Size() int {
	return len(bh.data)
}

func (bh *binaryHeap[T]) IsEmpty() bool {
	return bh.Size() == 0
}

func (bh *binaryHeap[T]) Clear() {
	bh.data = nil
}

func (bh *binaryHeap[T]) Iterator() iterator.Iterator[T] {
	return newBinaryHeapIterator[T](bh)
}

type binaryHeapIterator[T comparable] struct {
	currentIndex int
	h            *binaryHeap[T]
}

func newBinaryHeapIterator[T comparable](bh *binaryHeap[T]) *binaryHeapIterator[T] {
	return &binaryHeapIterator[T]{
		currentIndex: 0,
		h:            bh,
	}
}

func (bhi *binaryHeapIterator[T]) HasNext() bool {
	return bhi.currentIndex != bhi.h.Size()
}

func (bhi *binaryHeapIterator[T]) Next() (T, error) {
	if bhi.currentIndex >= bhi.h.Size() {
		return *new(T), emptyIteratorError("binaryHeapIterator.Next")
	}

	e := bhi.h.data[bhi.currentIndex]
	bhi.currentIndex++

	return e, nil
}
