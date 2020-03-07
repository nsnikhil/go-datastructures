package comparator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPrebuiltComparator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult func(res int) bool
		expectedError  error
	}{
		{
			name: "test integer comparator for greater",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, 1)
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test integer comparator for smaller",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(1, 2)
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test integer comparator for equals",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, 2)
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test integer comparator return error for type mismatch of first argument",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare("invalid", 2)
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected int got string"),
		},
		{
			name: "test integer comparator return error for type mismatch of second argument",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, "invalid")
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected int got string"),
		},
		{
			name: "test string comparator for greater",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("ab", "a")
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test string comparator for smaller",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("a", "ab")
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test string comparator for equals",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("a", "a")
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test string comparator return error for type for first mismatch",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare(1, "ab")
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected string got int"),
		},
		{
			name: "test string comparator return error for type for second mismatch",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("ab", 1)
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected string got int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.True(t, testCase.expectedResult(res))
		})
	}
}
