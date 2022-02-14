package list

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
)

type List[T comparable] interface {
	/*
		adds an element in the list
		if the list is empty and the type was never set then the
		elements type becomes the type of the list.

		params:
		e: element to add to list.

		returns:
		error: returns type mismatch error when the element to add is of different
		then what was set for list.
	*/
	Add(e T)

	/*
		adds an element at an specified index in the list
		the index has to be between 0 and (size of list - 1)
		if the list is empty and type was never set and the index is 0
		then the elements type becomes the type of the list.

		params:
		i: index at which element is supposed to be added.
		e: element to add to list.

		returns:
		error: returns indexOutOfBoundError if the index is < 0 or >= size or
		type mismatch error when the element to add is of different
		then what was set for list.
	*/
	AddAt(i int64, e T) error

	/*
		adds all the element specified in the list
		if the list is empty and type was never set then first elements type
		becomes the type of the list.

		params:
		l: list of elements to be added to list.

		returns:
		error: returns generic error if all element are not of same type or
		type mismatch error when the any of the element to add is of different
		then what was set for list.
	*/
	AddAll(l ...T)

	/*
		clears the content of list.
	*/
	Clear()

	/*
		return an new object that is the clone of current list or an error if the clone fails.
	*/
	Clone() List[T]

	/*
		use to check if the element is present in the list.

		params:
		e: the element to be searched in list.

		returns:
		bool: returns true if the element is present in the list.
		error: returns generic error if list is empty
		type mismatch error when the element to search is of different
		then what was set for list.
	*/
	Contains(e T) bool

	/*
		use to check if multiple elements are present in the list.

		params:
		l: the elements to be checked in list.

		returns:
		bool: returns true if all the elements are present in the list, else false.
		error: returns generic error if list is empty
		type mismatch error when the any of the element to search is of different
		then what was set for list.

	*/
	ContainsAll(l ...T) bool

	/*
		returns an element at a specified index.

		params:
		i: the index corresponding to which element has to be returned.

		returns:
		interface: returns the element at the specified index if the index is valid.
	*/
	Get(i int64) (T, error)

	Filter(predicate predicate.Predicate[T]) List[T]

	FindFirst(predicate predicate.Predicate[T]) (T, error)

	/*
		returns the index of an specified element.

		params:
		e: the element whose index is desired.

		returns:
		int: the index of the element if present else -1.
		error: returns generic error if the list is empty
		or type mismatch error if the element to be searched is of different type than what was
		set for the list, or generic error if the element is not found.

	*/
	IndexOf(e T) int64

	/*
		returns true if the list is empty else false.
	*/
	IsEmpty() bool

	/*
		returns an iterator for the list.
	*/
	Iterator() iterator.Iterator[T]

	/*
		returns a descending iterator for the list.
	*/
	DescendingIterator() iterator.Iterator[T]

	/*
		returns the last index for the element.

		params:
		e: the element whose last index is desired.

		returns:
		int: the index of the last element if present else -1.
		error: returns generic error if the list is empty
		or type mismatch error if the element to be searched is of different type than what was
		set for the list, or generic error if the element is not found.
	*/
	LastIndexOf(e T) int64

	/*
		use to remove and element from the list.

		params:
		e: the element to be removed.

		returns:
		bool: return true if the element was removed else false.
		error: generic error is the list is empty or,
		type mismatch error if the element to be removed is of different type than what was set
		for the list, or generic error if the element was not found in the list.
	*/
	Remove(e T) error

	/*
		use to remove and element at a specified index.

		params:
		i: the index corresponding to which element has to be removed.

		returns:
		bool: return true if the element was removed at the specified index else false.
		error: generic error is the list is empty or,
		indexOutOfBound error if the specified index is invalid.
	*/
	RemoveAt(i int64) (T, error)

	/*
		use to remove all the elements specified, if the element is not present in the list
		it is ignored.

		params:
		l: elements which are supposed to be removed.

		returns:
		bool: returns true if the elements were removed, else false.
		error: returns generic error if the list is empty, or
		type mismatch error if any element is of different type then what was set for the list.
	*/
	RemoveAll(l ...T) error

	/*
		replace and given value with another one.

		params:
		old: the value to be replaced.
		new: the value to replace with.
		error: return error if the list is empty or if the element is not found or if type mismatch.
	*/
	Replace(old, new T) error

	/*
		runs and function over all the element in the list, eg inc operator to increment all the
		elements of an integer list.

		params:
		uo: the function to be applied on every element of the list.
	*/
	ReplaceAll(uo operator.UnaryOperator[T])

	/*
		use to retain all the elements specified, other elements are removed for the list,
		if the element is not present in the list it is ignored.

		params:
		l: elements to retain.

		returns:
		bool: returns true if only element specified were retained else false.
		error: returns generic error if the list is empty, or
		type mismatch error if any element is of different type then what was set for the list.
	*/
	RetainAll(l ...T) error

	/*
		use to set an given element at a specified index.

		params:
		i: index at which the element has to be set.
		e: the element to set.

		returns:
		interface: returns the element if set was successful else nil.
		error: returns generic error if the list is empty or,
		return indexOutOfBound error if the specified index is invalid or,
		return type mismatch error if the elements type is different than one set for the list.
	*/
	Set(i int64, e T) (T, error)

	/*
		returns the count of number of elements in the list.
	*/
	Size() int64

	/*
		sorts the list based on the comparator specified.
	*/
	Sort(c comparator.Comparator[T])

	/*
		returns an sublist starting for index s to index e-1

		params:
		s: the start index for sublist.
		e: the end index for sublist, end index is 1 less than specified.

		returns:
		List: the sublist starting form s to e-1 from original list.
		error: returns generic error is end index is smaller than start or,
		returns indexOutOfBound error if either s or e is invalid index or,
		returns error if creation or adding to list fails
	*/
	SubList(s int64, e int64) (List[T], error)
}
