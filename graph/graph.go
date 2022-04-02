package graph

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
)

type Graph[T any] interface {
	AddNode(n *Node[T])

	CreateDiEdge(curr *Node[T], next *Node[T]) error

	CreateWeightedDiEdge(curr, next *Node[T], weight int64) error

	CreateBiEdge(curr *Node[T], next *Node[T]) error

	CreateWeightedBiEdge(curr, nodes *Node[T], weight int64) error

	DeleteNode(n *Node[T]) error

	DeleteEdge(start, end *Node[T]) error

	Contains(n *Node[T]) bool

	//print()
	DFSIterator() iterator.Iterator[*Node[T]]
	BFSIterator() iterator.Iterator[*Node[T]]

	HasLoop() bool
	HasCycle() bool
	//HasBridge() bool

	AreAdjacent(a, b *Node[T]) (bool, error)

	InDegreeOfNode(a *Node[T]) (int64, error)
	OutDegreeOfNode(a *Node[T]) (int64, error)

	Reverse()
	Clone() Graph[T]

	HasRoute(source, target *Node[T]) (bool, error)

	//IsDirected() bool

	//IsConnected() bool

	GetConnectedComponents() []list.List[*Node[T]]

	ShortestPath(source, target *Node[T], properties ...Property) (list.List[*Node[T]], error)
}
