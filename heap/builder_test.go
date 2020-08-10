package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapBuilder(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (error, []interface{}, map[interface{}]int)
		expectedResult  []interface{}
		expectedIndexes map[interface{}]int
		expectedError   error
	}{
		{
			name: "test build max heap from one element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10},
			expectedIndexes: map[interface{}]int{10: 0},
		},
		{
			name: "test build max heap from two element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10, 40}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{40, 10},
			expectedIndexes: map[interface{}]int{40: 0, 10: 1},
		},
		{
			name: "test build max heap from arbitrary elements",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{0, 1, 2, 3}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{3, 1, 2, 0},
			expectedIndexes: map[interface{}]int{3: 0, 1: 1, 2: 2, 0: 3},
		},
		{
			name: "test build another max heap from arbitrary elements",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10, 80, 60, 5, 9, 45, 72, 85, 120}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{120, 85, 72, 80, 9, 45, 60, 10, 5},
			expectedIndexes: map[interface{}]int{120: 0, 85: 1, 72: 2, 80: 3, 9: 4, 45: 5, 60: 6, 10: 7, 5: 8},
		},
		{
			name: "test build return error when max heapify fails due to invalid comparator",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{30, 20, 10}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewStringComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{30, 20, 10},
			expectedIndexes: map[interface{}]int{},
			expectedError:   liberror.NewTypeMismatchError("string", "int"),
		},
		{
			name: "test build min heap from one element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10},
			expectedIndexes: map[interface{}]int{10: 0},
		},
		{
			name: "test build min heap from two element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10, 40}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10, 40},
			expectedIndexes: map[interface{}]int{10: 0, 40: 1},
		},
		{
			name: "test build min heap from arbitrary elements",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{3, 2, 1, 0}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{0, 2, 1, 3},
			expectedIndexes: map[interface{}]int{0: 0, 2: 1, 1: 2, 3: 3},
		},
		{
			name: "test build another min heap from arbitrary elements",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10, 1, 60, 5, 9, 45, 12, 2, 0}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{0, 1, 12, 2, 9, 45, 60, 10, 5},
			expectedIndexes: map[interface{}]int{0: 0, 1: 1, 12: 2, 2: 3, 9: 4, 45: 5, 60: 6, 10: 7, 5: 8},
		},
		{
			name: "test build return error when min heapify fails due to invalid comparator",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{10, 20, 30}
				indexes := make(map[interface{}]int)
				return buildHeap(comparator.NewStringComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10, 20, 30},
			expectedError:   liberror.NewTypeMismatchError("string", "int"),
			expectedIndexes: map[interface{}]int{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res, idx := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedIndexes, idx)
		})
	}
}
