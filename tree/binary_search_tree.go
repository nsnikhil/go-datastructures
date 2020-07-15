package tree

import (
	"errors"
	"fmt"
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
	return bst.InsertCompare(e, bst.c)
}

func (bst *BinarySearchTree) Delete(e interface{}) error {
	if bst.Empty() {
		return errors.New("tree is empty")
	}

	return deleteNode(e, bst.c, bst.root)
}

func deleteNode(e interface{}, c comparator.Comparator, n *binaryNode) error {
	if n == nil {
		return fmt.Errorf("%v not found in the tree", e)
	}

	i, _ := c.Compare(n.data, e)

	if i > 0 {
		if err := deleteNode(e, c, n.left); err != nil {
			return err
		}
	} else if i < 0 {
		if err := deleteNode(e, c, n.right); err != nil {
			return err
		}
	} else {

		if n.isLeaf() {
			n.detach()
			return nil
		} else if n.left == nil {
			n.data = n.right.data
			n.right.detach()
			return nil
		} else if n.right == nil {
			n.data = n.left.data
			n.left.detach()
			return nil
		} else {
			sn := inOrderSuccessor(n.right)
			n.data = sn.data
			sn.detach()
		}

	}

	return nil
}

func inOrderSuccessor(n *binaryNode) *binaryNode {
	c := n
	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}
