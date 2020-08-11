package set

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/nsnikhil/go-datastructures/utils"
)

type TreeSet struct {
	c    comparator.Comparator
	data gmap.Map
}

func NewTreeSet(c comparator.Comparator, e ...interface{}) (Set, error) {
	tm, err := gmap.NewTreeMap(c)
	if err != nil {
		return nil, err
	}

	ts := &TreeSet{
		c:    c,
		data: tm,
	}

	return ts, nil
}

func (ts *TreeSet) Add(e interface{}) error {
	return nil
}

func (ts *TreeSet) AddAll(e ...interface{}) error {
	return nil
}

func (ts *TreeSet) Clear() {

}

func (ts *TreeSet) Contains(e interface{}) bool {
	return false
}

func (ts *TreeSet) ContainsAll(e ...interface{}) bool {
	return false
}

func (ts *TreeSet) Copy() Set {
	return nil
}

func (ts *TreeSet) IsEmpty() bool {
	return false
}

func (ts *TreeSet) Size() int {
	return utils.InvalidIndex
}

func (ts *TreeSet) Remove(e interface{}) error {
	return nil
}

func (ts *TreeSet) RemoveAll(e ...interface{}) error {
	return nil
}

func (ts *TreeSet) RetainAll(e ...interface{}) error {
	return nil
}

func (ts *TreeSet) Iterator() iterator.Iterator {
	return nil
}

func (ts *TreeSet) Union(s Set) (Set, error) {
	return nil, nil
}

func (ts *TreeSet) Intersection(s Set) (Set, error) {
	return nil, nil
}
