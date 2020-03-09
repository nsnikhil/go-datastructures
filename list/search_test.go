package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

var elements = func(size int) []interface{} {
	var nums []interface{}
	for i := size; i >= 0; i-- {
		nums = append(nums, i)
	}
	return nums
}

var arrayList = func(t *testing.T, size int) *ArrayList {
	al, err := NewArrayList(elements(size)...)
	require.NoError(t, err)

	return al
}

var linkedList = func(t *testing.T, size int) *LinkedList {
	ll, err := NewLinkedList(elements(size)...)
	require.NoError(t, err)

	return ll
}

func TestGetFinder(t *testing.T) {
	testCases := []struct {
		name           string
		actualFinder   func() finder
		expectedFinder finder
	}{
		{
			name: "test get linear finder",
			actualFinder: func() finder {
				return newFinder(linear)
			},
			expectedFinder: newLinearFinder(),
		},
		{
			name: "test get concurrent finder",
			actualFinder: func() finder {
				return newFinder(concurrent)
			},
			expectedFinder: newConcurrentFinder(),
		},
		{
			name: "test get doubly finder",
			actualFinder: func() finder {
				return newFinder(doubly)
			},
			expectedFinder: newDoublyFinder(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedFinder, testCase.actualFinder())
		})
	}
}

func TestFinderLinearSearch(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index when element is found",
			actualResult: func() (int, error) {
				return newLinearFinder().search(arrayList(t, 100), 4)
			},
			expectedResult: 96,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() (int, error) {
				return newLinearFinder().search(arrayList(t, 100), 105)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 105 not found in the list"),
		},
		{
			name: "return -1 with error when list is empty",
			actualResult: func() (int, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return newLinearFinder().search(al, 1)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return -1 with error when element is of different type",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return newLinearFinder().search(al, "a")
			},
			expectedResult: -1,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestFinderConcurrentSearch(t *testing.T) {
	sz := int(math.Pow(10, float64(6)))

	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index when element is found at start",
			actualResult: func() (int, error) {
				return newConcurrentFinder().search(arrayList(t, sz), 11000)
			},
			expectedResult: 989000,
		},
		{
			name: "return index when element is found at end",
			actualResult: func() (int, error) {
				return newConcurrentFinder().search(arrayList(t, sz), 0)
			},
			expectedResult: 1000000,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() (int, error) {
				return newConcurrentFinder().search(arrayList(t, sz), 1000005)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 1000005 not found in the list"),
		},
		{
			name: "return -1 with error when list is empty",
			actualResult: func() (int, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return newConcurrentFinder().search(al, 1)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return -1 with error when element is of different type",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return newConcurrentFinder().search(al, "a")
			},
			expectedResult: -1,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestFinderDoublySearch(t *testing.T) {

	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index when element is found at start",
			actualResult: func() (int, error) {
				return newDoublyFinder().search(linkedList(t, 100), 100)
			},
			expectedResult: 0,
		},
		{
			name: "return index when element is found at end",
			actualResult: func() (int, error) {
				return newDoublyFinder().search(linkedList(t, 100), 0)
			},
			expectedResult: 100,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() (int, error) {
				return newDoublyFinder().search(linkedList(t, 100), 105)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 105 not found in the list"),
		},
		{
			name: "return -1 with error when list is empty",
			actualResult: func() (int, error) {
				al, err := NewLinkedList()
				require.NoError(t, err)

				return newDoublyFinder().search(al, 1)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return -1 with error when element is of different type",
			actualResult: func() (int, error) {
				al, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return newDoublyFinder().search(al, "a")
			},
			expectedResult: -1,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func BenchmarkFinders(b *testing.B) {
	var lists []List

	for i := 1; i <= 8; i++ {

		sz := int(math.Pow(10, float64(i)))
		var ele []interface{}
		for j := 0; j < sz; j++ {
			ele = append(ele, j)
		}

		al, _ := NewArrayList(ele...)
		lists = append(lists, al)

	}

	lf := newLinearFinder()

	cf := newConcurrentFinder()

	for i, l := range lists {
		b.Run(fmt.Sprintf("lf %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				lf.search(l, math.MinInt32)
			}
		})
	}

	for i, l := range lists {
		b.Run(fmt.Sprintf("cf %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				cf.search(l, math.MinInt32)
			}
		})
	}
}
