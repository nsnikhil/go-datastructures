package list

type finder[T comparable] interface {

	/*
		returns the index of specified in the list.

		params:
		List: the list where elements has to be searched in.
		e: the element to search.

		returns:
		int: the index of the given element else -1.
	*/
	search(l List[T], e T) int64
}
