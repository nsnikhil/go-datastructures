package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapBuilder(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() []int
		expectedResult  []int
		expectedIndexes map[int]int
	}{
		{
			name: "test build max heap from one element",
			actualResult: func() []int {
				data := []int{10}
				buildHeap[int](comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{10},
			expectedIndexes: map[int]int{10: 0},
		},
		{
			name: "test build max heap from two element",
			actualResult: func() []int {
				data := []int{10, 40}
				buildHeap[int](comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{40, 10},
			expectedIndexes: map[int]int{40: 0, 10: 1},
		},
		{
			name: "test build max heap from arbitrary elements",
			actualResult: func() []int {
				data := []int{0, 1, 2, 3}
				buildHeap[int](comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{3, 1, 2, 0},
			expectedIndexes: map[int]int{3: 0, 1: 1, 2: 2, 0: 3},
		},
		{
			name: "test build another max heap from arbitrary elements",
			actualResult: func() []int {
				data := []int{10, 80, 60, 5, 9, 45, 72, 85, 120}
				buildHeap[int](comparator.NewIntegerComparator(), true, data)
				return data
			},
			expectedResult:  []int{120, 85, 72, 80, 9, 45, 60, 10, 5},
			expectedIndexes: map[int]int{120: 0, 85: 1, 72: 2, 80: 3, 9: 4, 45: 5, 60: 6, 10: 7, 5: 8},
		},
		{
			name: "test build min heap from one element",
			actualResult: func() []int {
				data := []int{10}
				buildHeap[int](comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{10},
			expectedIndexes: map[int]int{10: 0},
		},
		{
			name: "test build min heap from two element",
			actualResult: func() []int {
				data := []int{10, 40}
				buildHeap[int](comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{10, 40},
			expectedIndexes: map[int]int{10: 0, 40: 1},
		},
		{
			name: "test build min heap from arbitrary elements",
			actualResult: func() []int {
				data := []int{3, 2, 1, 0}
				buildHeap[int](comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{0, 2, 1, 3},
			expectedIndexes: map[int]int{0: 0, 2: 1, 1: 2, 3: 3},
		},
		{
			name: "test build another min heap from arbitrary elements",
			actualResult: func() []int {
				data := []int{10, 1, 60, 5, 9, 45, 12, 2, 0}
				buildHeap[int](comparator.NewIntegerComparator(), false, data)
				return data
			},
			expectedResult:  []int{0, 1, 12, 2, 9, 45, 60, 10, 5},
			expectedIndexes: map[int]int{0: 0, 1: 1, 12: 2, 2: 3, 9: 4, 45: 5, 60: 6, 10: 7, 5: 8},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
