package heap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
)

type binaryHeap[T comparable] struct {
	c         comparator.Comparator[T]
	isMaxHeap bool
	indexes   map[T]int
	data      []T
}

type MaxHeap[T comparable] struct {
	*binaryHeap[T]
}

func NewMaxHeap[T comparable](c comparator.Comparator[T], data ...T) (*MaxHeap[T], error) {
	heap, err := newBinaryHeap(c, true, data...)
	if err != nil {
		return nil, err
	}

	return &MaxHeap[T]{binaryHeap: heap}, nil
}

type MinHeap[T comparable] struct {
	*binaryHeap[T]
}

func NewMinHeap[T comparable](c comparator.Comparator[T], data ...T) (*MinHeap[T], error) {
	heap, err := newBinaryHeap(c, false, data...)
	if err != nil {
		return nil, err
	}

	return &MinHeap[T]{binaryHeap: heap}, nil
}

func newBinaryHeap[T comparable](c comparator.Comparator[T], isMaxHeap bool, data ...T) (*binaryHeap[T], error) {
	if len(data) == 0 {
		return &binaryHeap[T]{
			c:         c,
			isMaxHeap: isMaxHeap,
			indexes:   make(map[T]int),
		}, nil
	}

	indexes := make(map[T]int)
	if err := buildHeap(c, isMaxHeap, data, indexes); err != nil {
		return nil, err
	}

	return &binaryHeap[T]{
		c:         c,
		isMaxHeap: isMaxHeap,
		data:      data,
		indexes:   indexes,
	}, nil
}

func (bh *binaryHeap[T]) Add(data ...T) error {
	for _, d := range data {
		bh.data = append(bh.data, d)
		bh.indexes[d] = len(bh.data) - 1

		if err := heapify(len(bh.data)-1, bh.c, bh.isMaxHeap, bh.data, bh.indexes); err != nil {
			return err
		}
	}
	return nil
}

func (bh *binaryHeap[T]) Extract() (T, error) {
	if bh.IsEmpty() {
		return *new(T), errors.New("heap is empty")
	}

	ele := bh.data[0]

	bh.data[0] = bh.data[bh.Size()-1]
	bh.indexes[bh.data[bh.Size()-1]] = 0

	bh.data = bh.data[:bh.Size()-1]
	delete(bh.indexes, ele)

	if err := heapify(0, bh.c, bh.isMaxHeap, bh.data, bh.indexes); err != nil {
		return *new(T), err
	}

	return ele, nil
}

func (bh *binaryHeap[T]) Update(prev, new T) error {
	if bh.IsEmpty() {
		return errors.New("heap is empty")
	}

	idx, ok := bh.indexes[prev]
	if !ok {
		return fmt.Errorf("%v not found in heap", prev)
	}

	if prev == new {
		return fmt.Errorf("%v and %v are same", prev, new)
	}

	res := bh.c.Compare(prev, new)
	if res == 0 {
		return fmt.Errorf("comparator returned same for %v and %v", prev, new)
	}

	delete(bh.indexes, prev)

	bh.data[idx] = new
	bh.indexes[new] = idx

	if bh.isMaxHeap && res > 0 || !bh.isMaxHeap && res < 0 {
		return shiftDown(idx, bh.c, bh.isMaxHeap, bh.data, bh.indexes)
	}

	return shiftUp(idx, bh.c, bh.isMaxHeap, bh.data, bh.indexes)
}

func (bh *binaryHeap[T]) UpdateFunc(prev T, op func(T) T) error {
	if bh.IsEmpty() {
		return errors.New("heap is empty")
	}

	idx, ok := bh.indexes[prev]
	if !ok {
		return fmt.Errorf("%v not found in heap", prev)
	}

	updated := op(prev)

	delete(bh.indexes, prev)

	bh.data[idx] = updated
	bh.indexes[updated] = idx

	if bh.Size() == 1 {
		return nil
	}

	if err := shiftDown(idx, bh.c, bh.isMaxHeap, bh.data, bh.indexes); err != nil {
		return err
	}

	return shiftUp(idx, bh.c, bh.isMaxHeap, bh.data, bh.indexes)
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
		return *new(T), errors.New("") //TODO: FILL THIS
	}

	e := bhi.h.data[bhi.currentIndex]
	bhi.currentIndex++

	return e, nil
}
