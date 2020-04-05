package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapify(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, []interface{})
		expectedResult []interface{}
		expectedError  error
	}{
		{
			name: "max heapify for one element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{100}
				return heapify(0, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{100},
		},
		{
			name: "max heapify 1st element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{100, 140, 60}
				return heapify(0, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{140, 100, 60},
		},
		{
			name: "min heapify 1st element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{100, 40, 160}
				return heapify(0, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{40, 100, 160},
		},
		{
			name: "max heapify fourth element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{9, 8, 5, 6, 10}
				return heapify(4, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{10, 9, 5, 6, 8},
		},
		{
			name: "min heapify for one element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{1}
				return heapify(0, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{1},
		},
		{
			name: "min heapify fourth element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{1, 2, 3, 4, 0}
				return heapify(4, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{0, 1, 3, 4, 2},
		},
		{
			name: "max heapify 2nd element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{7, 0, 4, 1, 2}
				return heapify(1, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{7, 2, 4, 1, 0},
		},
		{
			name: "min heapify 2nd element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{5, 7, 8, 9, 6}
				return heapify(1, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{5, 6, 8, 9, 7},
		},
		{
			name: "max heapify last element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{9, 8, 5, 6, 10}
				return heapify(4, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{10, 9, 5, 6, 8},
		},
		{
			name: "max heapify last element two",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{5, 4, 3, 6}
				return heapify(3, comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{6, 5, 3, 4},
		},
		{
			name: "min heapify last element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{1, 2, 3, 4, 0}
				return heapify(4, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{0, 1, 3, 4, 2},
		},
		{
			name: "min heapify last element two",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{3, 4, 5, 2}
				return heapify(3, comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{2, 3, 5, 4},
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
