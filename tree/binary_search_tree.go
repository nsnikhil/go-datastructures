package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type BinarySearchTree[T comparable] struct {
	c comparator.Comparator[T]
	*BinaryTree[T]
}

func NewBinarySearchTree[T comparable](c comparator.Comparator[T], e ...T) *BinarySearchTree[T] {
	bt := NewBinaryTree[T]()

	bst := &BinarySearchTree[T]{
		c:          c,
		BinaryTree: bt,
	}

	sz := len(e)

	if sz == 0 {
		return bst
	}

	for _, k := range e {
		bst.InsertCompare(k, c)
	}

	return bst
}

func (bst *BinarySearchTree[T]) Insert(e T) {
	bst.InsertCompare(e, bst.c)
}

func (bst *BinarySearchTree[T]) Delete(e T) error {
	return bst.DeleteCompare(e, bst.c)
}

func (bst *BinarySearchTree[T]) Search(e T) (bool, error) {
	return bst.SearchCompare(e, bst.c)
}
