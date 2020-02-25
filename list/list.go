package list

import (
	"datastructures/functions/comparator"
	"datastructures/functions/iterator"
	"datastructures/functions/operator"
)

type List interface {
	Add(e interface{}) error
	AddAt(i int, e interface{}) error
	AddAll(l ...interface{}) error

	Clear()

	Contains(e interface{}) (bool, error)
	ContainsAll(l ...interface{}) (bool, error)

	Get(i int) interface{}
	IndexOf(e interface{}) (int, error)

	IsEmpty() bool

	Iterator() iterator.Iterator

	LastIndexOf(e interface{}) (int, error)

	Remove(e interface{}) (bool, error)
	RemoveAt(i int) (interface{}, error)
	RemoveAll(l ...interface{}) (bool, error)

	ReplaceAll(uo operator.UnaryOperator) error
	RetainAll(l ...interface{}) (bool, error)

	Set(i int, e interface{}) (interface{}, error)
	Size() int

	Sort(c comparator.Comparator)

	SubList(s int, e int) (List, error)
}
