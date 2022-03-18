package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewBinarySearchTree(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Tree[int]
		expectedResult func() Tree[int]
	}{
		{
			name: "test create new empty binary search tree",
			actualResult: func() Tree[int] {
				return NewBinarySearchTree(comparator.NewIntegerComparator())
			},
			expectedResult: func() Tree[int] {
				bt := &BinaryTree[int]{count: 0}
				return &BinarySearchTree[int]{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
		{
			name: "test create new empty binary search tree with one element",
			actualResult: func() Tree[int] {
				return NewBinarySearchTree(comparator.NewIntegerComparator(), 1)
			},
			expectedResult: func() Tree[int] {
				bt := &BinaryTree[int]{
					count: 1,
					root:  newBinaryNode(1),
				}
				return &BinarySearchTree[int]{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
		{
			name: "test create new empty binary search tree with multiple element",
			actualResult: func() Tree[int] {
				return NewBinarySearchTree(comparator.NewIntegerComparator(), 2, 1, 3, 4)
			},
			expectedResult: func() Tree[int] {
				bt := &BinaryTree[int]{
					count: 4,
				}
				bt.root = newBinaryNode(2)

				bt.root.left = newBinaryNode(1)
				bt.root.left.parent = bt.root

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.right.right = newBinaryNode(4)
				bt.root.right.right.parent = bt.root.right

				return &BinarySearchTree[int]{
					c:          comparator.NewIntegerComparator(),
					BinaryTree: bt,
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}
