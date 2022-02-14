package list

import (
	"fmt"
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

type doubler struct{}

func (db doubler) Apply(e int64) int64 {
	return e * 2
}

type intToString struct{}

func (intToString) Apply(e any) any {
	return fmt.Sprintf("%d", e)
}

type evenFilter struct{}

func (ev evenFilter) Test(e int64) bool {
	return e%2 == 0
}
