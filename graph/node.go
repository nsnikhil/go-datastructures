package main

import "math"

type node[T comparable] struct {
	data  T
	edges map[*edge[T]]bool

	// SHORTEST PATH
	costToReach int64
	predecessor *node[T]
}

func (n *node[T]) getData() T {
	return n.data
}

func (n *node[T]) getEdges() map[*edge[T]]bool {
	return n.edges
}

func (n *node[T]) clearEdges() {
	n.edges = make(map[*edge[T]]bool)
}

func (n *node[T]) removeEdge(e *edge[T]) {
	if !n.hasEdge(e) {
		return
	}

	delete(n.edges, e)
}

func (n *node[T]) addEdge(e *edge[T]) {
	if n.hasEdge(e) {
		return
	}

	n.edges[e] = true
}

func (n *node[T]) hasEdge(e *edge[T]) bool {
	return n.edges[e]
}

func (n *node[T]) addEdges(edges ...*edge[T]) {
	for _, e := range edges {
		n.addEdge(e)
	}
}

func (n *node[T]) areConnected(o *node[T]) bool {
	for k := range n.edges {
		if k.getNext() == o {
			return true
		}
	}

	return false
}

func (n *node[T]) copy() *node[T] {
	copyEdges := func(edges map[*edge[T]]bool) map[*edge[T]]bool {
		res := make(map[*edge[T]]bool)

		for e, ok := range edges {
			res[e.copy()] = ok
		}

		return res
	}

	return &node[T]{
		data:        n.data,
		costToReach: math.MaxInt64,
		edges:       copyEdges(n.edges),
	}
}

func newNode[T comparable](data T) *node[T] {
	return &node[T]{
		data:        data,
		costToReach: math.MaxInt64,
		edges:       make(map[*edge[T]]bool),
	}
}
