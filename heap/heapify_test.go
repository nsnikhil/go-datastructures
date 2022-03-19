package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapify(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() []int
		expectedResult  []int
		expectedIndexes map[int]int
		expectedError   error
	}{
		{
			name: "should max heapify for one element",
			actualResult: func() []int {
				data := []int{100}
				heapify(0, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{100},
			expectedIndexes: map[int]int{100: 0},
		},
		{
			name: "should max heapify 1st element",
			actualResult: func() []int {
				data := []int{100, 140, 60}
				heapify(0, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{140, 100, 60},
			expectedIndexes: map[int]int{140: 0, 100: 1, 60: 2},
		},
		{
			name: "should min heapify 1st element",
			actualResult: func() []int {
				data := []int{100, 40, 160}
				heapify(0, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{40, 100, 160},
			expectedIndexes: map[int]int{40: 0, 100: 1, 160: 2},
		},
		{
			name: "should max heapify fourth element",
			actualResult: func() []int {
				data := []int{9, 8, 5, 6, 10}
				heapify(4, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{10, 9, 5, 6, 8},
			expectedIndexes: map[int]int{10: 0, 9: 1, 5: 2, 6: 3, 8: 4},
		},
		{
			name: "should min heapify for one element",
			actualResult: func() []int {
				data := []int{1}
				heapify(0, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{1},
			expectedIndexes: map[int]int{1: 0},
		},
		{
			name: "should min heapify fourth element",
			actualResult: func() []int {
				data := []int{1, 2, 3, 4, 0}
				heapify(4, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{0, 1, 3, 4, 2},
			expectedIndexes: map[int]int{0: 0, 1: 1, 3: 2, 4: 3, 2: 4},
		},
		{
			name: "should max heapify 2nd element",
			actualResult: func() []int {
				data := []int{7, 0, 4, 1, 2}
				heapify(1, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{7, 2, 4, 1, 0},
			expectedIndexes: map[int]int{7: 0, 2: 1, 4: 2, 1: 3, 0: 4},
		},
		{
			name: "should min heapify 2nd element",
			actualResult: func() []int {
				data := []int{5, 7, 8, 9, 6}
				heapify(1, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{5, 6, 8, 9, 7},
			expectedIndexes: map[int]int{5: 0, 6: 1, 8: 2, 9: 3, 7: 4},
		},
		{
			name: "should max heapify last element",
			actualResult: func() []int {
				data := []int{9, 8, 5, 6, 10}
				heapify(4, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{10, 9, 5, 6, 8},
			expectedIndexes: map[int]int{10: 0, 9: 1, 5: 2, 6: 3, 8: 4},
		},
		{
			name: "should max heapify last element two",
			actualResult: func() []int {
				data := []int{5, 4, 3, 6}
				heapify(3, comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{6, 5, 3, 4},
			expectedIndexes: map[int]int{6: 0, 5: 4, 3: 2, 4: 3},
		},
		{
			name: "should min heapify last element",
			actualResult: func() []int {
				data := []int{1, 2, 3, 4, 0}
				heapify(4, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{0, 1, 3, 4, 2},
			expectedIndexes: map[int]int{0: 0, 1: 1, 3: 2, 4: 3, 2: 4},
		},
		{
			name: "should min heapify last element two",
			actualResult: func() []int {
				data := []int{3, 4, 5, 2}
				heapify(3, comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{2, 3, 5, 4},
			expectedIndexes: map[int]int{2: 0, 3: 1, 5: 2, 4: 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
