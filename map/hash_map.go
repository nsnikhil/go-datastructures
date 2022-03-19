package gmap

import (
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"golang.org/x/crypto/sha3"
	"hash"
	"reflect"
)

type factors struct {
	upperLoadFactor float64
	lowerLoadFactor float64

	scalingFactor int64
	capacity      int64
}

type counter struct {
	elementCount int64
	countMap     map[int64]bool
	uniqueCount  int64
}

type HashMap[K comparable, V comparable] struct {
	h hash.Hash

	*factors
	*counter

	data []*list.LinkedList[*Pair[K, V]]
}

func NewHashMap[K comparable, V comparable](values ...*Pair[K, V]) Map[K, V] {
	hm := newHashMap[K, V]()

	hm.insertAll(values...)

	return hm
}

func (hm *HashMap[K, V]) Put(key K, value V) V {
	ov, _ := hm.Get(key)

	hm.insertAll(NewPair[K, V](key, value))

	return ov
}

//TODO: SHOULD IT RETURN ERROR WHEN ARGS LIST IS EMPTY?
func (hm *HashMap[K, V]) PutAll(values ...*Pair[K, V]) {
	hm.insertAll(values...)
}

func (hm *HashMap[K, V]) Get(key K) (V, error) {
	if hm.IsEmpty() {
		return internal.ZeroValueOf[V](), emptyMapError("HashMap.Get")
	}

	cv, err := hm.get(key)
	if err != nil {
		return internal.ZeroValueOf[V](), err
	}

	return cv.second, nil
}

func (hm *HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	cv, err := hm.get(key)
	if err != nil {
		return defaultValue
	}

	return cv.second
}

func (hm *HashMap[K, V]) Remove(key K) (V, error) {
	cv, err := hm.get(key)
	if err != nil {
		return internal.ZeroValueOf[V](), err
	}

	if err := hm.remove(key); err != nil {
		return internal.ZeroValueOf[V](), err
	}

	return cv.second, nil
}

func (hm *HashMap[K, V]) RemoveWithVal(key K, value V) (V, error) {
	gv, err := hm.get(key)
	if err != nil {
		return internal.ZeroValueOf[V](), err
	}

	if gv.second != value {
		return internal.ZeroValueOf[V](), valueMisMatchError(value, gv.second, "HashMap.RemoveWithVal")
	}

	if err := hm.remove(key); err != nil {
		return internal.ZeroValueOf[V](), err
	}

	return value, nil
}

func (hm *HashMap[K, V]) Replace(key K, newValue V) error {
	ov, err := hm.Get(key)
	if err != nil {
		return err
	}

	return hm.ReplaceWithVal(key, ov, newValue)
}

func (hm *HashMap[K, V]) ReplaceWithVal(key K, oldValue V, newValue V) error {
	cv, err := hm.Get(key)
	if err != nil {
		return err
	}

	if cv != oldValue {
		return valueMisMatchError(cv, oldValue, "HashMap.RemoveWithVal")
	}

	hm.Put(key, newValue)

	return nil
}

func (hm *HashMap[K, V]) ReplaceAll(f function.BiFunction[K, V, V]) error {
	it := hm.Iterator()

	for it.HasNext() {
		p, _ := it.Next()

		nv := f.Apply(p.first, p.second)

		p.second = nv
	}

	return nil
}

func (hm *HashMap[K, V]) Compute(key K, f function.BiFunction[K, V, V]) (V, error) {
	cp, err := hm.get(key)
	if err != nil {
		return internal.ZeroValueOf[V](), err
	}

	nv := f.Apply(key, cp.second)

	cp.second = nv

	return nv, nil
}

func (hm *HashMap[K, V]) ContainsKey(key K) bool {
	_, err := hm.get(key)
	return err == nil
}

func (hm *HashMap[K, V]) ContainsValue(value V) bool {
	it := hm.Iterator()

	for it.HasNext() {
		cp, _ := it.Next()
		if cp.second == value {
			return true
		}
	}

	return false
}

func (hm *HashMap[K, V]) Size() int64 {
	return hm.elementCount
}

func (hm *HashMap[K, V]) Keys() (list.List[K], error) {
	return getKeys[K, V](hm)
}

func (hm *HashMap[K, V]) Values() (list.List[V], error) {
	return getValues[K, V](hm)
}

func (hm *HashMap[K, V]) Clear() {
	hm.h = sha3.New512()
	hm.capacity = initialCapacity
	hm.elementCount = internal.Zero
	hm.uniqueCount = internal.Zero
	hm.countMap = make(map[int64]bool)
	hm.data = make([]*list.LinkedList[*Pair[K, V]], initialCapacity)
}

func (hm *HashMap[K, V]) IsEmpty() bool {
	return hm.elementCount == internal.Zero
}

func (hm *HashMap[K, V]) Iterator() iterator.Iterator[*Pair[K, V]] {
	return newHashMapIterator(hm)
}

type hashMapIterator[K comparable, V comparable] struct {
	currIndex       int64
	currentIterator iterator.Iterator[*Pair[K, V]]
	hm              *HashMap[K, V]
}

func newHashMapIterator[K comparable, V comparable](hm *HashMap[K, V]) iterator.Iterator[*Pair[K, V]] {
	return &hashMapIterator[K, V]{
		currIndex: internal.Zero,
		hm:        hm,
	}
}

func (hmi *hashMapIterator[K, V]) HasNext() bool {
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

func (hmi *hashMapIterator[K, V]) Next() (*Pair[K, V], error) {
	return hmi.currentIterator.Next()
}

func (hm *HashMap[K, V]) insertAll(p ...*Pair[K, V]) error {
	if len(p) == internal.Zero {
		return nil
	}

	for i := 0; i < len(p); i++ {
		if err := hm.insert(p[i]); err != nil {
			return err
		}
	}

	return nil
}

func (hm *HashMap[K, V]) insert(p *Pair[K, V]) error {
	if hm.uniqueCount >= int64(float64(hm.capacity)*hm.upperLoadFactor) {
		hm.resizeUp()
	}

	idx, err := indexOf(&hm.h, p.first, hm.capacity)
	if err != nil {
		return err
	}

	if !hm.countMap[idx] {
		hm.countMap[idx] = true
		hm.uniqueCount++
	}

	ll := hm.data[idx]

	if ll == nil {
		tll := list.NewLinkedList[*Pair[K, V]]()

		hm.data[idx] = tll

		hm.elementCount++

		tll.AddLast(p)
		return nil
	}

	it := ll.Iterator()

	var curr *Pair[K, V]

	for it.HasNext() {
		pr, _ := it.Next()
		if pr.first == p.first {
			curr = pr
			break
		}
	}

	if curr == nil {
		hm.elementCount++
		ll.AddLast(p)
		return nil
	}

	curr.second = p.second

	return nil
}

func (hm *HashMap[K, V]) remove(key K) error {
	idx, err := indexOf(&hm.h, key, hm.capacity)
	if err != nil {
		return err
	}

	ll := hm.data[idx]
	if ll == nil {
		return keyNotFoundError(key, "HashMap.remove")
	}

	it := ll.Iterator()

	var curr *Pair[K, V]

	for it.HasNext() {
		pr, _ := it.Next()
		if pr.first == key {
			curr = pr
			break
		}
	}

	if curr == nil {
		return keyNotFoundError(key, "HashMap.remove")
	}

	if err := ll.Remove(curr); err != nil {
		return err
	}

	hm.elementCount--

	if ll.Size() == 0 {
		hm.data[idx] = (*list.LinkedList[*Pair[K, V]])(nil)
		delete(hm.countMap, idx)
		hm.uniqueCount--
	}

	if hm.capacity != initialCapacity && hm.uniqueCount <= int64(float64(hm.capacity)*hm.lowerLoadFactor) {
		hm.resizeDown()
	}

	return nil
}

func (hm *HashMap[K, V]) get(key K) (*Pair[K, V], error) {
	idx, err := indexOf(&hm.h, key, hm.capacity)
	if err != nil {
		return nil, err
	}

	ll := hm.data[idx]
	if ll == nil {
		return nil, keyNotFoundError(key, "HashMap.get")
	}

	it := ll.Iterator()

	for it.HasNext() {
		pr, _ := it.Next()
		if pr.first == key {
			return pr, nil
		}
	}

	return nil, keyNotFoundError(key, "HashMap.get")
}

//TODO: MERGE WITH GETVALUES
func getKeys[K comparable, V comparable](hm *HashMap[K, V]) (list.List[K], error) {
	keys := list.NewArrayList[K]()

	it := hm.Iterator()

	for it.HasNext() {
		cp, _ := it.Next()
		keys.Add(cp.first)
	}

	return keys, nil
}

//TODO: MERGE WITH GETKEYS
func getValues[K comparable, V comparable](hm *HashMap[K, V]) (list.List[V], error) {
	values := list.NewArrayList[V]()

	it := hm.Iterator()

	for it.HasNext() {
		cp, _ := it.Next()
		values.Add(cp.second)
	}

	return values, nil
}

func (hm *HashMap[K, V]) resizeUp() {
	resize(hm, hm.capacity*hm.scalingFactor)
}

func (hm *HashMap[K, V]) resizeDown() {
	resize(hm, hm.capacity/hm.scalingFactor)
}

func resize[K comparable, V comparable](hm *HashMap[K, V], newCap int64) {
	var data []*Pair[K, V]

	it := hm.Iterator()

	for it.HasNext() {
		e, _ := it.Next()
		data = append(data, e)
	}

	hm.Clear()

	hm.capacity = newCap

	hm.data = make([]*list.LinkedList[*Pair[K, V]], hm.capacity)

	hm.insertAll(data...)
}

func newHashMap[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{
		factors: &factors{upperLoadFactor: upperLoadFactor, lowerLoadFactor: lowerLoadFactor, scalingFactor: scalingFactor, capacity: initialCapacity},
		counter: &counter{elementCount: internal.Zero, countMap: make(map[int64]bool), uniqueCount: internal.Zero},
		h:       sha3.New512(),
		data:    make([]*list.LinkedList[*Pair[K, V]], initialCapacity),
	}
}
