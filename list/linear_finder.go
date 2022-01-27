package list

type linearFinder[T comparable] struct{}

func newLinearFinder[T comparable]() finder[T] {
	return linearFinder[T]{}
}

/*
	searches the list using the iterator for the element.
*/
func (lf linearFinder[T]) search(l List[T], e T) int64 {
	sz := l.Size()
	if sz == 0 {
		return -1
	}

	it := l.Iterator()

	count := int64(0)
	for it.HasNext() {
		v, _ := it.Next()
		if v == e {
			return count
		}
		count++
	}

	return -1

}
