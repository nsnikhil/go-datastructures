package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/liberror"
)

const (
	linear     = 0
	concurrent = 1
	doubly     = 2
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

func newFinder(t int) finder {
	switch t {
	case concurrent:
		return newConcurrentFinder()
	case doubly:
		return newDoublyFinder()
	default:
		return newLinearFinder()
	}
}

type linearFinder struct{}

func newLinearFinder() finder {
	return linearFinder{}
}

/*
	searches the list using the iterator for the element.
*/
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

func searchUtil(l List, e interface{}, start, end int, res chan<- int) {
	for start <= end {
		if l.Get(start) == e {
			res <- start
		}
		if l.Get(end) == e {
			res <- end
		}
		start++
		end--
	}
	res <- -1
}

type concurrentFinder struct {
	searchFactor  int
	partitionSize int
}

func newConcurrentFinder() finder {
	return concurrentFinder{
		searchFactor:  10000,
		partitionSize: 1000,
	}
}

/*
	concurrent finder splits the list into partitions and searches in each partition in a new go routine.

	concurrent finder is only invoked if the number of elements in list are greater than search factor else
	the linear finder is invoked

	below are the benchmark comparing linear and concurrent finder:

	linear finder:
		element count      iterations      iterations/sec
		10         	        3288336	        368 ns/op
		100         	    856701	        1376 ns/op
		1000         	    108962	        10903 ns/op
		10000         	    10000	        107407 ns/op
		100000         	    1130	        1042636 ns/op
		1000000         	100	            10808345 ns/op
		10000000         	10	            110666584 ns/op
		100000000         	1	            1074130085 ns/op

	concurrent finder:
		element count      iterations      iterations/sec
		10         	        2913409	        415 ns/op
		100         	    881508	        1337 ns/op
		1000         	    113041	        10654 ns/op
		10000         	    26298	        44269 ns/op
		100000         	    3853	        303841 ns/op
		1000000        	    435	            2731591 ns/op
		10000000         	40	            26074049 ns/op
		100000000         	4	            280420052 ns/op

*/
func (cf concurrentFinder) search(l List, e interface{}) (int, error) {
	sz := l.Size()
	if sz == 0 {
		return -1, fmt.Errorf("list is empty")
	}

	if getTypeName(l.Get(0)) != getTypeName(e) {
		return -1, liberror.NewTypeMismatchError(getTypeName(l.Get(0)), getTypeName(e))
	}

	if sz < cf.searchFactor {
		return newLinearFinder().search(l, e)
	}

	res := make(chan int, cf.partitionSize)

	for i := 0; i < sz; i += cf.partitionSize {
		go searchUtil(l, e, i, i+cf.partitionSize, res)
	}

	for i := 0; i < sz; i += cf.partitionSize {
		if idx := <-res; idx != -1 {
			return idx, nil
		}
	}

	return -1, fmt.Errorf("element %v not found in the list", e)
}

//should only be used by list with descending iterator like linked list here
type doublyFinder struct{}

func newDoublyFinder() doublyFinder {
	return doublyFinder{}
}

/*
	searches the list form both end simultaneously for the element.
*/
func (df doublyFinder) search(l List, e interface{}) (int, error) {
	sz := l.Size()
	if sz == 0 {
		return -1, fmt.Errorf("list is empty")
	}

	if getTypeName(l.Get(0)) != getTypeName(e) {
		return -1, liberror.NewTypeMismatchError(getTypeName(l.Get(0)), getTypeName(e))
	}

	it := l.Iterator()

	dit := l.(*LinkedList).DescendingIterator()

	forwardCount := 0
	backwardCount := 0

	for it.HasNext() && dit.HasNext() {

		if it.Next() == e {
			return forwardCount, nil
		}

		if dit.Next() == e {
			return l.Size() - backwardCount - 1, nil
		}

		forwardCount++
		backwardCount++
	}

	return -1, fmt.Errorf("element %v not found in the list", e)
}
