package set

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Set interface {
	Add(e interface{}) error

	AddAll(e ...interface{}) error

	Clear()

	Contains(e interface{}) bool

	ContainsAll(e ...interface{}) bool

	Copy() Set

	IsEmpty() bool

	Size() int

	Remove(e interface{}) error

	RemoveAll(e ...interface{}) error

	RetainAll(e ...interface{}) error

	Iterator() iterator.Iterator

	Union(s Set) (Set, error)

	Intersection(s Set) (Set, error)
}
