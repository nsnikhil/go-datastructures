package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
)

type AVLTree struct {
	c comparator.Comparator
	*BinarySearchTree
}

func NewAVLTree(c comparator.Comparator, e ...interface{}) (*AVLTree, error) {
	return nil, nil
}

func (avt *AVLTree) Insert(e interface{}) error {
	return nil
}

func (avt *AVLTree) Delete(e interface{}) error {
	return nil
}

func (avt *AVLTree) Search(e interface{}) error {
	return nil
}
