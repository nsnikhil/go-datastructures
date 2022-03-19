package list

import "github.com/nsnikhil/go-datastructures/internal"

//should only be used by list with descending iterator like linked list here
type doublyFinder[T comparable] struct{}

func newDoublyFinder[T comparable]() doublyFinder[T] {
	return doublyFinder[T]{}
}

/*
	searches the list form both end simultaneously for the element.
*/
func (df doublyFinder[T]) search(l List[T], e T) int64 {
	sz := l.Size()
	if sz == 0 {
		return internal.InvalidIndex
	}

	it := l.Iterator()

	dit := l.DescendingIterator()

	forwardCount := int64(0)
	backwardCount := int64(0)

	for it.HasNext() && dit.HasNext() {

		v, _ := it.Next()
		if v == e {
			return forwardCount
		}

		nv, _ := dit.Next()
		if nv == e {
			return l.Size() - backwardCount - 1
		}

		forwardCount++
		backwardCount++
	}

	return internal.InvalidIndex
}
