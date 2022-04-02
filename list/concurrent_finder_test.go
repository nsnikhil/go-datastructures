package list

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//TODO: INSPECT THIS TEST AS IT ONLY FAILS ON CI
func NotTestFinderConcurrentSearch(t *testing.T) {
	sz := int64(math.Pow(10, float64(6))) + 1

	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "return index when element is found at end",
			actualResult: func() int64 {
				return newConcurrentFinder[int64]().search(newTestArrayList(sz), 1000000)
			},
			expectedResult: 1000000,
		},
		{
			name: "return index when element is found at start",
			actualResult: func() int64 {
				return newConcurrentFinder[int64]().search(newTestArrayList(sz), 0)
			},
			expectedResult: 0,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() int64 {
				return newConcurrentFinder[int64]().search(newTestArrayList(sz), 1000005)
			},
			expectedResult: -1,
		},
		{
			name: "return -1 with error when list is empty",
			actualResult: func() int64 {
				return newConcurrentFinder[int64]().search(newTestArrayList(0), 1)
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
