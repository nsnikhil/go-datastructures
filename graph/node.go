package graph

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/nsnikhil/go-datastructures/stack"
)

type Node[T comparable] struct {
	data  T
	edges set.Set[*edge[T]]

	// SHORTEST PATH
	//costToReach int64
	//predecessor *Node[T]
}

func (n *Node[T]) addEdge(e *edge[T]) {
	if n.edges.Contains(e) {
		return
	}

	n.edges.Add(e)
}

func (n *Node[T]) removeEdge(e *edge[T]) error {
	if !n.edges.Contains(e) {
		return edgeNotFoundError(n.data, e.next.data, "Node.removeEdge")
	}

	if err := n.edges.Remove(e); err != nil {
		return edgeNotFoundError(n.data, e.next.data, "Node.removeEdge")
	}

	return nil
}

func (n *Node[T]) clearEdges() {
	n.edges.Clear()
}

func (n *Node[T]) findEdge(o *Node[T]) (*edge[T], error) {
	it := n.edges.Iterator()

	for it.HasNext() {
		e, _ := it.Next()
		if e.next == o {
			return e, nil
		}
	}

	return nil, edgeNotFoundError(n.data, o.data, "Node.findEdge")
}

func (n *Node[T]) copy() *Node[T] {
	copyEdges := func(edges set.Set[*edge[T]]) set.Set[*edge[T]] {
		res := set.NewHashSet[*edge[T]]()

		it := edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			res.Add(e.copy())
		}

		return res
	}

	return &Node[T]{
		data:  n.data,
		edges: copyEdges(n.edges),
	}
}

func (n *Node[T]) bfsIterator() iterator.Iterator[*Node[T]] {
	return newNodeBfsIterator[T](n)
}

func (n *Node[T]) dfsIterator() iterator.Iterator[*Node[T]] {
	return newNodeDfsIterator[T](n)
}

type nodeBfsIterator[T comparable] struct {
	qu queue.Queue[*Node[T]]
	vs set.Set[*Node[T]]
}

func (nbi *nodeBfsIterator[T]) HasNext() bool {
	return !nbi.qu.Empty()
}

func (nbi *nodeBfsIterator[T]) Next() (*Node[T], error) {
	v, err := nbi.qu.Remove()
	if err != nil {
		return nil, emptyIteratorError("nodeBfsIterator.Next")
	}

	it := v.edges.Iterator()
	for it.HasNext() {
		ed, _ := it.Next()
		if !nbi.vs.Contains(ed.next) {
			nbi.qu.Add(ed.next)
			nbi.vs.Add(ed.next)
		}
	}

	return v, nil
}

//TODO: REFACTOR REMOVE ADDING ALL EDGES AT ONCE
func newNodeBfsIterator[T comparable](n *Node[T]) iterator.Iterator[*Node[T]] {
	qu := queue.NewLinkedQueue[*Node[T]]()
	qu.Add(n)

	vs := set.NewHashSet[*Node[T]]()
	vs.Add(n)

	return &nodeBfsIterator[T]{
		qu: qu,
		vs: vs,
	}
}

//TODO: RENAME
func newNodeBfsIteratorWithVisited[T comparable](n *Node[T], vs set.Set[*Node[T]]) iterator.Iterator[*Node[T]] {
	qu := queue.NewLinkedQueue[*Node[T]]()
	qu.Add(n)

	vs.Add(n)

	return &nodeBfsIterator[T]{
		qu: qu,
		vs: vs,
	}
}

type nodeDfsIterator[T comparable] struct {
	st *stack.Stack[*Node[T]]
	vs set.Set[*Node[T]]
}

func (nbi *nodeDfsIterator[T]) HasNext() bool {
	return !nbi.st.Empty()
}

//TODO: REFACTOR REMOVE ADDING ALL EDGES AT ONCE
func (nbi *nodeDfsIterator[T]) Next() (*Node[T], error) {
	v, err := nbi.st.Pop()
	if err != nil {
		return nil, emptyIteratorError("nodeDfsIterator.Next")
	}

	it := v.edges.Iterator()
	for it.HasNext() {
		ed, _ := it.Next()
		if !nbi.vs.Contains(ed.next) {
			nbi.st.Push(ed.next)
			nbi.vs.Add(ed.next)
		}
	}

	return v, nil
}

func newNodeDfsIterator[T comparable](n *Node[T]) iterator.Iterator[*Node[T]] {
	st := stack.NewStack[*Node[T]]()
	st.Push(n)

	vs := set.NewHashSet[*Node[T]]()
	vs.Add(n)

	return &nodeDfsIterator[T]{
		st: st,
		vs: vs,
	}
}

//TODO: RENAME
func newNodeDfsIteratorWithVisited[T comparable](n *Node[T], vs set.Set[*Node[T]]) iterator.Iterator[*Node[T]] {
	st := stack.NewStack[*Node[T]]()
	st.Push(n)

	vs.Add(n)

	return &nodeDfsIterator[T]{
		st: st,
		vs: vs,
	}
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{
		data:  data,
		edges: set.NewHashSet[*edge[T]](),
	}
}
