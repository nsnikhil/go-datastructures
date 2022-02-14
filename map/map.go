package gmap

import (
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
)

type Map[K comparable, V comparable] interface {
	Put(key K, value V) V

	PutAll(values ...*Pair[K, V])

	Get(key K) (V, error)

	GetOrDefault(key K, defaultValue V) V

	Remove(key K) (V, error)

	RemoveWithVal(key K, value V) (V, error)

	Replace(key K, newValue V) error

	ReplaceWithVal(key K, oldValue V, newValue V) error

	ReplaceAll(f function.BiFunction[K, V, V]) error

	Compute(key K, f function.BiFunction[K, V, V]) (V, error)

	ContainsKey(key K) bool

	ContainsValue(value V) bool

	Size() int64

	Keys() (list.List[K], error)

	Values() (list.List[V], error)

	Clear()

	IsEmpty() bool

	Iterator() iterator.Iterator[*Pair[K, V]]
}
