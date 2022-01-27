package comparator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrebuiltComparator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult func(res int) bool
		expectedError  error
	}{
		{
			name: "test integer comparator for greater",
			actualResult: func() int {
				ic := NewIntegerComparator()
				return ic.Compare(2, 1)
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test integer comparator for smaller",
			actualResult: func() int {
				ic := NewIntegerComparator()
				return ic.Compare(1, 2)
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test integer comparator for equals",
			actualResult: func() int {
				ic := NewIntegerComparator()
				return ic.Compare(2, 2)
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test string comparator for greater",
			actualResult: func() int {
				ic := NewStringComparator()
				return ic.Compare("ab", "a")
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test string comparator for smaller",
			actualResult: func() int {
				ic := NewStringComparator()
				return ic.Compare("a", "ab")
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test string comparator for equals",
			actualResult: func() int {
				ic := NewStringComparator()
				return ic.Compare("a", "a")
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.True(t, testCase.expectedResult(res))
		})
	}
}
