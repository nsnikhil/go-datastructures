package gmap

//
//import (
//	"github.com/nsnikhil/go-datastructures/functions/comparator"
//	"github.com/nsnikhil/go-datastructures/functions/function"
//	"github.com/nsnikhil/go-datastructures/functions/iterator"
//	"github.com/nsnikhil/go-datastructures/list"
//	"hash"
//)
//
//type TreeMap[K comparable, V comparable] struct {
//	h hash.Hash
//
//	*factors
//	*counter
//
//	//data tree.Tree
//}
//
//func NewTreeMap[K comparable, V comparable](c comparator.Comparator[K], values ...*Pair[K, V]) Map[K, V] {
//	return nil
//	//bst, err := tree.NewBinarySearchTree(c)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//tm := &TreeMap{
//	//	data: bst,
//	//}
//	//
//	//return tm, nil
//}
//
//func (tm *TreeMap[K, V]) Put(k interface{}, v interface{}) (interface{}, error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) PutAll(values ...*Pair[K, V]) error {
//	return nil
//}
//
//func (tm *TreeMap[K, V]) Get(k interface{}) (interface{}, error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) GetOrDefault(k, d interface{}) interface{} {
//	return nil
//}
//
//func (tm *TreeMap[K, V]) Remove(k interface{}) (interface{}, error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) RemoveWithVal(k interface{}, v interface{}) (interface{}, error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) Replace(k interface{}, nv interface{}) error {
//	return nil
//}
//
//func (tm *TreeMap[K, V]) ReplaceWithVal(k interface{}, ov interface{}, nv interface{}) error {
//	return nil
//}
//
//func (tm *TreeMap[K, V]) ReplaceAll(f function.BiFunction) error {
//	return nil
//}
//
//func (tm *TreeMap[K, V]) Compute(k interface{}, f function.BiFunction) (interface{}, error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) ContainsKey(k interface{}) bool {
//	return false
//}
//
//func (tm *TreeMap[K, V]) ContainsValue(v interface{}) bool {
//	return false
//}
//
//func (tm *TreeMap[K, V]) Size() int {
//	return 0
//}
//
//func (tm *TreeMap[K, V]) Keys() (list.List[K], error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) Values() (list.List[V], error) {
//	return nil, nil
//}
//
//func (tm *TreeMap[K, V]) Clear() {
//
//}
//
//func (tm *TreeMap[K, V]) IsEmpty() bool {
//	return false
//}
//
//func (tm *TreeMap[K, V]) Iterator() iterator.Iterator[Pair[K, V]] {
//	return nil
//}
