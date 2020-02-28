package list

import (
	"fmt"
	"reflect"
)

type finder interface {
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

	if reflect.TypeOf(l.Get(0)).Name() != reflect.TypeOf(e).Name() {
		return -1, fmt.Errorf("type mismatch : expected %s got %s", reflect.TypeOf(l.Get(0)).Name(), reflect.TypeOf(e).Name())
	}

	res := make(chan int)

	go searchUtil(l, e, 0, sz-1, res)

	if idx := <-res; idx != -1 {
		return idx, nil
	}

	return -1, fmt.Errorf("element %v not found in the list", e)

}

func searchUtil(l List, e interface{}, start, end int, res chan<- int) {
	//fmt.Println(fmt.Sprintf("%d - %d", start, end))
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
		searchFactor:  1000,
		partitionSize: 1000,
	}
}

func (cf concurrentFinder) search(l List, e interface{}) (int, error) {
	sz := l.Size()
	if sz == 0 {
		return -1, fmt.Errorf("list is empty")
	}

	if reflect.TypeOf(l.Get(0)).Name() != reflect.TypeOf(e).Name() {
		return -1, fmt.Errorf("type mismatch : expected %s got %s", reflect.TypeOf(l.Get(0)).Name(), reflect.TypeOf(e).Name())
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
