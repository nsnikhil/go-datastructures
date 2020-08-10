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
	Update(prev, new interface{}) error

	/*

	 */
	UpdateFunc(prev interface{}, op func(interface{}) interface{}) error

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
