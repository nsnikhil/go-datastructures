package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapify(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (error, []interface{}, map[interface{}]int)
		expectedResult  []interface{}
		expectedIndexes map[interface{}]int
		expectedError   error
	}{
		{
			name: "max heapify for one element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{100}
				indexes := map[interface{}]int{100: 0}
				return heapify(0, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{100},
			expectedIndexes: map[interface{}]int{100: 0},
		},
		{
			name: "max heapify 1st element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{100, 140, 60}
				indexes := map[interface{}]int{100: 0, 140: 1, 60: 2}
				return heapify(0, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{140, 100, 60},
			expectedIndexes: map[interface{}]int{140: 0, 100: 1, 60: 2},
		},
		{
			name: "min heapify 1st element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{100, 40, 160}
				indexes := map[interface{}]int{100: 0, 40: 1, 160: 2}
				return heapify(0, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{40, 100, 160},
			expectedIndexes: map[interface{}]int{40: 0, 100: 1, 160: 2},
		},
		{
			name: "max heapify fourth element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{9, 8, 5, 6, 10}
				indexes := map[interface{}]int{9: 0, 8: 1, 5: 2, 6: 3, 10: 4}
				return heapify(4, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10, 9, 5, 6, 8},
			expectedIndexes: map[interface{}]int{10: 0, 9: 1, 5: 2, 6: 3, 8: 4},
		},
		{
			name: "min heapify for one element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{1}
				indexes := map[interface{}]int{1: 0}
				return heapify(0, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{1},
			expectedIndexes: map[interface{}]int{1: 0},
		},
		{
			name: "min heapify fourth element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{1, 2, 3, 4, 0}
				indexes := map[interface{}]int{1: 0, 2: 1, 3: 2, 4: 3, 0: 4}
				return heapify(4, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{0, 1, 3, 4, 2},
			expectedIndexes: map[interface{}]int{0: 0, 1: 1, 3: 2, 4: 3, 2: 4},
		},
		{
			name: "max heapify 2nd element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{7, 0, 4, 1, 2}
				indexes := map[interface{}]int{7: 0, 0: 1, 4: 2, 1: 3, 2: 4}
				return heapify(1, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{7, 2, 4, 1, 0},
			expectedIndexes: map[interface{}]int{7: 0, 2: 1, 4: 2, 1: 3, 0: 4},
		},
		{
			name: "min heapify 2nd element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{5, 7, 8, 9, 6}
				indexes := map[interface{}]int{5: 0, 7: 1, 8: 2, 9: 3, 6: 4}
				return heapify(1, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{5, 6, 8, 9, 7},
			expectedIndexes: map[interface{}]int{5: 0, 6: 1, 8: 2, 9: 3, 7: 4},
		},
		{
			name: "max heapify last element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{9, 8, 5, 6, 10}
				indexes := map[interface{}]int{9: 0, 8: 1, 5: 2, 6: 3, 10: 4}
				return heapify(4, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{10, 9, 5, 6, 8},
			expectedIndexes: map[interface{}]int{10: 0, 9: 1, 5: 2, 6: 3, 8: 4},
		},
		{
			name: "max heapify last element two",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{5, 4, 3, 6}
				indexes := map[interface{}]int{5: 0, 4: 4, 3: 2, 6: 3}
				return heapify(3, comparator.NewIntegerComparator(), true, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{6, 5, 3, 4},
			expectedIndexes: map[interface{}]int{6: 0, 5: 4, 3: 2, 4: 3},
		},
		{
			name: "min heapify last element",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{1, 2, 3, 4, 0}
				indexes := map[interface{}]int{1: 0, 2: 1, 3: 2, 4: 3, 0: 4}
				return heapify(4, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{0, 1, 3, 4, 2},
			expectedIndexes: map[interface{}]int{0: 0, 1: 1, 3: 2, 4: 3, 2: 4},
		},
		{
			name: "min heapify last element two",
			actualResult: func() (error, []interface{}, map[interface{}]int) {
				data := []interface{}{3, 4, 5, 2}
				indexes := map[interface{}]int{3: 0, 4: 1, 5: 2, 2: 3}
				return heapify(3, comparator.NewIntegerComparator(), false, data, indexes), data, indexes
			},
			expectedResult:  []interface{}{2, 3, 5, 4},
			expectedIndexes: map[interface{}]int{2: 0, 3: 1, 5: 2, 4: 3},
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
