package list

import "github.com/nsnikhil/go-datastructures/functions/comparator"

type sorter interface {

	/*
		sorts an list given a specified comparator.

		params:
		l: the list to sort.
		c: the comparator to use while sorting.
	*/
	sort(l List, c comparator.Comparator)
}

func newSorter() sorter {
	return newQuickSorter()
}

type quickSort struct{}

func newQuickSorter() sorter {
	return quickSort{}
}

func (qs quickSort) sort(l List, c comparator.Comparator) {
	quickSortUtil(l, c, 0, l.Size()-1)
}

func quickSortUtil(l List, c comparator.Comparator, s, p int) {
	if s < p {
		pivot := findPivot(l, c, s, p)
		quickSortUtil(l, c, s, pivot-1)
		quickSortUtil(l, c, pivot+1, p)
	}
}

func findPivot(l List, c comparator.Comparator, s, p int) int {
	i := s - 1
	j := s

	for j < p {
		r, _ := c.Compare(l.Get(j), l.Get(p))
		if r < 0 {
			i++
			swap(l, i, j)
		}
		j++
	}

	swap(l, i+1, p)

	return i + 1
}

func swap(l List, i, j int) {
	temp := l.Get(j)
	_, _ = l.Set(j, l.Get(i))
	_, _ = l.Set(i, temp)
}
