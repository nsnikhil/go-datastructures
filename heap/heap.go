package heap

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Heap[T comparable] interface {
	/*

	 */
	Add(e ...T) error

	/*

	 */
	Extract() (T, error)

	/*

	 */
	Update(prev, new T) error

	/*

	 */
	UpdateFunc(prev T, op func(T) T) error

	/*

	 */
	Clear()

	/*

	 */
	Iterator() iterator.Iterator[T]

	/*

	 */
	Delete() error

	/*

	 */
	Size() int

	/*

	 */
	IsEmpty() bool
}
