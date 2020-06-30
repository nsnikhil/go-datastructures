package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewBinaryTree(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Tree, error)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test create empty binary tree",
			actualResult: func() (Tree, error) {
				return NewBinaryTree()
			},
			expectedResult: func() Tree {
				return &BinaryTree{
					count:   0,
					typeURL: "na",
				}
			},
		},
		{
			name: "test binary tree with values",
			actualResult: func() (Tree, error) {
				return NewBinaryTree(1, 2, 3, 4)
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   4,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)

				bt.root.left = newBinaryNode(2)
				bt.root.left.parent = bt.root

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.left.left = newBinaryNode(4)
				bt.root.left.left.parent = bt.root.left

				return bt
			},
		},
		{
			name: "test failed to create binary tree due to type mismatch",
			actualResult: func() (Tree, error) {
				return NewBinaryTree(1, 2, 'a')
			},
			expectedResult: func() Tree {
				return (*BinaryTree)(nil)
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

func TestBinaryTreeInsert(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test insert one node in binary tree",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.Insert(1), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)
				return bt
			},
		},
		{
			name: "test insert multiple node in binary tree",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				require.NoError(t, bt.Insert(1))
				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				return bt.Insert(4), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   4,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)

				bt.root.left = newBinaryNode(2)
				bt.root.left.parent = bt.root

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.left.left = newBinaryNode(4)
				bt.root.left.left.parent = bt.root.left

				return bt
			},
		},
		{
			name: "test failed to insert due to type mismatch",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				require.NoError(t, bt.Insert(1))
				return bt.Insert('a'), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)
				return bt
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeInsertCompare(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test insert one node in binary tree with compare",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.InsertCompare(1, comparator.NewIntegerComparator()), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)
				return bt
			},
		},
		{
			name: "test insert multiple node in binary tree with compare",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(1, c))
				require.NoError(t, bt.InsertCompare(3, c))
				return bt.InsertCompare(4, c), bt
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

				return bt
			},
		},
		{
			name: "test insert compare return error when type mismatch",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(1, c))

				return bt.InsertCompare('a', c), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)
				return bt
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test insert compare return error when comparator type is different",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				c := comparator.NewStringComparator()

				require.NoError(t, bt.InsertCompare(1, c))

				return bt.InsertCompare(2, c), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{
					count:   1,
					typeURL: "int",
				}
				bt.root = newBinaryNode(1)
				return bt
			},
			expectedError: liberror.NewTypeMismatchError("string", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeDelete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test delete root when tree consists of only one element",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.Delete(1), bt
			},
			expectedResult: func() Tree {
				return &BinaryTree{
					count:   0,
					typeURL: "int",
				}
			},
		},
		{
			name: "test delete root when tree consists of root and left child",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))

				return bt.Delete(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test delete root when tree consists root and right child",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.InsertCompare(2, comparator.NewIntegerComparator()))

				return bt.Delete(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test delete left most element when right is absent",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(4, c))
				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(1, c))

				return bt.Delete(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(4, c))
				require.NoError(t, bt.InsertCompare(2, c))

				return bt
			},
		},
		{
			name: "test delete left most element when right is present",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(4, c))
				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(1, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(11, c))

				return bt.Delete(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(4, c))
				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(12, c))

				curr := bt.root

				for curr.left != nil {
					curr = curr.left
				}

				curr.left = newBinaryNode(11)
				curr.left.parent = curr
				bt.count++

				return bt
			},
		},
		{
			name: "test delete right most element when left is absent",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(14, c))

				return bt.Delete(14), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(12, c))

				return bt
			},
		},
		{
			name: "test delete right most element when left is present",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(0, c))
				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(4, c))

				return bt.Delete(14), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(8)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(0, c))
				require.NoError(t, bt.InsertCompare(2, c))
				require.NoError(t, bt.InsertCompare(4, c))

				return bt
			},
		},
		{
			name: "test delete mid element",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				return bt.Delete(7), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(16))
				require.NoError(t, bt.Insert(14))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(8))
				require.NoError(t, bt.Insert(12))

				return bt
			},
		},
		{
			name: "test delete return error when element is not present",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(1, c))
				require.NoError(t, bt.InsertCompare(3, c))

				return bt.Delete(7), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(1, c))
				require.NoError(t, bt.InsertCompare(3, c))

				return bt
			},
			expectedError: errors.New("7 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreePreOrderIterator(t *testing.T) {
	testCase := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test preorder iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.PreOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test preorder iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.PreOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "test preorder iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.PreOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 7, 6, 8, 14, 12, 16},
		},
		{
			name: "test preorder iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.PreOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 20, 15, 12, 17, 19, 25},
		},
		{
			name: "test preorder iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.PreOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{25, 17, 14, 10, 7, 15, 20},
		},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreePostOrderIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test post order iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.PostOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test post order iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.PostOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name: "test post order iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.PostOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{6, 8, 7, 12, 16, 14, 10},
		},
		{
			name: "test post order iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.PostOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{12, 19, 17, 15, 25, 20, 10},
		},
		{
			name: "test post order iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.PostOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{7, 10, 15, 14, 20, 17, 25},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeInOrderIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test in order iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.InOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test in order iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.InOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "test in order iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.InOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{6, 7, 8, 10, 12, 14, 16},
		},
		{
			name: "test in order iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.InOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 12, 15, 17, 19, 20, 25},
		},
		{
			name: "test in order iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.InOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{7, 10, 14, 15, 17, 20, 25},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeLevelOrderIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test level order iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.LevelOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test level order iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.LevelOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "test level order iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.LevelOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 7, 14, 6, 8, 12, 16},
		},
		{
			name: "test level order iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.LevelOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 20, 15, 25, 12, 17, 19},
		},
		{
			name: "test level order iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.LevelOrderIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{25, 17, 14, 20, 10, 15, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeSearch(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		{
			name: "test search in binary tree",
			actualResult: func() (bool, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4)
				require.NoError(t, err)

				return bt.Search(1)
			},
			expectedResult: true,
		},
		{
			name: "test search in binary tree two",
			actualResult: func() (bool, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4)
				require.NoError(t, err)

				return bt.Search(4)
			},
			expectedResult: true,
		},
		{
			name: "test search in binary tree return false when element is not present",
			actualResult: func() (bool, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4)
				require.NoError(t, err)

				return bt.Search(5)
			},
			expectedResult: false,
			expectedError:  errors.New("5 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestBinaryTreeCount(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test count element in binary tree",
			actualResult: func() int {
				bt, err := NewBinaryTree(1, 2, 3, 4)
				require.NoError(t, err)

				return bt.count
			},
			expectedResult: 4,
		},
		{
			name: "test count element in binary tree two",
			actualResult: func() int {
				bt, err := NewBinaryTree(1, 2)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Delete(3))
				require.NoError(t, bt.InsertCompare(6, comparator.NewIntegerComparator()))
				require.NoError(t, bt.InsertCompare(0, comparator.NewIntegerComparator()))
				require.NoError(t, bt.Delete(2))

				return bt.count
			},
			expectedResult: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when tree is empty",
			actualResult: func() bool {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true when tree is empty after operations",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1, 2)
				require.NoError(t, err)

				require.NoError(t, bt.Delete(2))
				require.NoError(t, bt.Delete(1))

				return bt.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when tree is not empty",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1, 2)
				require.NoError(t, err)

				return bt.Empty()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Tree
		expectedResult func() Tree
	}{
		{
			name: "test clear empty tree",
			actualResult: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				bt.Clear()

				return bt
			},
			expectedResult: func() Tree {
				return &BinaryTree{typeURL: "na", count: 0}
			},
		},
		{
			name: "test clear tree with data",
			actualResult: func() Tree {
				bt, err := NewBinaryTree('a')
				require.NoError(t, err)

				require.NoError(t, bt.Insert('b'))

				bt.Clear()

				return bt
			},
			expectedResult: func() Tree {
				return &BinaryTree{typeURL: "int32", count: 0}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestBinaryTreeMirror(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error, Tree)
		expectedResult bool
		expectedTree   func() Tree
		expectedError  error
	}{
		{
			name: "test mirror tree with one node",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror tree with multiple nodes",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(7))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(4))

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror tree with multiple nodes two",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(25)

				bt.root.right = newBinaryNode(17)
				bt.root.right.parent = bt.root

				bt.root.right.left = newBinaryNode(20)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.right = newBinaryNode(14)
				bt.root.right.right.parent = bt.root.right

				bt.root.right.right.left = newBinaryNode(15)
				bt.root.right.right.left.parent = bt.root.right.right

				bt.root.right.right.right = newBinaryNode(10)
				bt.root.right.right.right.parent = bt.root.right.right

				bt.root.right.right.right.right = newBinaryNode(7)
				bt.root.right.right.right.right.parent = bt.root.right.right.right

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror tree with multiple nodes three",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(10)

				bt.root.left = newBinaryNode(20)
				bt.root.left.parent = bt.root

				bt.root.left.left = newBinaryNode(25)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.right = newBinaryNode(15)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.right.right = newBinaryNode(12)
				bt.root.left.right.right.parent = bt.root.left.right

				bt.root.left.right.left = newBinaryNode(17)
				bt.root.left.right.left.parent = bt.root.left.right

				bt.root.left.right.left.left = newBinaryNode(19)
				bt.root.left.right.left.left.parent = bt.root.left.right.left

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror tree twice gives back original tree",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				ok, err := bt.Mirror()
				require.NoError(t, err)
				require.True(t, ok)

				ok, err = bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror empty tree return false",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedResult: false,
			expectedError:  errors.New("tree is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ok, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, ok)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTree(), res)
		})
	}
}

func TestBinaryTreeMirrorAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error, Tree)
		expectedResult bool
		expectedTree   func() Tree
		expectedError  error
	}{
		{
			name: "test mirror at tree with one node",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				ok, err := bt.MirrorAt(1)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror at tree with multiple nodes",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				ok, err := bt.MirrorAt(2)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror at tree with multiple nodes two",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				ok, err := bt.MirrorAt(14)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(25)

				bt.root.left = newBinaryNode(17)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(20)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.left = newBinaryNode(14)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(15)
				bt.root.left.left.left.parent = bt.root.left.left

				bt.root.left.left.right = newBinaryNode(10)
				bt.root.left.left.right.parent = bt.root.left.left

				bt.root.left.left.right.right = newBinaryNode(7)
				bt.root.left.left.right.right.parent = bt.root.left.left.right

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror at tree with multiple nodes three",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				ok, err := bt.MirrorAt(15)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(10)

				bt.root.right = newBinaryNode(20)
				bt.root.right.parent = bt.root

				bt.root.right.right = newBinaryNode(25)
				bt.root.right.right.parent = bt.root.right

				bt.root.right.left = newBinaryNode(15)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.left.right = newBinaryNode(12)
				bt.root.right.left.right.parent = bt.root.right.left

				bt.root.right.left.left = newBinaryNode(17)
				bt.root.right.left.left.parent = bt.root.right.left

				bt.root.right.left.left.left = newBinaryNode(19)
				bt.root.right.left.left.left.parent = bt.root.right.left.left

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror at tree twice gives back original tree",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				ok, err := bt.MirrorAt(3)
				require.NoError(t, err)
				require.True(t, ok)

				ok, err = bt.MirrorAt(3)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				return bt
			},
			expectedResult: true,
		},
		{
			name: "test mirror empty tree return false",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				ok, err := bt.Mirror()

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedResult: false,
			expectedError:  errors.New("tree is empty"),
		},
		{
			name: "test mirror return false when value is not present",
			actualResult: func() (bool, error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5)
				require.NoError(t, err)

				ok, err := bt.MirrorAt(6)

				return ok, err, bt
			},
			expectedTree: func() Tree {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return bt
			},
			expectedResult: false,
			expectedError:  errors.New("6 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ok, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, ok)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTree(), res)
		})
	}
}

func TestBinaryTreeHeight(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test get height of empty tree",
			actualResult: func() int {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.Height()
			},
			expectedResult: 0,
		},
		{
			name: "test get height with multiple nodes",
			actualResult: func() int {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.Height()
			},
			expectedResult: 3,
		},
		{
			name: "test get height with multiple nodes two",
			actualResult: func() int {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.Height()
			},
			expectedResult: 5,
		},
		{
			name: "test get height with multiple nodes three",
			actualResult: func() int {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				return bt.Height()
			},
			expectedResult: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeDiameter(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test get diameter of empty tree",
			actualResult: func() int {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.Diameter()
			},
			expectedResult: 0,
		},
		{
			name: "test get diameter with multiple nodes",
			actualResult: func() int {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.Diameter()
			},
			expectedResult: 5,
		},
		{
			name: "test get diameter with multiple nodes two",
			actualResult: func() int {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(30, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.Diameter()
			},
			expectedResult: 6,
		},
		{
			name: "test get diameter with multiple nodes three",
			actualResult: func() int {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				return bt.Diameter()
			},
			expectedResult: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeIsBalanced(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return true for empty tree",
			actualResult: func() bool {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.IsBalanced()
			},
			expectedResult: true,
		},
		{
			name: "return true for one node tree",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.IsBalanced()
			},
			expectedResult: true,
		},
		{
			name: "return true when tree is balanced",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				return bt.IsBalanced()
			},
			expectedResult: true,
		},
		{
			name: "return true when tree is balanced two",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))
				require.NoError(t, bt.Insert(8))
				require.NoError(t, bt.Insert(9))

				return bt.IsBalanced()
			},
			expectedResult: true,
		},
		{
			name: "return false when tree is not balanced",
			actualResult: func() bool {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.IsBalanced()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeRotateRight(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test rotate tree with one node",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.RotateRight(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.RotateRight(), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(2)

				bt.root.left = newBinaryNode(4)
				bt.root.left.parent = bt.root

				bt.root.right = newBinaryNode(1)
				bt.root.right.parent = bt.root

				bt.root.right.left = newBinaryNode(5)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.right = newBinaryNode(3)
				bt.root.right.right.parent = bt.root.right

				bt.root.right.right.left = newBinaryNode(6)
				bt.root.right.right.left.parent = bt.root.right.right

				bt.root.right.right.right = newBinaryNode(7)
				bt.root.right.right.right.parent = bt.root.right.right

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes two",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.RotateRight(), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(17)

				bt.root.right = newBinaryNode(25)
				bt.root.right.parent = bt.root

				bt.root.right.left = newBinaryNode(20)
				bt.root.right.left.parent = bt.root.right

				bt.root.left = newBinaryNode(14)
				bt.root.left.parent = bt.root

				bt.root.left.left = newBinaryNode(10)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.right = newBinaryNode(15)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(7)
				bt.root.left.left.left.parent = bt.root.left.left

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes three",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.RotateRight(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt
			},
		},
		{
			name: "test rotate return error when tree is empty",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.RotateRight(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("tree is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeRotateLeft(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test rotate tree with one node",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.RotateLeft(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.RotateLeft(), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(3)

				bt.root.right = newBinaryNode(7)
				bt.root.right.parent = bt.root

				bt.root.left = newBinaryNode(1)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(6)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.left = newBinaryNode(2)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(4)
				bt.root.left.left.left.parent = bt.root.left.left

				bt.root.left.left.right = newBinaryNode(5)
				bt.root.left.left.right.parent = bt.root.left.left

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes two",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.RotateLeft(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes three",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.RotateLeft(), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(20)

				bt.root.right = newBinaryNode(25)
				bt.root.right.parent = bt.root

				bt.root.left = newBinaryNode(10)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(15)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.right.left = newBinaryNode(12)
				bt.root.left.right.left.parent = bt.root.left.right

				bt.root.left.right.right = newBinaryNode(17)
				bt.root.left.right.right.parent = bt.root.left.right

				bt.root.left.right.right.right = newBinaryNode(19)
				bt.root.left.right.right.right.parent = bt.root.left.right.right

				return bt
			},
		},
		{
			name: "test rotate return error when tree is empty",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.RotateLeft(), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("tree is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeRotateLeftAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test rotate left at tree with one node",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.RotateLeftAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				err = bt.RotateLeftAt(2)

				return err, bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(1)

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.right.left = newBinaryNode(6)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.right = newBinaryNode(7)
				bt.root.right.right.parent = bt.root.right

				bt.root.left = newBinaryNode(5)
				bt.root.left.parent = bt.root

				bt.root.left.left = newBinaryNode(2)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(4)
				bt.root.left.left.left.parent = bt.root.left.left

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes two",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.RotateLeftAt(14), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(25)

				bt.root.left = newBinaryNode(17)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(20)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.left = newBinaryNode(15)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(14)
				bt.root.left.left.left.parent = bt.root.left.left

				bt.root.left.left.left.left = newBinaryNode(10)
				bt.root.left.left.left.left.parent = bt.root.left.left.left

				bt.root.left.left.left.left.left = newBinaryNode(7)
				bt.root.left.left.left.left.left.parent = bt.root.left.left.left.left

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes three",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.RotateLeftAt(15), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(10)

				bt.root.right = newBinaryNode(20)
				bt.root.right.parent = bt.root

				bt.root.right.right = newBinaryNode(25)
				bt.root.right.right.parent = bt.root.right

				bt.root.right.left = newBinaryNode(17)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.left.right = newBinaryNode(19)
				bt.root.right.left.right.parent = bt.root.right.left

				bt.root.right.left.left = newBinaryNode(15)
				bt.root.right.left.left.parent = bt.root.right.left

				bt.root.right.left.left.left = newBinaryNode(12)
				bt.root.right.left.left.left.parent = bt.root.right.left.left

				return bt
			},
		},
		{
			name: "test rotate return error when tree is empty",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.RotateLeftAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("tree is empty"),
		},
		{
			name: "test rotate return error when value is not found",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt.RotateLeftAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("1 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeRotateRightAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Tree)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test rotate left at tree with one node",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.RotateRightAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				err = bt.RotateRightAt(2)

				return err, bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(1)

				bt.root.right = newBinaryNode(3)
				bt.root.right.parent = bt.root

				bt.root.right.left = newBinaryNode(6)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.right = newBinaryNode(7)
				bt.root.right.right.parent = bt.root.right

				bt.root.left = newBinaryNode(4)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(2)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.right.right = newBinaryNode(5)
				bt.root.left.right.right.parent = bt.root.left.right

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes two",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.RotateRightAt(14), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(25)

				bt.root.left = newBinaryNode(17)
				bt.root.left.parent = bt.root

				bt.root.left.right = newBinaryNode(20)
				bt.root.left.right.parent = bt.root.left

				bt.root.left.left = newBinaryNode(10)
				bt.root.left.left.parent = bt.root.left

				bt.root.left.left.left = newBinaryNode(7)
				bt.root.left.left.left.parent = bt.root.left.left

				bt.root.left.left.right = newBinaryNode(14)
				bt.root.left.left.right.parent = bt.root.left.left

				bt.root.left.left.right.right = newBinaryNode(15)
				bt.root.left.left.right.right.parent = bt.root.left.left.right

				return bt
			},
		},
		{
			name: "test rotate tree with multiple nodes three",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.RotateRightAt(15), bt
			},
			expectedResult: func() Tree {
				bt := &BinaryTree{typeURL: "int", count: 7}

				bt.root = newBinaryNode(10)

				bt.root.right = newBinaryNode(20)
				bt.root.right.parent = bt.root

				bt.root.right.right = newBinaryNode(25)
				bt.root.right.right.parent = bt.root.right

				bt.root.right.left = newBinaryNode(12)
				bt.root.right.left.parent = bt.root.right

				bt.root.right.left.right = newBinaryNode(15)
				bt.root.right.left.right.parent = bt.root.right.left

				bt.root.right.left.right.right = newBinaryNode(17)
				bt.root.right.left.right.right.parent = bt.root.right.left.right

				bt.root.right.left.right.right.right = newBinaryNode(19)
				bt.root.right.left.right.right.right.parent = bt.root.right.left.right.right

				return bt
			},
		},
		{
			name: "test rotate return error when tree is empty",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt.RotateRightAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree()
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("tree is empty"),
		},
		{
			name: "test rotate return error when value is not found",
			actualResult: func() (error, Tree) {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt.RotateRightAt(1), bt
			},
			expectedResult: func() Tree {
				bt, err := NewBinaryTree(2)
				require.NoError(t, err)

				return bt
			},
			expectedError: errors.New("1 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinaryTreeIsFull(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test tree with one node is full",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.IsFull()
			},
			expectedResult: true,
		},
		{
			name: "test return true when all nodes have 2 or 0 children",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.IsFull()
			},
			expectedResult: true,
		},
		{
			name: "test return false when all any node has 1 child",
			actualResult: func() bool {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.IsFull()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeIsPerfect(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test tree with one node is perfect",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.IsPerfect()
			},
			expectedResult: true,
		},
		{
			name: "test return true when tree is perfect",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.IsPerfect()
			},
			expectedResult: true,
		},
		{
			name: "test return false when tree is not perfect",
			actualResult: func() bool {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.IsPerfect()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeIsComplete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test tree with one node is complete",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				return bt.IsComplete()
			},
			expectedResult: true,
		},
		{
			name: "test return true when tree is complete",
			actualResult: func() bool {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return bt.IsComplete()
			},
			expectedResult: true,
		},
		{
			name: "test return false when tree is not complete",
			actualResult: func() bool {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(16, c))

				return bt.IsComplete()
			},
			expectedResult: false,
		},
		{
			name: "test return false when tree is not complete two",
			actualResult: func() bool {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(9, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(5, c))
				require.NoError(t, bt.InsertCompare(7, c))

				return bt.IsComplete()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeLowestCommonAncestor(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (interface{}, error)
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test get lowest common ancestor",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.LowestCommonAncestor(4, 7)
			},
			expectedResult: 1,
		},
		{
			name: "test get lowest common ancestor two",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.LowestCommonAncestor(5, 3)
			},
			expectedResult: 1,
		},
		{
			name: "test get lowest common ancestor three",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.LowestCommonAncestor(6, 7)
			},
			expectedResult: 3,
		},
		{
			name: "test get lowest common ancestor four",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return bt.LowestCommonAncestor(3, 7)
			},
			expectedResult: 3,
		},
		{
			name: "test get lowest common ancestor five",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.LowestCommonAncestor(19, 25)
			},
			expectedResult: 20,
		},
		{
			name: "test get lowest common ancestor six",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.LowestCommonAncestor(19, 12)
			},
			expectedResult: 15,
		},
		{
			name: "test get lowest common return error when node is not present",
			actualResult: func() (interface{}, error) {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				return bt.LowestCommonAncestor(1, 12)
			},
			expectedError: errors.New("1 not found in the tree"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestBinaryTreeLeftViewIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test level order iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.LeftViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test level order iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.LeftViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 4},
		},
		{
			name: "test level order iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.LeftViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 7, 6},
		},
		{
			name: "test level order iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.LeftViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 20, 15, 12, 19},
		},
		{
			name: "test level order iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.LeftViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{25, 17, 14, 10, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestBinaryTreeRightViewIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test level order iterator when tree only contains of one node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				res := make([]interface{}, 0)

				it := bt.RightViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "test level order iterator when tree only contains multiple node",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(1)
				require.NoError(t, err)

				require.NoError(t, bt.Insert(2))
				require.NoError(t, bt.Insert(3))
				require.NoError(t, bt.Insert(4))
				require.NoError(t, bt.Insert(5))
				require.NoError(t, bt.Insert(6))
				require.NoError(t, bt.Insert(7))

				res := make([]interface{}, 0)

				it := bt.RightViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 3, 7},
		},
		{
			name: "test level order iterator when tree only contains multiple node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(7, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(6, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(8, c))
				require.NoError(t, bt.InsertCompare(16, c))

				res := make([]interface{}, 0)

				it := bt.RightViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 14, 16},
		},
		{
			name: "test level order iterator when tree only contains right node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(10)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(25, c))
				require.NoError(t, bt.InsertCompare(12, c))
				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(19, c))

				res := make([]interface{}, 0)

				it := bt.RightViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{10, 20, 25, 17, 19},
		},
		{
			name: "test level order iterator when tree only contains left node with compare",
			actualResult: func() []interface{} {
				bt, err := NewBinaryTree(25)
				require.NoError(t, err)

				c := comparator.NewIntegerComparator()

				require.NoError(t, bt.InsertCompare(17, c))
				require.NoError(t, bt.InsertCompare(20, c))
				require.NoError(t, bt.InsertCompare(14, c))
				require.NoError(t, bt.InsertCompare(10, c))
				require.NoError(t, bt.InsertCompare(15, c))
				require.NoError(t, bt.InsertCompare(7, c))

				res := make([]interface{}, 0)

				it := bt.RightViewIterator()

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{25, 17, 20, 15, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
