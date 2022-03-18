package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, []int)
		expectedResult []int
		expectedError  error
	}{
		{
			name: "test sort integer array descending",
			actualResult: func() (error, []int) {
				hs := newHeapSort[int]()
				data := []int{8, 4, 7, 2, 9, 0, 1, 3, 5, 6}

				err := hs.sort(comparator.NewIntegerComparator(), true, &data)

				return err, data
			},
			expectedResult: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "test sort integer array ascending",
			actualResult: func() (error, []int) {
				hs := newHeapSort[int]()
				data := []int{8, 4, 7, 2, 9, 0, 1, 3, 5, 6}

				err := hs.sort(comparator.NewIntegerComparator(), false, &data)

				return err, data
			},
			expectedResult: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
