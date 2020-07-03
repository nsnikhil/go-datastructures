package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type BinarySearchTree struct {
	c comparator.Comparator
	*BinaryTree
}

func NewBinarySearchTree(c comparator.Comparator, e ...interface{}) (*BinarySearchTree, error) {
	bt, err := NewBinaryTree()
	if err != nil {
		return nil, err
	}

	bst := &BinarySearchTree{
		c:          c,
		BinaryTree: bt,
	}

	for _, k := range e {
		if err := bst.InsertCompare(k, c); err != nil {
			return nil, err
		}
	}

	return bst, nil
}

func (bst *BinarySearchTree) Insert(e interface{}) error {
	return errors.New("NOT IMPLEMENTED")
}

func (bst *BinarySearchTree) Delete(e interface{}) error {
	return errors.New("NOT IMPLEMENTED")
}