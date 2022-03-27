package graph

import (
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/set"
	"sort"
)

//TODO: FAILS FOR CYCLE
func areNodeEqual(a, b *Node[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.data != b.data || !areEdgesEqual(a.edges, b.edges) {
		return false
	}

	return true
}

func areEdgesEqual(a, b set.Set[*edge[int]]) bool {
	if a.Size() != b.Size() {
		return false
	}

	ak := getKeys(a)
	bk := getKeys(b)

	if ak.Size() != bk.Size() {
		return false
	}

	aIt := ak.Iterator()
	bIt := bk.Iterator()

	for aIt.HasNext() && bIt.HasNext() {
		an, _ := aIt.Next()
		bn, _ := bIt.Next()

		if an.weight != bn.weight {
			return false
		}

		if !areNodeEqual(an.next, bn.next) {
			return false
		}
	}

	return true
}

type edgeComparator struct{}

func (e edgeComparator) Compare(a, b *edge[int]) int {
	return a.next.data - b.next.data
}

func getKeys(s set.Set[*edge[int]]) list.List[*edge[int]] {
	res := list.NewArrayList[*edge[int]]()

	it := s.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		res.Add(v)
	}

	res.Sort(edgeComparator{})

	return res
}

func simplifyGraph(g *listGraph[int]) map[int][]int {
	res := make(map[int][]int)

	ni := g.nodes.Iterator()
	for ni.HasNext() {
		n, _ := ni.Next()
		vi := n.edges.Iterator()
		for vi.HasNext() {
			v, _ := vi.Next()
			res[n.data] = append(res[n.data], v.next.data)
		}
	}

	return res
}

type intSliceComparator struct{}

func (isc intSliceComparator) Compare(a, b []int) int {
	asz := len(a)
	bsz := len(b)

	if asz != bsz {
		return asz - bsz
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < asz; i++ {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}

	return 0
}
