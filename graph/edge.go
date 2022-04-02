package graph

import (
	"github.com/nsnikhil/go-datastructures/internal"
)

type edge[T any] struct {
	next   *Node[T]
	weight int64
}

func (e *edge[T]) changeWeight(weight int64) {
	e.weight = weight
}

func (e *edge[T]) changeNext(next *Node[T]) {
	e.next = next
}

func (e *edge[T]) copy() *edge[T] {
	return &edge[T]{
		next:   e.next.copy(),
		weight: e.weight,
	}
}

func newDiEdge[T any](next *Node[T]) *edge[T] {
	return newEdge[T](next, internal.Zero)
}

func newWeightedDiEdge[T any](next *Node[T], weight int64) *edge[T] {
	return newEdge[T](next, weight)
}

func newEdge[T any](next *Node[T], weight int64) *edge[T] {
	return &edge[T]{
		next:   next,
		weight: weight,
	}
}
