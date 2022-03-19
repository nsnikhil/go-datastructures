package main

type edge[T comparable] struct {
	next   *node[T]
	weight int64
}

func (e *edge[T]) getNext() *node[T] {
	return e.next
}

func (e *edge[T]) getWeight() int64 {
	return e.weight
}

func (e *edge[T]) changeNext(next *node[T]) {
	e.next = next
}

func (e *edge[T]) copy() *edge[T] {
	return &edge[T]{
		next:   e.next.copy(),
		weight: e.weight,
	}
}

func newDiEdge[T comparable](next *node[T]) *edge[T] {
	return newWeightedDiEdge[T](next, 0)
}

func newWeightedDiEdge[T comparable](next *node[T], weight int64) *edge[T] {
	return &edge[T]{
		next:   next,
		weight: weight,
	}
}
