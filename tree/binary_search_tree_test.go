package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewBinarySearchTree(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Tree, error)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test create new empty binary search tree",
			actualResult: func() (Tree, error) {
				return NewBinarySearchTree(comparator.NewIntegerComparator())
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{count: 0, typeURL: "na"}
				return &BinarySearchTree{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
		{
			name: "test create new empty binary search tree with one element",
			actualResult: func() (Tree, error) {
				return NewBinarySearchTree(comparator.NewIntegerComparator(), 1)
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
					root:    newBinaryNode(1),
				}
				return &BinarySearchTree{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
		{
			name: "test create new empty binary search tree with multiple element",
			actualResult: func() (Tree, error) {
				return NewBinarySearchTree(comparator.NewIntegerComparator(), 2, 1, 3, 4)
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   4,
					typeURL: "int",
				}
				bt.root = newBinaryNode(2)

				bt.root.left = newBinaryNode(1)
				bt.root.left.parent = bt.root

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.right.right = newBinaryNode(4)
				bt.root.right.right.parent = bt.root.right

				return &BinarySearchTree{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
		{
			name: "test create new return error when element are of different type",
			actualResult: func() (Tree, error) {
				return NewBinarySearchTree(comparator.NewIntegerComparator(), 1, 'a')
			},
			expectedResult: func() Tree {
				return (*BinarySearchTree)(nil)
			},
			expectedError: errors.New("all elements must be of same type"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}