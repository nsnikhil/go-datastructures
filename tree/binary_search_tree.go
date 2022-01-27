package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/utils"
)

type BinarySearchTree[T comparable] struct {
	c comparator.Comparator[T]
	*BinaryTree[T]
}

func NewBinarySearchTree[T comparable](c comparator.Comparator[T], e ...T) (*BinarySearchTree[T], error) {
	bt, err := NewBinaryTree[T]()
	if err != nil {
		return nil, err
	}

	bst := &BinarySearchTree[T]{
		c:          c,
		BinaryTree: bt,
	}

	sz := len(e)

	if sz == 0 {
		return bst, nil
	}

	et := utils.GetTypeName(e[0])

	for i := 1; i < sz; i++ {
		if et != utils.GetTypeName(e[i]) {
			return nil, errors.New("all elements must be of same type")
		}
	}

	for _, k := range e {
		if err := bst.InsertCompare(k, c); err != nil {
			return nil, err
		}
	}

	return bst, nil
}

func (bst *BinarySearchTree[T]) Insert(e T) error {
	return bst.InsertCompare(e, bst.c)
}

func (bst *BinarySearchTree[T]) Delete(e T) error {
	return bst.DeleteCompare(e, bst.c)
}

func (bst *BinarySearchTree[T]) Search(e T) (bool, error) {
	return bst.SearchCompare(e, bst.c)
}
