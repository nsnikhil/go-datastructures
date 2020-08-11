package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/utils"
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

func (bst *BinarySearchTree) Insert(e interface{}) error {
	return bst.InsertCompare(e, bst.c)
}

func (bst *BinarySearchTree) Delete(e interface{}) error {
	return bst.DeleteCompare(e, bst.c)
}

func (bst *BinarySearchTree) Search(e interface{}) (bool, error) {
	return bst.SearchCompare(e, bst.c)
}
