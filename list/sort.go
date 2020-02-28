package list

import "datastructures/functions/comparator"

type sorter interface {
	sort(l List, c comparator.Comparator)
}

func newSorter() sorter {
	return newQuickSorter()
}

type mergeSort struct{}

func newMergeSorter() sorter {
	return mergeSort{}
}

func (ms mergeSort) sort(l List, c comparator.Comparator) {
	mergeSortUtil(l, c, 0, l.Size()-1)
}

func merge(l List, c comparator.Comparator, s, m, e int) {
	i := s
	j := m + 1

	temp, _ := NewArrayList()

	for i <= m && j <= e {
		res, _ := c.Compare(l.Get(i), l.Get(j))
		if res < 0 {
			_ = temp.Add(l.Get(i))
			i++
		} else {
			_ = temp.Add(l.Get(j))
			j++
		}
	}

	for i <= m {
		_ = temp.Add(l.Get(i))
		i++
	}

	for j <= e {
		_ = temp.Add(l.Get(j))
		j++
	}

	it := temp.Iterator()

	k := s
	for it.HasNext() && k <= e {
		_, _ = l.Set(k, it.Next())
		k++
	}

}

func mergeSortUtil(l List, c comparator.Comparator, s, e int) {
	if s < e {
		mid := (s + e) / 2
		mergeSortUtil(l, c, s, mid)
		mergeSortUtil(l, c, mid+1, e)
		merge(l, c, s, mid, e)
	}
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
