package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFinderDoublySearch(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "return index when element is found at end",
			actualResult: func() int64 {
				return newDoublyFinder[int64]().search(newTestLinkedList(101), 100)
			},
			expectedResult: 100,
		},
		{
			name: "return index when element is found at start",
			actualResult: func() int64 {
				return newDoublyFinder[int64]().search(newTestLinkedList(100), 0)
			},
			expectedResult: 0,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() int64 {
				return newDoublyFinder[int64]().search(newTestLinkedList(100), 105)
			},
			expectedResult: -1,
		},
		{
			name: "return -1 with error when list is empty",
			actualResult: func() int64 {
				return newDoublyFinder[int64]().search(newTestLinkedList(0), 1)
			},
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
