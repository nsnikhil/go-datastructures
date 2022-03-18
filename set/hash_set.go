package set

import (
	"errors"
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/internal"
	gmap "github.com/nsnikhil/go-datastructures/map"
)

var errorDifferentTypes = errors.New("all elements must be of same type")

type present struct{}

type HashSet[T comparable] struct {
	data gmap.Map[T, present]
}

func NewHashSet[T comparable](e ...T) *HashSet[T] {
	hm := gmap.NewHashMap[T, present]()

	hs := newHashSet[T](hm)

	if len(e) == 0 {
		return hs
	}

	hs.insert(e...)

	return hs
}

func (hs *HashSet[T]) Add(e T) {
	hs.insert(e)
}

//TODO: SHOULD IT THROW ERRO IF ARGS ARE EMPTY?
func (hs *HashSet[T]) AddAll(e ...T) {
	hs.insert(e...)
}

func (hs *HashSet[T]) Clear() {
	hs.data.Clear()
}

func (hs *HashSet[T]) Contains(e T) bool {
	return hs.contains(e)
}

func (hs *HashSet[T]) ContainsAll(e ...T) bool {
	return hs.contains(e...)
}

func (hs *HashSet[T]) Copy() Set[T] {
	dt := make([]T, 0)

	it := hs.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		dt = append(dt, v)
	}

	return NewHashSet(dt...)
}

func (hs *HashSet[T]) IsEmpty() bool {
	return hs.data.IsEmpty()
}

func (hs *HashSet[T]) Size() int64 {
	return hs.data.Size()
}

func (hs *HashSet[T]) Remove(e T) error {
	return hs.remove(false, e)
}

func (hs *HashSet[T]) RemoveAll(e ...T) error {
	return hs.remove(true, e...)
}

func (hs *HashSet[T]) RetainAll(e ...T) error {
	if hs.IsEmpty() {
		return errors.New("set is empty")
	}

	tm := make(map[T]bool)

	for _, k := range e {
		tm[k] = true
	}

	dl := make([]T, 0)

	it := hs.Iterator()
	for it.HasNext() {
		e, _ := it.Next()

		if !tm[e] {
			dl = append(dl, e)
		}
	}

	if len(dl) == 0 {
		return nil
	}

	return hs.remove(false, dl...)
}

func (hs *HashSet[T]) Iterator() iterator.Iterator[T] {
	return newHashIterator[T](hs.data.Iterator())
}

func (hs *HashSet[T]) Union(s Set[T]) (Set[T], error) {
	ns := NewHashSet[T]()

	if err := ns.union(hs); err != nil {
		return nil, err
	}

	if err := ns.union(s.(*HashSet[T])); err != nil {
		return nil, err
	}

	return ns, nil
}

func (hs *HashSet[T]) Intersection(s Set[T]) (Set[T], error) {
	ns := NewHashSet[T]()

	tm := make(map[T]bool)
	it := hs.Iterator()
	for it.HasNext() {
		e, _ := it.Next()

		tm[e] = true
	}

	it = s.Iterator()
	for it.HasNext() {
		e, _ := it.Next()

		if tm[e] {
			ns.Add(e)
		}
	}

	return ns, nil
}

type hashSetIterator[T comparable] struct {
	it iterator.Iterator[*gmap.Pair[T, present]]
}

func newHashIterator[T comparable](it iterator.Iterator[*gmap.Pair[T, present]]) *hashSetIterator[T] {
	return &hashSetIterator[T]{
		it: it,
	}
}

func (hsi *hashSetIterator[T]) HasNext() bool {
	return hsi.it.HasNext()
}

func (hsi *hashSetIterator[T]) Next() (T, error) {
	if !hsi.it.HasNext() {
		return internal.ZeroValueOf[T](), errors.New("FILL IT") //TODO: FILL IT
	}

	v, err := hsi.it.Next()
	if err != nil {
		return internal.ZeroValueOf[T](), err
	}

	return v.First(), nil
}

func newHashSet[T comparable](data gmap.Map[T, present]) *HashSet[T] {
	return &HashSet[T]{data: data}
}

func (hs *HashSet[T]) insert(e ...T) {
	hs.data.PutAll(toPairs(e...)...)
}

func (hs *HashSet[T]) remove(ignore bool, e ...T) error {
	if hs.IsEmpty() {
		return errors.New("set is empty")
	}

	sz := len(e)
	if sz == 0 {
		return errors.New("argument list is empty")
	}

	for i := 0; i < sz; i++ {
		_, err := hs.data.Remove(e[i])
		if err != nil {
			//TODO: REFACTOR
			if ignore {
				if ex, ok := err.(*erx.Erx); ok {
					if ex.Kind() == "keyNotFoundError" {
						continue
					}
				}
			}

			return err
		}
	}

	return nil
}

func (hs *HashSet[T]) contains(e ...T) bool {
	for _, k := range e {
		if !hs.data.ContainsKey(k) {
			return false
		}
	}

	return true
}

func (hs *HashSet[T]) union(b *HashSet[T]) error {
	it := b.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		hs.Add(v)
	}

	return nil
}

func toPairs[T comparable](e ...T) []*gmap.Pair[T, present] {
	pr := func(k T) *gmap.Pair[T, present] { return gmap.NewPair[T, present](k, present{}) }

	sz := len(e)
	res := make([]*gmap.Pair[T, present], sz)

	for i := 0; i < sz; i++ {
		res[i] = pr(e[i])
	}

	return res
}
