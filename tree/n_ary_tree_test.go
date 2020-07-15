package tree

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewNaryTree(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Tree, error)
		expectedResult func() Tree
		expectedError  error
	}{
		{
			name: "test create empty nary tree",
			actualResult: func() (Tree, error) {
				return NewNAryTree(1)
			},
			expectedResult: func() Tree {
				return &NAryTree{
					count:       0,
					typeURL:     "na",
					maxChildren: 1,
				}
			},
		},
		{
			name: "test nary tree with values",
			actualResult: func() (Tree, error) {
				return NewNAryTree(2, 1, 2, 3, 4)
			},
			expectedResult: func() Tree {
				nt := &NAryTree{
					count:       4,
					typeURL:     "int",
					maxChildren: 2,
				}

				a := newNode(1)
				b := newNode(2)
				c := newNode(3)
				d := newNode(4)

				nt.root = a

				require.NoError(t, a.add(b))
				b.parent = a

				require.NoError(t, a.add(c))
				c.parent = a

				require.NoError(t, b.add(d))
				d.parent = b

				return nt
			},
		},
		{
			name: "test nary tree with values and child count 3",
			actualResult: func() (Tree, error) {
				return NewNAryTree(3, 1, 2, 3, 4)
			},
			expectedResult: func() Tree {
				nt := &NAryTree{
					count:       4,
					typeURL:     "int",
					maxChildren: 3,
				}

				a := newNode(1)
				b := newNode(2)
				c := newNode(3)
				d := newNode(4)

				nt.root = a

				require.NoError(t, a.add(b))
				b.parent = a

				require.NoError(t, a.add(c))
				c.parent = a

				require.NoError(t, a.add(d))
				d.parent = a

				return nt
			},
		},
		{
			name: "test failed to create nary tree due to type mismatch",
			actualResult: func() (Tree, error) {
				return NewNAryTree(2, 1, 2, 'a')
			},
			expectedResult: func() Tree {
				return (*NAryTree)(nil)
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
