package heap

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Heap interface {
	/*

	 */
	Add(e ...interface{}) error

	/*

	 */
	Extract() (interface{}, error)

	/*

	 */
	Clear()

	/*

	 */
	Iterator() iterator.Iterator

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
