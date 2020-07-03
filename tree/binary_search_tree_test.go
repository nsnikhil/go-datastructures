package tree

import (
	"testing"
)

func TestCreateNewBinarySearchTree(t *testing.T)  {
	testCases := []struct {
		name           string
		actualResult   func() (Tree, error)
		expectedResult func() Tree
		expectedError  error
	}{
		{},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//res, err := testCase.actualResult()
			//
			//assert.Equal(t, testCase.expectedError, err)
			//assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestBinarySearchTreeInsert(t *testing.T)  {

}

func TestBinarySearchTreeDelete(t *testing.T)  {

}
