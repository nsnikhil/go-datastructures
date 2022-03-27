package graph

import "github.com/nsnikhil/go-datastructures/functions/iterator"

type Graph[T comparable] interface {
	AddNode(n *Node[T])

	CreateDiEdge(curr *Node[T], next *Node[T]) error

	CreateWeightedDiEdge(curr, next *Node[T], weight int64) error

	CreateBiEdge(curr *Node[T], next *Node[T]) error

	CreateWeightedBiEdge(curr, nodes *Node[T], weight int64) error

	DeleteNode(n *Node[T]) error

	DeleteEdge(start, end *Node[T]) error

	//print()
	DFSIterator() iterator.Iterator[*Node[T]]
	BFSIterator() iterator.Iterator[*Node[T]]

	HasLoop() bool
	HasCycle() bool
	AreAdjacent(a, b *Node[T]) bool
	DegreeOfNode(a *Node[T]) int
	HasBridge() bool

	Reverse()
	Clone() Graph[T]

	HasRoute(source, target *Node[T]) bool

	IsDirected() bool

	IsConnected() bool

	GetConnectedComponents() [][]*Node[T]

	ShortestPath(source, target *Node[T]) []*Node[T]
}
