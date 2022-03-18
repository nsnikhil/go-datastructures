package set

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Set[T comparable] interface {
	Add(e T)

	AddAll(e ...T)

	Clear()

	Contains(e T) bool

	ContainsAll(e ...T) bool

	Copy() Set[T]

	IsEmpty() bool

	Size() int64

	Remove(e T) error

	RemoveAll(e ...T) error

	RetainAll(e ...T) error

	Iterator() iterator.Iterator[T]

	Union(s Set[T]) (Set[T], error)

	Intersection(s Set[T]) (Set[T], error)
}
