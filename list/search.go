package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/liberror"
)

type finder interface {

	/*
		returns the index of specified in the list.

		params:
		List: the list where elements has to be searched in.
		e: the element to search.

		returns:
		int: the index of the given element else -1.
		error: returns generic error is the list is empty or,
		returns type mismatch error if the type of the element to search is different then the type
		set for the list.
	*/
	search(l List, e interface{}) (int, error)
}

func newFinder() finder {
	return newLinearFinder()
}

type linearFinder struct{}

func newLinearFinder() finder {
	return linearFinder{}
}

func (lf linearFinder) search(l List, e interface{}) (int, error) {
	sz := l.Size()
	if sz == 0 {
		return -1, fmt.Errorf("list is empty")
	}

	if getTypeName(l.Get(0)) != getTypeName(e) {
		return -1, liberror.NewTypeMismatchError(getTypeName(l.Get(0)), getTypeName(e))
	}

	it := l.Iterator()

	count := 0
	for it.HasNext() {
		if it.Next() == e {
			return count, nil
		}
		count++
	}

	return -1, fmt.Errorf("element %v not found in the list", e)

}
