package set

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberr"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/nsnikhil/go-datastructures/utils"
)

var errorDifferentTypes = errors.New("all elements must be of same type")

type present struct{}

type HashSet struct {
	data gmap.Map
}

func NewHashSet(e ...interface{}) (Set, error) {
	hm, err := gmap.NewHashMap()
	if err != nil {
		return nil, err
	}

	hs := newHashSet(hm)

	if len(e) == 0 {
		return hs, nil
	}

	err = insert(hs, e...)
	if err != nil {
		return nil, err
	}

	return hs, nil
}

func (hs *HashSet) Add(e interface{}) error {
	return insert(hs, e)
}

func (hs *HashSet) AddAll(e ...interface{}) error {
	return insert(hs, e...)
}

func (hs *HashSet) Clear() {
	hs.data.Clear()
}

func (hs *HashSet) Contains(e interface{}) bool {
	return contains(hs, e)
}

func (hs *HashSet) ContainsAll(e ...interface{}) bool {
	return contains(hs, e...)
}

func (hs *HashSet) Copy() Set {
	dt := make([]interface{}, 0)

	it := hs.Iterator()
	for it.HasNext() {
		dt = append(dt, it.Next())
	}

	cs, _ := NewHashSet(dt...)
	return cs
}

func (hs *HashSet) IsEmpty() bool {
	return hs.data.IsEmpty()
}

func (hs *HashSet) Size() int {
	return hs.data.Size()
}

func (hs *HashSet) Remove(e interface{}) error {
	return remove(hs, false, e)
}

func (hs *HashSet) RemoveAll(e ...interface{}) error {
	return remove(hs, true, e...)
}

func (hs *HashSet) RetainAll(e ...interface{}) error {
	if hs.IsEmpty() {
		return errors.New("set is empty")
	}

	kt, err := validaTypes(e...)
	if err != nil && err == errorDifferentTypes {
		return err
	}

	tm := make(map[interface{}]bool)

	for _, k := range e {
		tm[k] = true
	}

	validated := false
	dl := make([]interface{}, 0)

	it := hs.Iterator()
	for it.HasNext() {
		e := it.Next()

		if !validated {
			ct := utils.GetTypeName(e)
			if kt != utils.NA && kt != ct {
				return liberr.TypeMismatchError(ct, kt)
			}
			validated = true
		}

		if !tm[e] {
			dl = append(dl, e)
		}
	}

	if len(dl) == 0 {
		return nil
	}

	return remove(hs, false, dl...)
}

func (hs *HashSet) Iterator() iterator.Iterator {
	return newHashIterator(hs.data.Iterator())
}

func (hs *HashSet) Union(s Set) (Set, error) {
	ns, err := NewHashSet()
	if err != nil {
		return nil, err
	}

	if err = union(ns.(*HashSet), hs); err != nil {
		return nil, err
	}

	if err = union(ns.(*HashSet), s.(*HashSet)); err != nil {
		return nil, err
	}

	return ns, nil
}

func (hs *HashSet) Intersection(s Set) (Set, error) {
	ns, err := NewHashSet()
	if err != nil {
		return nil, err
	}

	var kt string
	tm := make(map[interface{}]bool)
	it := hs.Iterator()
	for it.HasNext() {
		e := it.Next()
		if len(kt) == 0 {
			kt = utils.GetTypeName(e)
		}

		tm[e] = true
	}

	var ct string
	it = s.Iterator()
	for it.HasNext() {
		e := it.Next()
		if len(ct) == 0 {
			ct := utils.GetTypeName(e)

			if kt != ct {
				return nil, liberr.TypeMismatchError(kt, ct)
			}
		}

		if tm[e] {
			if err := ns.Add(e); err != nil {
				return nil, err
			}

		}
	}

	return ns, nil
}

type hashSetIterator struct {
	it iterator.Iterator
}

func newHashIterator(it iterator.Iterator) *hashSetIterator {
	return &hashSetIterator{
		it: it,
	}
}

func (hsi *hashSetIterator) HasNext() bool {
	return hsi.it.HasNext()
}

func (hsi *hashSetIterator) Next() interface{} {
	return hsi.it.Next().(*gmap.Pair).GetKey()
}

func newHashSet(data gmap.Map) *HashSet {
	return &HashSet{data: data}
}

func insert(hs *HashSet, e ...interface{}) error {
	if len(e) == 0 {
		return errors.New("argument list is empty")
	}

	return hs.data.PutAll(toPairs(e...)...)
}

func remove(hs *HashSet, ignore bool, e ...interface{}) error {
	if hs.IsEmpty() {
		return errors.New("set is empty")
	}

	sz := len(e)
	if sz == 0 {
		return errors.New("argument list is empty")
	}

	//TODO OPTIMIZE TWO PASS
	kt := utils.GetTypeName(e[0])
	for i := 1; i < sz; i++ {
		ct := utils.GetTypeName(e[i])
		if kt != ct {
			return liberr.TypeMismatchError(kt, ct)
		}
	}

	for i := 0; i < sz; i++ {
		_, err := hs.data.Remove(e[i])
		if err != nil {
			if ignore && err.Error() == fmt.Sprintf("no value found against the key: %v", e[i]) {
				continue
			}

			return err
		}
	}

	return nil
}

func contains(hs *HashSet, e ...interface{}) bool {
	for _, k := range e {
		if !hs.data.ContainsKey(k) {
			return false
		}
	}

	return true
}

func union(a, b *HashSet) error {
	it := b.Iterator()
	for it.HasNext() {
		if err := a.Add(it.Next()); err != nil {
			return err
		}
	}

	return nil
}

func toPairs(e ...interface{}) []*gmap.Pair {
	pr := func(k interface{}) *gmap.Pair { return gmap.NewPair(k, present{}) }

	sz := len(e)
	res := make([]*gmap.Pair, sz)

	for i := 0; i < sz; i++ {
		res[i] = pr(e[i])
	}

	return res
}

func validaTypes(e ...interface{}) (string, error) {
	if len(e) == 0 {
		return utils.NA, errors.New("empty slice")
	}

	kt := utils.GetTypeName(e[0])
	sz := len(e)
	for i := 1; i < sz; i++ {
		ct := utils.GetTypeName(e[i])
		if kt != ct {
			return utils.NA, errorDifferentTypes
		}
	}

	return kt, nil
}
