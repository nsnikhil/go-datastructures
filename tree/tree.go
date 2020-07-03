package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
)

type Tree interface {
	Insert(e interface{}) error
	Delete(e interface{}) error
	Search(e interface{}) (bool, error)

	Count() int
	Height() int
	Diameter() int
	Empty() bool

	Clear()

	Mirror() (bool, error)
	MirrorAt(e interface{}) (bool, error)

	RotateLeft() error
	RotateRight() error
	RotateLeftAt(e interface{}) error
	RotateRightAt(e interface{}) error

	IsFull() bool
	IsBalanced() bool
	IsPerfect() bool
	IsComplete() bool

	//Balance() error

	LowestCommonAncestor(a, b interface{}) (interface{}, error)

	Paths() (list.List, error)

	PreOrderIterator() iterator.Iterator
	PostOrderIterator() iterator.Iterator
	InOrderIterator() iterator.Iterator
	LevelOrderIterator() iterator.Iterator

	VerticalViewIterator() iterator.Iterator
	LeftViewIterator() iterator.Iterator
	RightViewIterator() iterator.Iterator
	TopViewIterator() iterator.Iterator
	BottomViewIterator() iterator.Iterator
}
