package list

type reverseFinder[T comparable] struct{}

func newReverseFinder[T comparable]() finder[T] {
	return reverseFinder[T]{}
}

/*
	searches the list using the iterator for the element.
*/
func (lf reverseFinder[T]) search(l List[T], e T) int64 {
	sz := l.Size()
	if sz == 0 {
		return -1
	}

	it := l.DescendingIterator()

	count := l.Size() - 1

	for it.HasNext() {
		v, _ := it.Next()
		if v == e {
			return count
		}
		count--
	}

	return -1

}
