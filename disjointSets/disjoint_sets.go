package disjointSets

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/internal"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/nsnikhil/go-datastructures/set"
)

type DisjointSets[T comparable] interface {
	MakeSet(e T)
	FindSet(e T) (T, error)
	Union(a, b T) error
	AreInSameSet(a, b T) (bool, error)
	SetsCount() int64
	GetAllSets() []set.Set[T]
	Clear()
}

type node[T comparable] struct {
	data   T
	weight int64
	par    *node[T]
}

func (n *node[T]) findParent() *node[T] {
	if n.par == n {
		return n
	}

	n.par = n.par.findParent()
	return n.par
}

func newNode[T comparable](data T) *node[T] {
	nd := &node[T]{
		data:   data,
		weight: 1,
	}

	nd.par = nd

	return nd
}

//TODO: RENAME DEFAULT
type defaultDisjointSets[T comparable] struct {
	sets gmap.Map[T, *node[T]]
}

func (dds *defaultDisjointSets[T]) MakeSet(e T) {
	if !dds.sets.ContainsKey(e) {
		dds.sets.Put(e, newNode[T](e))
	}
}

func (dds *defaultDisjointSets[T]) FindSet(e T) (T, error) {
	n, err := dds.findParent(e)
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("defaultDisjointSets.FindSet"), err)
	}

	return n.data, nil
}

func (dds *defaultDisjointSets[T]) Union(a, b T) error {
	ap, err := dds.findParent(a)
	if err != nil {
		return erx.WithArgs(erx.Kind("defaultDisjointSets.Union"), err)
	}

	bp, err := dds.findParent(b)
	if err != nil {
		return erx.WithArgs(erx.Kind("defaultDisjointSets.Union"), err)
	}

	if ap == bp {
		return nil
	}

	if bp.weight > ap.weight {
		ap.par = bp
		bp.weight++
	} else {
		bp.par = ap
		ap.weight++
	}

	return nil
}

func (dds *defaultDisjointSets[T]) AreInSameSet(a, b T) (bool, error) {
	ap, err := dds.findParent(a)
	if err != nil {
		return false, erx.WithArgs(erx.Kind("defaultDisjointSets.AreInSameSet"), err)
	}

	bp, err := dds.findParent(b)
	if err != nil {
		return false, erx.WithArgs(erx.Kind("defaultDisjointSets.AreInSameSet"), err)
	}

	return ap == bp, nil
}

func (dds *defaultDisjointSets[T]) SetsCount() int64 {
	return int64(len(dds.getAllUniqueSets()))
}

func (dds *defaultDisjointSets[T]) GetAllSets() []set.Set[T] {
	uniqueSets := dds.getAllUniqueSets()

	var res []set.Set[T]

	for _, v := range uniqueSets {
		res = append(res, set.NewHashSet[T](v...))
	}

	return res
}

func (dds *defaultDisjointSets[T]) Clear() {
	dds.sets = gmap.NewHashMap[T, *node[T]]()
}

func (dds *defaultDisjointSets[T]) getAllUniqueSets() map[T][]T {
	temp := map[T][]T{}

	it := dds.sets.Iterator()

	for it.HasNext() {
		v, _ := it.Next()

		vPar, _ := dds.findParent(T(v.First()))

		temp[vPar.data] = append(temp[vPar.data], v.First())
	}

	return temp
}

func (dds *defaultDisjointSets[T]) findParent(e T) (*node[T], error) {
	n, err := dds.sets.Get(e)
	if err != nil {
		return nil, erx.WithArgs(erx.Kind("defaultDisjointSets.findParent"), err)
	}

	return n.findParent(), nil
}

func NewDisjointSets[T comparable](elements ...T) DisjointSets[T] {
	ds := &defaultDisjointSets[T]{
		sets: gmap.NewHashMap[T, *node[T]](),
	}

	for _, e := range elements {
		ds.MakeSet(e)
	}

	return ds
}
