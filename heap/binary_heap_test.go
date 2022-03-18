package heap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMaxHeap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *MaxHeap[int]
		expectedResult *MaxHeap[int]
		expectedError  error
	}{
		{
			name: "test create empty max heap",
			actualResult: func() *MaxHeap[int] {
				return NewMaxHeap[int](comparator.NewIntegerComparator())
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
			}},
		},
		{
			name: "test create heap of one element",
			actualResult: func() *MaxHeap[int] {
				return NewMaxHeap[int](comparator.NewIntegerComparator(), 1)
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{1},
			}},
		},
		{
			name: "test create max heap with multiple elements",
			actualResult: func() *MaxHeap[int] {
				return NewMaxHeap[int](comparator.NewIntegerComparator(), 1, 2, 3, 4)
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{4, 2, 3, 1},
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
		})
	}
}

func TestCreateMinHeap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *MinHeap[int]
		expectedResult *MinHeap[int]
		expectedError  error
	}{
		{
			name: "test create empty min heap",
			actualResult: func() *MinHeap[int] {
				return NewMinHeap[int](comparator.NewIntegerComparator())
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c: comparator.NewIntegerComparator(),
			}},
		},
		{
			name: "test create min heap with one element",
			actualResult: func() *MinHeap[int] {
				return NewMinHeap[int](comparator.NewIntegerComparator(), 1)
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{1},
			}},
		},
		{
			name: "test create heap with many elements",
			actualResult: func() *MinHeap[int] {
				return NewMinHeap[int](comparator.NewIntegerComparator(), 4, 3, 2, 1)
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{1, 3, 2, 4},
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
		})
	}
}

func TestMaxHeapAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MaxHeap[int])
		expectedResult *MaxHeap[int]
		expectedError  error
	}{
		{
			name: "test heap add one element",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				h.Add(10)
				return nil, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{10},
			}},
		},
		{
			name: "test add will heapify one element",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				h.Add(100, 40, 60, 80)
				h.Add(100, 40, 60, 80)
				return nil, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{100, 100, 60, 80, 80, 40, 60, 40},
			}},
		},
		{
			name: "test add will heapify two element",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				h.Add(100, 40, 60, 120)
				return nil, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{120, 100, 60, 40},
			}},
		},
		{
			name: "test add will heapify two element two",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				h.Add(100, 110, 120)
				return nil, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{120, 100, 110},
			}},
		},
		{
			name: "test add will heapify all element",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				h.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
				return nil, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{9, 8, 5, 6, 7, 1, 4, 0, 3, 2},
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMinHeapAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MinHeap[int])
		expectedResult *MinHeap[int]
		expectedError  error
	}{
		{
			name: "test heap add one element",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				h.Add(10)
				return nil, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{10},
			}},
		},
		{
			name: "test add will heapify one element",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				h.Add(10, 40, 60, 20)
				return nil, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{10, 20, 60, 40},
			}},
		},
		{
			name: "test add will heapify two element",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				h.Add(20, 40, 60, 10)
				return nil, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{10, 20, 60, 40},
			}},
		},
		{
			name: "test add will heapify two element two",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				h.Add(30, 20, 10)
				return nil, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{10, 30, 20},
			}},
		},
		{
			name: "test add will heapify all element",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				h.Add(9, 8, 7, 6, 5, 4, 3, 2, 1, 0)
				return nil, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{0, 1, 4, 3, 2, 8, 5, 9, 6, 7},
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

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
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				return h.IsEmpty()
			},
		},
		{
			name: "return true when Heap is empty",
			actualResult: func() bool {
				h := NewMaxHeap(comparator.NewIntegerComparator())

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
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				return h.IsEmpty()
			},
		},
		{
			name: "return true when Heap is empty",
			actualResult: func() bool {
				h := NewMinHeap(comparator.NewIntegerComparator())

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
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return true when Heap is empty after Clear two",
			actualResult: func() bool {
				h := NewMaxHeap(comparator.NewStringComparator(), "a", "b")

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
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				h.Clear()

				return h.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return true when Heap is empty after Clear two",
			actualResult: func() bool {
				h := NewMinHeap(comparator.NewStringComparator(), "a", "b")

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
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "get Size of empty Heap as 0",
			actualResult: func() int64 {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				return h.Size()
			},
		},
		{
			name: "get Size of empty Heap as 2",
			actualResult: func() int64 {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)

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
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "get Size of empty Heap as 0",
			actualResult: func() int64 {
				h := NewMinHeap(comparator.NewIntegerComparator())

				return h.Size()
			},
		},
		{
			name: "get Size of empty Heap as 2",
			actualResult: func() int64 {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)

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
		actualResult    func() (int, error, *MaxHeap[int])
		expectedElement int
		expectedResult  *MaxHeap[int]
		expectedError   error
	}{
		{
			name: "extract first element of the max heap",
			actualResult: func() (int, error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 2,
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{1},
			}},
		},
		{
			name: "extract first element of the max heap two",
			actualResult: func() (int, error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 10,
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{9, 7, 8},
			}},
		},
		{
			name: "extract first element of the max heap three",
			actualResult: func() (int, error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 9, 8, 5, 6, 7, 1, 4, 0, 3, 2)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 9,
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{8, 7, 5, 6, 2, 1, 4, 0, 3},
			}},
		},
		{
			name: "extract return error when heap is empty",
			actualResult: func() (int, error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				e, err := h.Extract()
				return e, err, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int(nil),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ele, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
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
		actualResult    func() (int, error, *MinHeap[int])
		expectedElement int
		expectedResult  *MinHeap[int]
		expectedError   error
	}{
		{
			name: "extract first element of the min heap",
			actualResult: func() (int, error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 1,
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{2},
			}},
		},
		{
			name: "extract first element of the min heap two",
			actualResult: func() (int, error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 7,
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{8, 9, 10},
			}},
		},
		{
			name: "extract first element of the min heap three",
			actualResult: func() (int, error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 3, 2, 8, 5, 9, 6, 7)

				e, err := h.Extract()
				return e, err, h
			},
			expectedElement: 0,
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{1, 2, 4, 3, 7, 8, 5, 9, 6},
			}},
		},
		{
			name: "extract return error when heap is empty",
			actualResult: func() (int, error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				e, err := h.Extract()
				return e, err, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int(nil),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ele, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, ele)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

type value struct{ val int }

func newValue(v int) *value { return &value{val: v} }

func (v *value) String() string {
	return fmt.Sprintf("%d", v.val)
}

type pointerValueComparator struct{}

func (vc *pointerValueComparator) Compare(one *value, two *value) int {
	return one.val - two.val
}

type valueComparator struct{}

func (vc valueComparator) Compare(one value, two value) int {
	return one.val - two.val
}

func TestMaxHeapDelete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MaxHeap[int])
		expectedResult *MaxHeap[int]
		expectedError  error
	}{
		{
			name: "delete first element of the max heap",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{1},
			}},
		},
		{
			name: "delete first element of the max heap two",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{9, 7, 8},
			}},
		},
		{
			name: "delete first element of the max heap three",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 9, 8, 5, 6, 7, 1, 4, 0, 3, 2)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int{8, 7, 5, 6, 2, 1, 4, 0, 3},
			}},
		},
		{
			name: "delete return error when heap is empty",
			actualResult: func() (error, *MaxHeap[int]) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				err := h.Delete()
				return err, h
			},
			expectedResult: &MaxHeap[int]{&binaryHeap[int]{
				isMaxHeap: true,
				c:         comparator.NewIntegerComparator(),
				data:      []int(nil),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

			if res != nil {
				assert.Equal(t, testCase.expectedResult.binaryHeap, res.binaryHeap)
			}
		})
	}
}

func TestMinHeapDelete(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *MinHeap[int])
		expectedResult *MinHeap[int]
		expectedError  error
	}{
		{
			name: "delete first element of the min Heap",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{2},
			}},
		},
		{
			name: "delete first element of the min Heap two",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 10, 9, 8, 7)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{8, 9, 10},
			}},
		},
		{
			name: "delete first element of the min Heap three",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 3, 2, 8, 5, 9, 6, 7)

				err := h.Delete()
				return err, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int{1, 2, 4, 3, 7, 8, 5, 9, 6},
			}},
		},
		{
			name: "delete return error when heap is empty",
			actualResult: func() (error, *MinHeap[int]) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				err := h.Delete()
				return err, h
			},
			expectedResult: &MinHeap[int]{&binaryHeap[int]{
				c:    comparator.NewIntegerComparator(),
				data: []int(nil),
			}},
			expectedError: errors.New("heap is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

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
				h := NewMaxHeap(comparator.NewIntegerComparator())

				return h.Iterator().HasNext()
			},
		},
		{
			name: "test has next return true for non empty Heap",
			actualResult: func() bool {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1)

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
				h := NewMinHeap(comparator.NewIntegerComparator())

				return h.Iterator().HasNext()
			},
		},
		{
			name: "test has next return true for non empty Heap",
			actualResult: func() bool {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1)

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
		actualResult   func() ([]int, error)
		expectedResult []int
		expectedError  error
	}{
		{
			name: "test get nil for empty Heap",
			actualResult: func() ([]int, error) {
				h := NewMaxHeap(comparator.NewIntegerComparator())

				_, err := h.Iterator().Next()

				return []int(nil), err
			},
			expectedError: errors.New("iterator is empty"),
		},
		{
			name: "test get all items from Heap one",
			actualResult: func() ([]int, error) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1)

				v, err := h.Iterator().Next()

				return []int{v}, err
			},
			expectedResult: []int{1},
		},
		{
			name: "test get all items from Heap two",
			actualResult: func() ([]int, error) {
				h := NewMaxHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				i := h.Iterator()

				var res []int
				for i.HasNext() {
					v, _ := i.Next()
					res = append(res, v)
				}

				return res, nil
			},
			expectedResult: []int{4, 2, 3, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestMinHeapIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]int, error)
		expectedResult []int
		expectedError  error
	}{
		{
			name: "test get nil for empty Heap",
			actualResult: func() ([]int, error) {
				h := NewMinHeap(comparator.NewIntegerComparator())

				_, err := h.Iterator().Next()

				return []int(nil), err
			},
			expectedError: errors.New("iterator is empty"),
		},
		{
			name: "test get all items from Heap one",
			actualResult: func() ([]int, error) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1)

				v, err := h.Iterator().Next()

				return []int{v}, err
			},
			expectedResult: []int{1},
		},
		{
			name: "test get all items from Heap two",
			actualResult: func() ([]int, error) {
				h := NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				i := h.Iterator()

				var res []int
				for i.HasNext() {
					v, _ := i.Next()
					res = append(res, v)
				}

				return res, nil
			},
			expectedResult: []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
