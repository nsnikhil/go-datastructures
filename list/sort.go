package list

import "github.com/nsnikhil/go-datastructures/functions/comparator"

type sorter[T comparable] interface {

	/*
		sorts an list given a specified comparator.

		params:
		l: the list to sort.
		c: the comparator to use while sorting.
	*/
	sort(l List[T], c comparator.Comparator[T])
}

type quickSort[T comparable] struct{}

func newQuickSorter[T comparable]() sorter[T] {
	return quickSort[T]{}
}

func (qs quickSort[T]) sort(l List[T], c comparator.Comparator[T]) {
	quickSortUtil(l, c, 0, l.Size()-1)
}

func quickSortUtil[T comparable](l List[T], c comparator.Comparator[T], s, p int64) {
	if s < p {
		pivot := findPivot(l, c, s, p)
		quickSortUtil(l, c, s, pivot-1)
		quickSortUtil(l, c, pivot+1, p)
	}
}

func findPivot[T comparable](l List[T], c comparator.Comparator[T], s, p int64) int64 {
	i := s - 1
	j := s

	for j < p {
		je, _ := l.Get(j)
		pe, _ := l.Get(p)

		r := c.Compare(je, pe)
		if r < 0 {
			i++
			swap(l, i, j)
		}
		j++
	}

	swap(l, i+1, p)

	return i + 1
}

func swap[T comparable](l List[T], i, j int64) {
	tl, _ := l.Get(j)
	temp := tl

	el, _ := l.Get(i)
	_, _ = l.Set(j, el)
	_, _ = l.Set(i, temp)
}
