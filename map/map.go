package gmap

import (
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
)

type Map interface {
	Put(k interface{}, v interface{}) (interface{}, error)

	PutAll(values ...*Pair) error

	Get(k interface{}) (interface{}, error)

	GetOrDefault(k, d interface{}) interface{}

	Remove(k interface{}) (interface{}, error)

	RemoveWithVal(k interface{}, v interface{}) (interface{}, error)

	Replace(k interface{}, nv interface{}) error

	ReplaceWithVal(k interface{}, ov interface{}, nv interface{}) error

	ReplaceAll(f function.BiFunction) error

	Compute(k interface{}, f function.BiFunction) (interface{}, error)

	ContainsKey(k interface{}) bool

	ContainsValue(v interface{}) bool

	Size() int

	Keys() (list.List, error)

	Values() (list.List, error)

	Clear()

	IsEmpty() bool

	Iterator() iterator.Iterator
}
