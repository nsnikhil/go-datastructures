package main

import "math"

type node struct {
	data  int
	edges map[*edge]bool

	// SHORTEST PATH
	costToReach int
	predecessor *node
}

func (n *node) getData() int {
	return n.data
}

func (n *node) getEdges() map[*edge]bool {
	return n.edges
}

func (n *node) clearEdges() {
	n.edges = make(map[*edge]bool)
}

func (n *node) removeEdge(e *edge) {
	if !n.hasEdge(e) {
		return
	}

	delete(n.edges, e)
}

func (n *node) addEdge(e *edge) {
	if n.hasEdge(e) {
		return
	}

	n.edges[e] = true
}

func (n *node) hasEdge(e *edge) bool {
	return n.edges[e]
}

func (n *node) addEdges(edges ...*edge) {
	for _, e := range edges {
		n.addEdge(e)
	}
}

func (n *node) areConnected(o *node) bool {
	for k := range n.edges {
		if k.getNext() == o {
			return true
		}
	}

	return false
}

func newNode(data int) *node {
	return &node{
		data:        data,
		costToReach: math.MaxInt32,
		edges:       make(map[*edge]bool),
	}
}
