package list

import (
	"github.com/nsnikhil/go-datastructures/internal"
)

func toSlice[T comparable](list List[T]) []T {
	res := make([]T, 0)
	it := list.Iterator()

	for it.HasNext() {
		v, _ := it.Next()
		res = append(res, v)
	}

	return res
}

func newTestArrayList(sz int64) *ArrayList[int64] {
	return NewArrayList(internal.SliceGenerator{Size: sz}.Generate()...)
}

func newTestLinkedList(sz int64) *LinkedList[int64] {
	return NewLinkedList(internal.SliceGenerator{Size: sz}.Generate()...)
}
