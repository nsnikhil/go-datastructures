package list

import (
	"datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult List
	}{
		{
			name: "test Sort integer List",
			actualResult: func() List {
				al, err := NewArrayList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				sortList(al, comparator.NewIntegerComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func addElement(t *testing.T, l List, n int) {
	for i := 0; i < n; i++ {
		require.NoError(t, l.Add(i))
	}
}

func TestSearchList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test linear search return index when element is present",
			actualResult: func() int {
				al, err := NewArrayList()
				require.NoError(t, err)

				addElement(t, al, 100)

				return searchList(al, 85)

			},
			expectedResult: 85,
		},
		{
			name: "test linear search return -1 when element is not present",
			actualResult: func() int {
				al, err := NewArrayList()
				require.NoError(t, err)

				addElement(t, al, 100)

				return searchList(al, 111)

			},
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func BenchMarkSearchList(b *testing.B) {

}
