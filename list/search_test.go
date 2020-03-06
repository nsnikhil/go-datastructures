package list

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func getList() List {
	nums := func() []interface{} {
		var nums []interface{}
		for i := 100; i >= 0; i-- {
			nums = append(nums, i)
		}
		return nums
	}

	al, _ := NewArrayList(nums()...)
	return al
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
				return newLinearFinder().search(getList(), 4)
			},
			expectedResult: 96,
		},
		{
			name: "return -1 with error when element is not found",
			actualResult: func() (int, error) {
				return newLinearFinder().search(getList(), 105)
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

func generateList(size int) List {
	nums := func() []interface{} {
		var nums []interface{}
		for i := size; i >= 0; i-- {
			nums = append(nums, i)
		}
		return nums
	}

	al, _ := NewArrayList(nums()...)
	return al
}

func BenchmarkLinearFinder(b *testing.B) {
	s := b.N
	l := generateList(s)
	f := newLinearFinder()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < s; i++ {
		e := rand.Intn(s-0+1) + 0
		_, _ = f.search(l, e)
	}
}
