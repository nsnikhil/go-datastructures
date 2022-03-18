package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
)

type Tree[T comparable] interface {
	Insert(e T)
	Delete(e T) error
	Search(e T) (bool, error)

	Count() int
	Height() int
	Diameter() int
	Empty() bool

	Clear()
	Clone() Tree[T]

	//Mirror() (bool, error)
	//MirrorAt(e T) (bool, error)
	//
	//RotateLeft() error
	//RotateRight() error
	//RotateLeftAt(e T) error
	//RotateRightAt(e T) error

	IsFull() bool
	IsBalanced() bool
	IsPerfect() bool
	IsComplete() bool

	//Balance() error

	LowestCommonAncestor(a, b T) (T, error)

	Paths() ([][]T, error)

	// TEMPORARY
	Mode() (list.List[T], error)
	Equal(t Tree[T]) (bool, error)
	//Symmetric() bool
	//Invert()

	InOrderSuccessor(e T) (T, error)
	PreOrderSuccessor(e T) (T, error)
	PostOrderSuccessor(e T) (T, error)
	LevelOrderSuccessor(e T) (T, error)

	PreOrderIterator() iterator.Iterator[T]
	PostOrderIterator() iterator.Iterator[T]
	InOrderIterator() iterator.Iterator[T]
	LevelOrderIterator() iterator.Iterator[T]

	VerticalViewIterator() iterator.Iterator[T]
	LeftViewIterator() iterator.Iterator[T]
	RightViewIterator() iterator.Iterator[T]
	TopViewIterator() iterator.Iterator[T]
	BottomViewIterator() iterator.Iterator[T]
}
