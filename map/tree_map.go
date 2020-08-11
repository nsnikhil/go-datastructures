package gmap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/tree"
	"hash"
)

type TreeMap struct {
	h hash.Hash

	*typeURL
	*factors
	*counter

	data tree.Tree
}

func NewTreeMap(c comparator.Comparator, values ...*Pair) (Map, error) {
	bst, err := tree.NewBinarySearchTree(c)
	if err != nil {
		return nil, err
	}

	tm := &TreeMap{
		data: bst,
	}

	return tm, nil
}

func (tm *TreeMap) Put(k interface{}, v interface{}) (interface{}, error) {
	return nil, nil
}

func (tm *TreeMap) PutAll(values ...*Pair) error {
	return nil
}

func (tm *TreeMap) Get(k interface{}) (interface{}, error) {
	return nil, nil
}

func (tm *TreeMap) GetOrDefault(k, d interface{}) interface{} {
	return nil
}

func (tm *TreeMap) Remove(k interface{}) (interface{}, error) {
	return nil, nil
}

func (tm *TreeMap) RemoveWithVal(k interface{}, v interface{}) (interface{}, error) {
	return nil, nil
}

func (tm *TreeMap) Replace(k interface{}, nv interface{}) error {
	return nil
}

func (tm *TreeMap) ReplaceWithVal(k interface{}, ov interface{}, nv interface{}) error {
	return nil
}

func (tm *TreeMap) ReplaceAll(f function.BiFunction) error {
	return nil
}

func (tm *TreeMap) Compute(k interface{}, f function.BiFunction) (interface{}, error) {
	return nil, nil
}

func (tm *TreeMap) ContainsKey(k interface{}) bool {
	return false
}

func (tm *TreeMap) ContainsValue(v interface{}) bool {
	return false
}

func (tm *TreeMap) Size() int {
	return 0
}

func (tm *TreeMap) Keys() (list.List, error) {
	return nil, nil
}

func (tm *TreeMap) Values() (list.List, error) {
	return nil, nil
}

func (tm *TreeMap) Clear() {

}

func (tm *TreeMap) IsEmpty() bool {
	return false
}

func (tm *TreeMap) Iterator() iterator.Iterator {
	return nil
}
