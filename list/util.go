package list

import (
	"datastructures/functions/comparator"
)

const (
	searchFactor = 5
)

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

func mergeSort(l List, c comparator.Comparator, s, e int) {
	if s < e {
		mid := (s + e) / 2
		mergeSort(l, c, s, mid)
		mergeSort(l, c, mid+1, e)
		merge(l, c, s, mid, e)
	}
}

func sortList(l List, c comparator.Comparator) {
	mergeSort(l, c, 0, l.Size()-1)
}

func searchUtil(start, end int, element interface{}, l List, res chan<- int) {
	i := start
	for i <= end {
		if l.Get(i) == element {
			res <- i
			return
		}
		i++
	}
}

//TODO
func concurrentSearch(l List, e interface{}) int {
	sz := l.Size()

	if sz <= searchFactor {
		return linearSearch(l, e)
	}

	parts := sz / searchFactor

	res := make(chan int)

	for i := 0; i < sz; i = i + parts {
		go searchUtil(i, i+parts, e, l, res)
	}

	select {
	case r := <-res:
		return r
	}

	//TODO FIX THIS

}

func searchList(l List, e interface{}) int {
	return linearSearch(l, e)
}

func linearSearch(l List, e interface{}) int {
	it := l.Iterator()
	i := 0

	for it.HasNext() {
		if it.Next() == e {
			return i
		}
		i++
	}

	return -1
}
