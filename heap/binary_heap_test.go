package heap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCreateMaxHeap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (*MaxHeap, error)
		expectedResult *MaxHeap
		expectedError  error
	}{
		{
			name: "test create empty max heap",
			actualResult: func() (*MaxHeap, error) {
				return NewMaxHeap(comparator.NewIntegerComparator())
			},
			expectedResult: &MaxHeap{&binaryHeap{
				isMaxHeap: true,
				typeURL:   "na",
				c:         comparator.NewIntegerComparator(),
				indexes:   make(map[interface{}]int),
			}},
		},
		{
			name: "test create heap of one element",
			actualResult: func() (*MaxHeap, error) {
				return NewMaxHeap(comparator.NewIntegerComparator(), 1)
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
		},
		{
			name: "test create max heap with multiple elements",
			actualResult: func() (*MaxHeap, error) {
				return NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{4, 2, 3, 1},
				indexes:   map[interface{}]int{4: 0, 2: 1, 3: 2, 1: 3},
			}},
		},
		{
			name: "test create return error when type of elements are not same",
			actualResult: func() (*MaxHeap, error) {
				return NewMaxHeap(comparator.NewIntegerComparator(), 1, "a")
			},
			expectedResult: (*MaxHeap)(nil),
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test create return error when comparator return error",
			actualResult: func() (*MaxHeap, error) {
				return NewMaxHeap(comparator.NewStringComparator(), 1, 2)
			},
			expectedResult: (*MaxHeap)(nil),
			expectedError:  liberror.NewTypeMismatchError("string", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestCreateMinHeap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (*MinHeap, error)
		expectedResult *MinHeap
		expectedError  error
	}{
		{
			name: "test create empty min heap",
			actualResult: func() (*MinHeap, error) {
				return NewMinHeap(comparator.NewIntegerComparator())
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "na",
				c:       comparator.NewIntegerComparator(),
				indexes: make(map[interface{}]int),
			}},
		},
		{
			name: "test create min heap with one element",
			actualResult: func() (*MinHeap, error) {
				return NewMinHeap(comparator.NewIntegerComparator(), 1)
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{1},
				indexes: map[interface{}]int{1: 0},
			}},
		},
		{
			name: "test create heap with many elements",
			actualResult: func() (*MinHeap, error) {
				return NewMinHeap(comparator.NewIntegerComparator(), 4, 3, 2, 1)
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{1, 3, 2, 4},
				indexes: map[interface{}]int{1: 0, 3: 1, 2: 2, 4: 3},
			}},
		},
		{
			name: "test create return error when type of elements are not same",
			actualResult: func() (*MinHeap, error) {
				return NewMinHeap(comparator.NewIntegerComparator(), 1, "a")
			},
			expectedResult: (*MinHeap)(nil),
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test create return error when when comparator return error",
			actualResult: func() (*MinHeap, error) {
				return NewMinHeap(comparator.NewStringComparator(), 1, 2)
			},
			expectedResult: (*MinHeap)(nil),
			expectedError:  liberror.NewTypeMismatchError("string", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMaxHeapAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MaxHeap)
		expectedResult *MaxHeap
		expectedError  error
	}{
		{
			name: "test heap add one element",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(10))
				return nil, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{10},
				indexes:   map[interface{}]int{10: 0},
			}},
		},
		{
			name: "test add will heapify one element",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(100, 40, 60, 80))
				return nil, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{100, 80, 60, 40},
				indexes:   map[interface{}]int{100: 0, 80: 1, 60: 2, 40: 3},
			}},
		},
		{
			name: "test add will heapify two element",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(100, 40, 60, 120))
				return nil, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{120, 100, 60, 40},
				indexes:   map[interface{}]int{120: 0, 100: 1, 60: 2, 40: 3},
			}},
		},
		{
			name: "test add will heapify two element two",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(100, 110, 120))
				return nil, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{120, 100, 110},
				indexes:   map[interface{}]int{120: 0, 100: 1, 110: 2},
			}},
		},
		{
			name: "test add will heapify all element",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
				return nil, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{9, 8, 5, 6, 7, 1, 4, 0, 3, 2},
				indexes:   map[interface{}]int{9: 0, 8: 1, 5: 2, 6: 3, 7: 4, 1: 5, 4: 6, 0: 7, 3: 8, 2: 9},
			}},
		},
		{
			name: "test add returns error when type is different",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Add(1, "a"), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "na",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   make(map[interface{}]int),
			}},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add return error when adding different type element to cleared list",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				h.Clear()

				return h.Add("a"), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   map[interface{}]int{2: 0, 1: 1},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add return error when comparator returns error",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewStringComparator())
				require.NoError(t, err)

				return h.Add(1, 2, 3), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewStringComparator(),
				data:      []interface{}{1, 2},
				indexes:   map[interface{}]int{1: 0, 2: 1},
			}},
			expectedError: liberror.NewTypeMismatchError("string", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMinHeapAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MinHeap)
		expectedResult *MinHeap
		expectedError  error
	}{
		{
			name: "test heap add one element",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(10))
				return nil, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{10},
				indexes: map[interface{}]int{10: 0},
			}},
		},
		{
			name: "test add will heapify one element",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(10, 40, 60, 20))
				return nil, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{10, 20, 60, 40},
				indexes: map[interface{}]int{10: 0, 20: 1, 60: 2, 40: 3},
			}},
		},
		{
			name: "test add will heapify two element",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(20, 40, 60, 10))
				return nil, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{10, 20, 60, 40},
				indexes: map[interface{}]int{10: 0, 20: 1, 60: 2, 40: 3},
			}},
		},
		{
			name: "test add will heapify two element two",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(30, 20, 10))
				return nil, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{10, 30, 20},
				indexes: map[interface{}]int{10: 0, 30: 1, 20: 2},
			}},
		},
		{
			name: "test add will heapify all element",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, h.Add(9, 8, 7, 6, 5, 4, 3, 2, 1, 0))
				return nil, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{0, 1, 4, 3, 2, 8, 5, 9, 6, 7},
				indexes: map[interface{}]int{0: 0, 1: 1, 4: 2, 3: 3, 2: 4, 8: 5, 5: 6, 9: 7, 6: 8, 7: 9},
			}},
		},
		{
			name: "test add returns error when type is different",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Add(1, "a"), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "na",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}(nil),
				indexes: make(map[interface{}]int),
			}},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add return error when adding different type element to cleared list",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				h.Clear()

				return h.Add("a"), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}(nil),
				indexes: map[interface{}]int{1: 0, 2: 1},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add return error when comparator returns error",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewStringComparator())
				require.NoError(t, err)

				return h.Add(1, 2, 3), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewStringComparator(),
				data:    []interface{}{1, 2},
				indexes: map[interface{}]int{1: 0, 2: 1},
			}},
			expectedError: liberror.NewTypeMismatchError("string", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMaxHeapIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return false when Heap is not empty",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				return h.IsEmpty()
			},
		},
		{
			name: "return true when Heap is empty",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.IsEmpty()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMinHeapIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return false when Heap is not empty",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				return h.IsEmpty()
			},
		},
		{
			name: "return true when Heap is empty",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.IsEmpty()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMaxHeapClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return true when Heap is empty after Clear",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return true when Heap is empty after Clear two",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewStringComparator(), "a", "b")
				require.NoError(t, err)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMinHeapClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return true when Heap is empty after Clear",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return true when Heap is empty after Clear two",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewStringComparator(), "a", "b")
				require.NoError(t, err)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMaxHeapSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "get Size of empty Heap as 0",
			actualResult: func() int {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Size()
			},
		},
		{
			name: "get Size of empty Heap as 2",
			actualResult: func() int {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				return h.Size()
			},
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMinHeapSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "get Size of empty Heap as 0",
			actualResult: func() int {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Size()
			},
		},
		{
			name: "get Size of empty Heap as 2",
			actualResult: func() int {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				return h.Size()
			},
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMaxHeapExtract(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, *MaxHeap)
		expectedElement interface{}
		expectedResult  *MaxHeap
		expectedError   error
	}{
		{
			name: "extract first element of the max heap",
			actualResult: func() (interface{}, error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 2,
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
		},
		{
			name: "extract first element of the max heap two",
			actualResult: func() (interface{}, error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 10,
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{9, 7, 8},
				indexes:   map[interface{}]int{9: 0, 7: 1, 8: 2},
			}},
		},
		{
			name: "extract first element of the max heap three",
			actualResult: func() (interface{}, error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 9, 8, 5, 6, 7, 1, 4, 0, 3, 2)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 9,
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{8, 7, 5, 6, 2, 1, 4, 0, 3},
				indexes:   map[interface{}]int{8: 0, 7: 1, 5: 2, 6: 3, 2: 4, 1: 5, 4: 6, 0: 7, 3: 8},
			}},
		},
		{
			name: "extract return error when heap is empty",
			actualResult: func() (interface{}, error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "na",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ele, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, ele)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMinHeapExtract(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, *MinHeap)
		expectedElement interface{}
		expectedResult  *MinHeap
		expectedError   error
	}{
		{
			name: "extract first element of the min heap",
			actualResult: func() (interface{}, error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 1,
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{2},
				indexes: map[interface{}]int{2: 0},
			}},
		},
		{
			name: "extract first element of the min heap two",
			actualResult: func() (interface{}, error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 7,
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{8, 9, 10},
				indexes: map[interface{}]int{8: 0, 9: 1, 10: 2},
			}},
		},
		{
			name: "extract first element of the min heap three",
			actualResult: func() (interface{}, error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 3, 2, 8, 5, 9, 6, 7)
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 0,
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{1, 2, 4, 3, 7, 8, 5, 9, 6},
				indexes: map[interface{}]int{1: 0, 2: 1, 4: 2, 3: 3, 7: 4, 8: 5, 5: 6, 9: 7, 6: 8},
			}},
		},
		{
			name: "extract return error when heap is empty",
			actualResult: func() (interface{}, error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := h.Extract()
				return e, err, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "na",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}(nil),
				indexes: make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ele, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, ele)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMaxHeapUpdate(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Heap)
		expectedResult Heap
		expectedError  error
	}{
		{
			name: "test decrease value in max heap",
			actualResult: func() (error, Heap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 6, 4, 2, 1)
				require.NoError(t, err)

				return h.Update(4, 0), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{6, 1, 2, 0},
				indexes:   map[interface{}]int{6: 0, 1: 1, 2: 2, 0: 3},
			}},
		},
		{
			name: "test increase value in max heap",
			actualResult: func() (error, Heap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 6, 4, 2, 1)
				require.NoError(t, err)

				return h.Update(4, 7), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{7, 6, 2, 1},
				indexes:   map[interface{}]int{7: 0, 6: 1, 2: 2, 1: 3},
			}},
		},
		{
			name: "test decrease value in min heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 4, 6)
				require.NoError(t, err)

				return h.Update(2, 0), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{0, 1, 4, 6},
				indexes:   map[interface{}]int{0: 0, 1: 1, 4: 2, 6: 3},
			}},
		},
		{
			name: "test increase value in min heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 4, 6)
				require.NoError(t, err)

				return h.Update(2, 7), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1, 6, 4, 7},
				indexes:   map[interface{}]int{1: 0, 6: 1, 4: 2, 7: 3},
			}},
		},
		{
			name: "test return error when heap is empty",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Update(2, 1), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "na",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
		{
			name: "test return error when prev type is different",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Update('a', 2), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when new type is different",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Update(1, 'a'), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when prev is not present in heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Update(2, 3), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: errors.New("2 not found in heap"),
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

type value struct{ val int }

func newValue(v int) *value { return &value{val: v} }

func (v *value) String() string {
	return fmt.Sprintf("%d", v.val)
}

type pointerValueComparator struct{}

func (vc *pointerValueComparator) Compare(one interface{}, two interface{}) (int, error) {
	return one.(*value).val - two.(*value).val, nil
}

type valueComparator struct{}

func (vc valueComparator) Compare(one interface{}, two interface{}) (int, error) {
	return one.(value).val - two.(value).val, nil
}

func TestMaxHeapUpdateFunc(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Heap)
		expectedResult Heap
		expectedError  error
	}{
		{
			name: "test decrease value in max heap",
			actualResult: func() (error, Heap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 6, 4, 2, 1)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 0
				}

				return h.UpdateFunc(4, uf), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{6, 1, 2, 0},
				indexes:   map[interface{}]int{6: 0, 1: 1, 2: 2, 0: 3},
			}},
		},
		{
			name: "test decrease pointer reference value in max heap",
			actualResult: func() (error, Heap) {
				f := &value{4}
				h, err := NewMaxHeap(&pointerValueComparator{}, &value{6}, f, &value{2}, &value{1})
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					e.(*value).val = 0
					return e
				}

				return h.UpdateFunc(f, uf), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "value",
				isMaxHeap: true,
				c:         &pointerValueComparator{},
				data:      []interface{}{&value{6}, &value{1}, &value{2}, &value{0}},
				indexes: map[interface{}]int{
					&value{1}: 1,
					&value{6}: 0,
					&value{0}: 3,
					&value{2}: 2,
				},
			}},
		},
		{
			name: "test decrease reference value in max heap",
			actualResult: func() (error, Heap) {
				f := value{4}
				h, err := NewMaxHeap(valueComparator{}, value{6}, f, value{2}, value{1})
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return value{0}
				}

				return h.UpdateFunc(f, uf), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "value",
				isMaxHeap: true,
				c:         valueComparator{},
				data:      []interface{}{value{6}, value{1}, value{2}, value{0}},
				indexes: map[interface{}]int{
					value{1}: 1,
					value{6}: 0,
					value{0}: 3,
					value{2}: 2,
				},
			}},
		},
		{
			name: "test increase value in max heap",
			actualResult: func() (error, Heap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 6, 4, 2, 1)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 7
				}

				return h.UpdateFunc(4, uf), h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{7, 6, 2, 1},
				indexes:   map[interface{}]int{7: 0, 6: 1, 2: 2, 1: 3},
			}},
		},
		{
			name: "test decrease value in min heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 4, 6)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 0
				}

				return h.UpdateFunc(2, uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{0, 1, 4, 6},
				indexes:   map[interface{}]int{0: 0, 1: 1, 4: 2, 6: 3},
			}},
		},
		{
			name: "test increase value in min heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 4, 6)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 7
				}

				return h.UpdateFunc(2, uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1, 6, 4, 7},
				indexes:   map[interface{}]int{1: 0, 6: 1, 4: 2, 7: 3},
			}},
		},
		{
			name: "test return error when heap is empty",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 1
				}

				return h.UpdateFunc(2, uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "na",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
		{
			name: "test return error when prev type is different",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 2
				}

				return h.UpdateFunc('a', uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when new type is different",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 'a'
				}

				return h.UpdateFunc(1, uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when prev is not present in heap",
			actualResult: func() (error, Heap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				uf := func(e interface{}) interface{} {
					return 3
				}

				return h.UpdateFunc(2, uf), h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: false,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
			expectedError: errors.New("2 not found in heap"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if reflect.TypeOf(res).String() == "*heap.MaxHeap" {
				assert.Equal(t, testCase.expectedResult.(*MaxHeap).data, res.(*MaxHeap).data)
				assert.Equal(t, testCase.expectedResult.(*MaxHeap).isMaxHeap, res.(*MaxHeap).isMaxHeap)
				assert.Equal(t, testCase.expectedResult.(*MaxHeap).c, res.(*MaxHeap).c)
				assert.Equal(t, testCase.expectedResult.(*MaxHeap).typeURL, res.(*MaxHeap).typeURL)
				em := testCase.expectedResult.(*MaxHeap).indexes
				am := testCase.expectedResult.(*MaxHeap).indexes
				for k := range em {
					assert.Equal(t, em[k], am[k])
				}

			} else {
				assert.Equal(t, testCase.expectedResult.(*MinHeap).data, res.(*MinHeap).data)
				assert.Equal(t, testCase.expectedResult.(*MinHeap).isMaxHeap, res.(*MinHeap).isMaxHeap)
				assert.Equal(t, testCase.expectedResult.(*MinHeap).c, res.(*MinHeap).c)
				assert.Equal(t, testCase.expectedResult.(*MinHeap).typeURL, res.(*MinHeap).typeURL)

				em := testCase.expectedResult.(*MinHeap).indexes
				am := testCase.expectedResult.(*MinHeap).indexes
				for k := range em {
					assert.Equal(t, em[k], am[k])
				}
			}

		})
	}
}

func TestMaxHeapDelete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MaxHeap)
		expectedResult *MaxHeap
		expectedError  error
	}{
		{
			name: "delete first element of the max heap",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{1},
				indexes:   map[interface{}]int{1: 0},
			}},
		},
		{
			name: "delete first element of the max heap two",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{9, 7, 8},
				indexes:   map[interface{}]int{9: 0, 7: 1, 8: 2},
			}},
		},
		{
			name: "delete first element of the max heap three",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 9, 8, 5, 6, 7, 1, 4, 0, 3, 2)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "int",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}{8, 7, 5, 6, 2, 1, 4, 0, 3},
				indexes:   map[interface{}]int{8: 0, 7: 1, 5: 2, 6: 3, 2: 4, 1: 5, 4: 6, 0: 7, 3: 8},
			}},
		},
		{
			name: "delete return error when heap is empty",
			actualResult: func() (error, *MaxHeap) {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap{&binaryHeap{
				typeURL:   "na",
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []interface{}(nil),
				indexes:   make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMinHeapDelete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MinHeap)
		expectedResult *MinHeap
		expectedError  error
	}{
		{
			name: "delete first element of the min Heap",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{2},
				indexes: map[interface{}]int{2: 0},
			}},
		},
		{
			name: "delete first element of the min Heap two",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{8, 9, 10},
				indexes: map[interface{}]int{8: 0, 9: 1, 10: 2},
			}},
		},
		{
			name: "delete first element of the min Heap three",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 3, 2, 8, 5, 9, 6, 7)
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "int",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}{1, 2, 4, 3, 7, 8, 5, 9, 6},
				indexes: map[interface{}]int{1: 0, 2: 1, 4: 2, 3: 3, 7: 4, 8: 5, 5: 6, 9: 7, 6: 8},
			}},
		},
		{
			name: "delete return error when heap is empty",
			actualResult: func() (error, *MinHeap) {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				err = h.Delete()
				return err, h
			},
			expectedResult: &MinHeap{&binaryHeap{
				typeURL: "na",
				c:       comparator.NewIntegerComparator(),
				data:    []interface{}(nil),
				indexes: make(map[interface{}]int),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMaxHeapIteratorHasNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test has next return false for empty Heap",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Iterator().HasNext()
			},
		},
		{
			name: "test has next return true for non empty Heap",
			actualResult: func() bool {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Iterator().HasNext()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMinHeapIteratorHasNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test has next return false for empty Heap",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Iterator().HasNext()
			},
		},
		{
			name: "test has next return true for non empty Heap",
			actualResult: func() bool {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Iterator().HasNext()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMaxHeapIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test get nil for empty Heap",
			actualResult: func() interface{} {
				h, err := NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Iterator().Next()
			},
		},
		{
			name: "test get all items from Heap one",
			actualResult: func() interface{} {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Iterator().Next()
			},
			expectedResult: 1,
		},
		{
			name: "test get all items from Heap two",
			actualResult: func() interface{} {
				h, err := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				i := h.Iterator()

				var res []interface{}
				for i.HasNext() {
					res = append(res, i.Next())
				}

				return res
			},
			expectedResult: []interface{}{4, 2, 3, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestMinHeapIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test get nil for empty Heap",
			actualResult: func() interface{} {
				h, err := NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return h.Iterator().Next()
			},
		},
		{
			name: "test get all items from Heap one",
			actualResult: func() interface{} {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return h.Iterator().Next()
			},
			expectedResult: 1,
		},
		{
			name: "test get all items from Heap two",
			actualResult: func() interface{} {
				h, err := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				i := h.Iterator()

				var res []interface{}
				for i.HasNext() {
					res = append(res, i.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
