package unionfind

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/internal"
	gmap "github.com/nsnikhil/go-datastructures/map"
)

type UnionFind[T comparable] interface {
	Add(e T)
	Find(e T) (T, error)
	Union(a, b T) error
	AreInSameSet(a, b T) (bool, error)
}

type node[T comparable] struct {
	data   T
	weight int64
	par    *node[T]
}

func newNode[T comparable](data T) *node[T] {
	nd := &node[T]{
		data:   data,
		weight: 1,
	}

	nd.par = nd

	return nd
}

type defaultUnionFind[T comparable] struct {
	cache gmap.Map[T, *node[T]]
}

func (duf *defaultUnionFind[T]) Add(e T) {
	if !duf.cache.ContainsKey(e) {
		duf.cache.Put(e, newNode[T](e))
	}
}

func (duf *defaultUnionFind[T]) Find(e T) (T, error) {
	n, err := duf.findParent(e)
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("defaultUnionFind.Find"), err)
	}

	return n.data, nil
}

func (duf *defaultUnionFind[T]) Union(a, b T) error {
	ap, err := duf.findParent(a)
	if err != nil {
		return erx.WithArgs(erx.Kind("defaultUnionFind.Union"), err)
	}

	bp, err := duf.findParent(b)
	if err != nil {
		return erx.WithArgs(erx.Kind("defaultUnionFind.Union"), err)
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

func (duf *defaultUnionFind[T]) AreInSameSet(a, b T) (bool, error) {
	ap, err := duf.findParent(a)
	if err != nil {
		return false, erx.WithArgs(erx.Kind("defaultUnionFind.AreInSameSet"), err)
	}

	bp, err := duf.findParent(b)
	if err != nil {
		return false, erx.WithArgs(erx.Kind("defaultUnionFind.AreInSameSet"), err)
	}

	return ap == bp, nil
}

func (duf *defaultUnionFind[T]) findParent(e T) (*node[T], error) {
	var findParentUtil func(n *node[T]) *node[T]

	findParentUtil = func(n *node[T]) *node[T] {
		if n.par == n {
			return n
		}

		n.par = findParentUtil(n.par)
		return n.par
	}

	n, err := duf.cache.Get(e)
	if err != nil {
		return nil, erx.WithArgs(erx.Kind("defaultUnionFind.findParent"), err)
	}

	return findParentUtil(n), nil
}

func NewUnionFind[T comparable]() UnionFind[T] {
	return &defaultUnionFind[T]{
		cache: gmap.NewHashMap[T, *node[T]](),
	}
}
