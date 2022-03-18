package heap

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Heap[T comparable] interface {
	/*

	 */
	Add(e ...T)

	/*

	 */
	Extract() (T, error)

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
