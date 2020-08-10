package heap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/utils"
)

type binaryHeap struct {
	typeURL   string
	c         comparator.Comparator
	isMaxHeap bool
	indexes   map[interface{}]int
	data      []interface{}
}

type MaxHeap struct {
	*binaryHeap
}

func NewMaxHeap(c comparator.Comparator, data ...interface{}) (*MaxHeap, error) {
	heap, err := newBinaryHeap(c, true, data...)
	if err != nil {
		return nil, err
	}

	return &MaxHeap{binaryHeap: heap}, nil
}

type MinHeap struct {
	*binaryHeap
}

func NewMinHeap(c comparator.Comparator, data ...interface{}) (*MinHeap, error) {
	heap, err := newBinaryHeap(c, false, data...)
	if err != nil {
		return nil, err
	}

	return &MinHeap{binaryHeap: heap}, nil
}

func newBinaryHeap(c comparator.Comparator, isMaxHeap bool, data ...interface{}) (*binaryHeap, error) {
	if len(data) == 0 {
		return &binaryHeap{
			c:         c,
			typeURL:   na,
			isMaxHeap: isMaxHeap,
			indexes:   make(map[interface{}]int),
		}, nil
	}

	typeURL := utils.GetTypeName(data[0])

	for i := 1; i < len(data); i++ {
		if utils.GetTypeName(data[i]) != typeURL {
			return nil, liberror.NewTypeMismatchError(typeURL, utils.GetTypeName(data[i]))
		}
	}

	indexes := make(map[interface{}]int)
	if err := buildHeap(c, isMaxHeap, data, indexes); err != nil {
		return nil, err
	}

	return &binaryHeap{
		c:         c,
		typeURL:   typeURL,
		isMaxHeap: isMaxHeap,
		data:      data,
		indexes:   indexes,
	}, nil
}

func (bh *binaryHeap) Add(data ...interface{}) error {
	s := 0
	typeURL := bh.typeURL

	if typeURL == na {
		s++
		typeURL = utils.GetTypeName(data[0])
	}

	for i := s; i < len(data); i++ {
		if utils.GetTypeName(data[i]) != typeURL {
			return liberror.NewTypeMismatchError(typeURL, utils.GetTypeName(data[i]))
		}
	}

	if bh.typeURL == na {
		bh.typeURL = typeURL
	}

	for _, d := range data {
		bh.data = append(bh.data, d)
		bh.indexes[d] = len(bh.data) - 1

		if err := heapify(len(bh.data)-1, bh.c, bh.isMaxHeap, bh.data, bh.indexes); err != nil {
			return err
		}
	}
	return nil
}

func (bh *binaryHeap) Extract() (interface{}, error) {
	if bh.IsEmpty() {
		return nil, errors.New("heap is empty")
	}

	ele := bh.data[0]

	bh.data[0] = bh.data[bh.Size()-1]
	bh.indexes[bh.data[bh.Size()-1]] = 0

	bh.data = bh.data[:bh.Size()-1]
	delete(bh.indexes, ele)

	if err := heapify(0, bh.c, bh.isMaxHeap, bh.data, bh.indexes); err != nil {
		return nil, err
	}

	return ele, nil
}

func (bh *binaryHeap) Update(prev, new interface{}) error {
	if bh.IsEmpty() {
		return errors.New("heap is empty")
	}

	pt := utils.GetTypeName(prev)
	if pt != bh.typeURL {
		return liberror.NewTypeMismatchError(bh.typeURL, pt)
	}

	nt := utils.GetTypeName(new)
	if nt != bh.typeURL {
		return liberror.NewTypeMismatchError(bh.typeURL, nt)
	}

	idx, ok := bh.indexes[prev]
	if !ok {
		return fmt.Errorf("%v not found in heap", prev)
	}

	if prev == new {
		return fmt.Errorf("%v and %v are same", prev, new)
	}

	res, err := bh.c.Compare(prev, new)
	if err != nil {
		return err
	}

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

func (bh *binaryHeap) UpdateFunc(prev interface{}, op func(interface{}) interface{}) error {
	if bh.IsEmpty() {
		return errors.New("heap is empty")
	}

	pt := utils.GetTypeName(prev)
	if pt != bh.typeURL {
		return liberror.NewTypeMismatchError(bh.typeURL, pt)
	}

	idx, ok := bh.indexes[prev]
	if !ok {
		return fmt.Errorf("%v not found in heap", prev)
	}

	updated := op(prev)
	nt := utils.GetTypeName(updated)
	if nt != bh.typeURL {
		return liberror.NewTypeMismatchError(bh.typeURL, nt)
	}

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

func (bh *binaryHeap) Delete() error {
	if _, err := bh.Extract(); err != nil {
		return err
	}

	return nil
}

func (bh *binaryHeap) Size() int {
	return len(bh.data)
}

func (bh *binaryHeap) IsEmpty() bool {
	return bh.Size() == 0
}

func (bh *binaryHeap) Clear() {
	bh.data = nil
}

func (bh *binaryHeap) Iterator() iterator.Iterator {
	return newBinaryHeapIterator(bh)
}

type binaryHeapIterator struct {
	currentIndex int
	h            *binaryHeap
}

func newBinaryHeapIterator(bh *binaryHeap) *binaryHeapIterator {
	return &binaryHeapIterator{
		currentIndex: 0,
		h:            bh,
	}
}

func (bhi *binaryHeapIterator) HasNext() bool {
	return bhi.currentIndex != bhi.h.Size()
}

func (bhi *binaryHeapIterator) Next() interface{} {
	if bhi.currentIndex >= bhi.h.Size() {
		return nil
	}

	e := bhi.h.data[bhi.currentIndex]
	bhi.currentIndex++

	return e
}
