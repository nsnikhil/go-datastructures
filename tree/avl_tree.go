package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type AVLTree struct {
	c comparator.Comparator
	*BinarySearchTree
}

func NewAVLTree(c comparator.Comparator, e ...interface{}) (*AVLTree, error) {
	bst, err := NewBinarySearchTree(c)
	if err != nil {
		return nil, err
	}

	at := &AVLTree{
		c:                c,
		BinarySearchTree: bst,
	}

	for _, k := range e {
		if err := at.Insert(k); err != nil {
			return nil, err
		}
	}

	return at, nil
}

func Insert(e interface{}) error {
	return errors.New("NOT IMPLEMENTED")
}
