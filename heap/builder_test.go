package heap

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapBuilder(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, []interface{})
		expectedResult []interface{}
		expectedError  error
	}{
		{
			name: "test build max heap from one element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10}
				return buildHeap(comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{10},
		},
		{
			name: "test build max heap from two element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10, 40}
				return buildHeap(comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{40, 10},
		},
		{
			name: "test build max heap from arbitrary elements",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{0, 1, 2, 3}
				return buildHeap(comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{3, 1, 2, 0},
		},
		{
			name: "test build another max heap from arbitrary elements",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10, 80, 60, 5, 9, 45, 72, 85, 120}
				return buildHeap(comparator.NewIntegerComparator(), true, data), data
			},
			expectedResult: []interface{}{120, 85, 72, 80, 9, 45, 60, 10, 5},
		},
		{
			name: "test build return error when max heapify fails due to invalid comparator",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{30, 20, 10}
				return buildHeap(comparator.NewStringComparator(), true, data), data
			},
			expectedResult: []interface{}{30, 20, 10},
			expectedError:  liberror.NewTypeMismatchError("string", "int"),
		},
		{
			name: "test build min heap from one element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10}
				return buildHeap(comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{10},
		},
		{
			name: "test build min heap from two element",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10, 40}
				return buildHeap(comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{10, 40},
		},
		{
			name: "test build min heap from arbitrary elements",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{3, 2, 1, 0}
				return buildHeap(comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{0, 2, 1, 3},
		},
		{
			name: "test build another min heap from arbitrary elements",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10, 1, 60, 5, 9, 45, 12, 2, 0}
				return buildHeap(comparator.NewIntegerComparator(), false, data), data
			},
			expectedResult: []interface{}{0, 1, 12, 2, 9, 45, 60, 10, 5},
		},
		{
			name: "test build return error when min heapify fails due to invalid comparator",
			actualResult: func() (error, []interface{}) {
				data := []interface{}{10, 20, 30}
				return buildHeap(comparator.NewStringComparator(), false, data), data
			},
			expectedResult: []interface{}{10, 20, 30},
			expectedError:  liberror.NewTypeMismatchError("string", "int"),
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
