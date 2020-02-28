package list

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type testObj struct {
	elements []int
}

func newTestObj(elements ...int) testObj {
	return testObj{
		elements: elements,
	}
}

func (to testObj) sum() int {
	total := 0
	for _, e := range to.elements {
		total += e
	}
	return total
}

type testObjComparator struct{}

func (testObjComparator) Compare(one interface{}, two interface{}) (int, error) {
	return (one).(testObj).sum() - (two).(testObj).sum(), nil
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult List
	}{
		{
			name: "sort integer list",
			actualResult: func() List {
				al, err := NewArrayList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				newMergeSorter().sort(al, comparator.NewIntegerComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "sort string list",
			actualResult: func() List {
				al, err := NewArrayList("e", "d", "c", "b", "a")
				require.NoError(t, err)

				newMergeSorter().sort(al, comparator.NewStringComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "b", "c", "d", "e"},
			},
		},
		{
			name: "sort object list",
			actualResult: func() List {
				al, err := NewArrayList(newTestObj(2, 3), newTestObj(4, 6), newTestObj(1, 4))
				require.NoError(t, err)

				newMergeSorter().sort(al, testObjComparator{})

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "testObj",
				data:    []interface{}{newTestObj(1, 4), newTestObj(2, 3), newTestObj(4, 6)},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult List
	}{
		{
			name: "sort integer list",
			actualResult: func() List {
				al, err := NewArrayList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				newQuickSorter().sort(al, comparator.NewIntegerComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "sort string list",
			actualResult: func() List {
				al, err := NewArrayList("e", "d", "c", "b", "a")
				require.NoError(t, err)

				newQuickSorter().sort(al, comparator.NewStringComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "b", "c", "d", "e"},
			},
		},
		{
			name: "sort object list",
			actualResult: func() List {
				al, err := NewArrayList(newTestObj(2, 3), newTestObj(4, 6), newTestObj(1, 4))
				require.NoError(t, err)

				newQuickSorter().sort(al, testObjComparator{})

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "testObj",
				data:    []interface{}{newTestObj(1, 4), newTestObj(2, 3), newTestObj(4, 6)},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
