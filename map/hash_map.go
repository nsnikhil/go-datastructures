package gmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/utils"
	"golang.org/x/crypto/sha3"
	"hash"
	"reflect"
)

const (
	notFoundError = "no value found against the key: %v"
	keys          = "keys"
	values        = "values"
)

type factors struct {
	upperLoadFactor float64
	lowerLoadFactor float64

	scalingFactor int
	capacity      int
}

type counter struct {
	elementCount int
	countMap     map[int]bool
	uniqueCount  int
}

type HashMap[T comparable] struct {
	h hash.Hash

	*factors
	*counter

	data *list.ArrayList[*list.LinkedList[T]]
}

func NewHashMap(values ...*Pair) (Map, error) {
	hm := newHashMap()

	if len(values) == 0 {
		return hm, nil
	}

	if err := insertAll(hm, values...); err != nil {
		return nil, err
	}

	return hm, nil
}

func (hm *HashMap) Put(k interface{}, v interface{}) (interface{}, error) {
	ov, err := hm.Get(k)
	if err != nil && errors.Is(err, &json.UnsupportedTypeError{}) {
		return nil, err
	}

	return ov, insertAll(hm, NewPair(k, v))
}

func (hm *HashMap) PutAll(values ...*Pair) error {
	return insertAll(hm, values...)
}

func (hm *HashMap) Get(k interface{}) (interface{}, error) {
	if hm.IsEmpty() {
		return nil, errors.New("map is empty")
	}

	kt := utils.GetTypeName(k)
	if hm.keyTypeURL != kt {
		return nil, liberr.TypeMismatchError(hm.keyTypeURL, kt)
	}

	cp, err := get(hm, k)
	if err != nil {
		return nil, err
	}

	return cp.v, nil
}

func (hm *HashMap) GetOrDefault(k, d interface{}) interface{} {
	v, err := hm.Get(k)
	if err != nil {
		return d
	}

	return v
}

func (hm *HashMap) Remove(k interface{}) (interface{}, error) {
	v, err := hm.Get(k)
	if err != nil {
		return nil, err
	}

	if err := remove(hm, k); err != nil {
		return nil, err
	}

	return v, nil
}

func (hm *HashMap) RemoveWithVal(k interface{}, v interface{}) (interface{}, error) {
	gv, err := hm.Get(k)
	if err != nil {
		return nil, err
	}

	if gv != v {
		return nil, fmt.Errorf("value mismatch: expected %v, got %v", v, gv)
	}

	if err := remove(hm, k); err != nil {
		return nil, err
	}

	return v, nil
}

func (hm *HashMap) Replace(k interface{}, nv interface{}) error {
	ov, err := hm.Get(k)
	if err != nil {
		return err
	}

	return hm.ReplaceWithVal(k, ov, nv)
}

func (hm *HashMap) ReplaceWithVal(k interface{}, ov interface{}, nv interface{}) error {
	cv, err := hm.Get(k)
	if err != nil {
		return err
	}

	if cv != ov {
		return fmt.Errorf("value mismatch: expected %v, got %v", cv, ov)
	}

	_, err = hm.Put(k, nv)
	return err
}

func (hm *HashMap) ReplaceAll(f function.BiFunction) error {
	it := hm.Iterator()

	for it.HasNext() {
		p := it.Next().(*Pair)
		nv := f.Apply(p.k, p.v)

		if hm.valueTypeURL != utils.GetTypeName(nv) {
			return liberr.TypeMismatchError(hm.valueTypeURL, utils.GetTypeName(nv))
		}

		p.v = nv
	}

	return nil
}

func (hm *HashMap) Compute(k interface{}, f function.BiFunction) (interface{}, error) {
	cp, err := get(hm, k)
	if err != nil {
		return nil, err
	}

	nv := f.Apply(k, cp.v)

	if hm.valueTypeURL != utils.GetTypeName(nv) {
		return nil, liberr.TypeMismatchError(hm.valueTypeURL, utils.GetTypeName(nv))
	}

	cp.v = nv

	return nv, nil
}

func (hm *HashMap) ContainsKey(k interface{}) bool {
	_, err := get(hm, k)
	return err == nil
}

func (hm *HashMap) ContainsValue(v interface{}) bool {
	it := hm.Iterator()

	for it.HasNext() {
		cp := it.Next().(*Pair)
		if cp.v == v {
			return true
		}
	}

	return false
}

func (hm *HashMap) Size() int {
	return hm.elementCount
}

func (hm *HashMap) Keys() (list.List, error) {
	return getAll(hm, keys)
}

func (hm *HashMap) Values() (list.List, error) {
	return getAll(hm, values)
}

func (hm *HashMap) Clear() {
	hm.h = sha3.New512()
	hm.capacity = initialCapacity
	hm.elementCount = nought
	hm.uniqueCount = nought
	hm.countMap = make(map[int]bool)
	hm.data = make([]*list.LinkedList, initialCapacity)
}

func (hm *HashMap) IsEmpty() bool {
	return hm.elementCount == nought
}

func (hm *HashMap) Iterator() iterator.Iterator {
	return newHashMapIterator(hm)
}

type hashMapIterator struct {
	currIndex       int
	currentIterator iterator.Iterator
	hm              *HashMap
}

func newHashMapIterator(hm *HashMap) iterator.Iterator {
	return &hashMapIterator{
		currIndex: nought,
		hm:        hm,
	}
}

func (hmi *hashMapIterator) HasNext() bool {
	for hmi.currIndex < hmi.hm.capacity {

		if !reflect.ValueOf(hmi.hm.data[hmi.currIndex]).IsNil() {
			if hmi.currentIterator == nil {
				hmi.currentIterator = hmi.hm.data[hmi.currIndex].Iterator()
			}

			if hmi.currentIterator.HasNext() {
				return true
			}
		}

		hmi.currentIterator = nil
		hmi.currIndex++
	}

	return hmi.currIndex < hmi.hm.capacity
}

func (hmi *hashMapIterator) Next() interface{} {
	return hmi.currentIterator.Next()
}

func insertAll(hm *HashMap, p ...*Pair) error {
	if len(p) == nought {
		return errors.New("argument list is empty")
	}

	kt, vt, err := validateType(p...)
	if err != nil {
		return err
	}

	if hm.keyTypeURL == utils.NA && hm.valueTypeURL == utils.NA {
		hm.keyTypeURL = kt
		hm.valueTypeURL = vt
	} else if hm.keyTypeURL != kt {
		return liberr.TypeMismatchError(hm.keyTypeURL, kt)
	} else if hm.valueTypeURL != vt {
		return liberr.TypeMismatchError(hm.valueTypeURL, vt)
	}

	for i := 0; i < len(p); i++ {
		if err := insert(hm, p[i]); err != nil {
			return err
		}
	}

	return nil
}

func insert(hm *HashMap, p *Pair) error {
	if hm.uniqueCount >= int(float64(hm.capacity)*hm.upperLoadFactor) {
		resizeUp(hm)
	}

	idx, err := indexOf(&hm.h, p.k, float64(hm.capacity))
	if err != nil {
		return err
	}

	if !hm.countMap[idx] {
		hm.countMap[idx] = true
		hm.uniqueCount++
	}

	ll := hm.data[idx]

	if ll == nil {
		tll, err := list.NewLinkedList()
		if err != nil {
			return err
		}

		hm.data[idx] = tll

		hm.elementCount++
		return tll.AddLast(p)
	}

	it := ll.Iterator()

	var curr *Pair

	for it.HasNext() {
		pr := it.Next().(*Pair)
		if pr.k == p.k {
			curr = pr
			break
		}
	}

	if curr == nil {
		hm.elementCount++
		return ll.AddLast(p)
	}

	curr.v = p.v

	return nil
}

func remove(hm *HashMap, k interface{}) error {
	idx, err := indexOf(&hm.h, k, float64(hm.capacity))
	if err != nil {
		return err
	}

	ll := hm.data[idx]
	if ll == nil {
		return fmt.Errorf(notFoundError, k)
	}

	it := ll.Iterator()

	var curr *Pair

	for it.HasNext() {
		pr := it.Next().(*Pair)
		if pr.k == k {
			curr = pr
			break
		}
	}

	if curr == nil {
		return fmt.Errorf(notFoundError, k)
	}

	if _, err := ll.Remove(curr); err != nil {
		return err
	}

	hm.elementCount--

	if ll.Size() == 0 {
		hm.data[idx] = (*list.LinkedList)(nil)
		delete(hm.countMap, idx)
		hm.uniqueCount--
	}

	if hm.capacity != initialCapacity && hm.uniqueCount <= int(float64(hm.capacity)*hm.lowerLoadFactor) {
		if err := resizeDown(hm); err != nil {
			return err
		}
	}

	return nil
}

func get(hm *HashMap, k interface{}) (*Pair, error) {
	idx, err := indexOf(&hm.h, k, float64(hm.capacity))
	if err != nil {
		return nil, err
	}

	ll := hm.data[idx]
	if ll == nil {
		return nil, fmt.Errorf(notFoundError, k)
	}

	it := ll.Iterator()

	for it.HasNext() {
		pr := it.Next().(*Pair)
		if pr.k == k {
			return pr, nil
		}
	}

	return nil, fmt.Errorf(notFoundError, k)
}

func getAll(hm *HashMap, data string) (list.List, error) {
	keys, err := list.NewLinkedList()
	if err != nil {
		return nil, err
	}

	it := hm.Iterator()
	for it.HasNext() {
		cp := it.Next().(*Pair)

		iv := cp.k
		if data == values {
			iv = cp.v
		}

		if err := keys.AddLast(iv); err != nil {
			return nil, err
		}
	}

	return keys, nil
}

func newHashMap() *HashMap {
	return &HashMap{
		typeURL: &typeURL{keyTypeURL: utils.NA, valueTypeURL: utils.NA},
		factors: &factors{upperLoadFactor: upperLoadFactor, lowerLoadFactor: lowerLoadFactor, scalingFactor: scalingFactor, capacity: initialCapacity},
		counter: &counter{elementCount: nought, countMap: make(map[int]bool), uniqueCount: nought},
		h:       sha3.New512(),
		data:    make([]*list.LinkedList, initialCapacity),
	}
}

func resizeUp(hm *HashMap) {
	hm.capacity *= hm.scalingFactor

	temp := make([]*list.LinkedList, hm.capacity)

	sz := len(hm.data)
	for i := 0; i < sz; i++ {
		temp[i] = hm.data[i]
	}

	hm.data = temp
}

func resizeDown(hm *HashMap) error {
	var data []*Pair

	it := hm.Iterator()
	for it.HasNext() {
		e := it.Next().(*Pair)
		data = append(data, e)
	}

	newCap := hm.capacity / hm.scalingFactor

	hm.Clear()

	hm.capacity = newCap

	err := insertAll(hm, data...)

	return err
}

func validateType(values ...*Pair) (string, string, error) {
	kt, vt := utils.GetTypeName(values[0].k), utils.GetTypeName(values[0].v)
	sz := len(values)

	for i := 1; i < sz; i++ {
		if utils.GetTypeName(values[i].k) != kt {
			return utils.NA, utils.NA, liberr.TypeMismatchError(kt, utils.GetTypeName(values[i].k))
		}

		if utils.GetTypeName(values[i].v) != vt {
			return utils.NA, utils.NA, liberr.TypeMismatchError(vt, utils.GetTypeName(values[i].v))
		}
	}

	return kt, vt, nil
}
