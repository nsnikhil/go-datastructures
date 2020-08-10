package main

type edge struct {
	next   *node
	weight int
}

func (e *edge) getNext() *node {
	return e.next
}

func (e *edge) getWeight() int {
	return e.weight
}

func (e *edge) changeNext(next *node) {
	e.next = next
}

func newDiEdge(next *node) *edge {
	return newWeightedDiEdge(next, 0)
}

func newWeightedDiEdge(next *node, weight int) *edge {
	return &edge{
		next:   next,
		weight: weight,
	}
}
