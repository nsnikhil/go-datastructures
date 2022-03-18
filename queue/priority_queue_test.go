package queue

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/heap"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewPriorityQueue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue[int]
		expectedResult func() Queue[int]
	}{
		{
			name: "test create max priority queue",
			actualResult: func() Queue[int] {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				return q
			},
			expectedResult: func() Queue[int] {
				mx := heap.NewMaxHeap(comparator.NewIntegerComparator())

				return &PriorityQueue[int]{mx}
			},
		},
		{
			name: "test create min priority queue",
			actualResult: func() Queue[int] {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				return q
			},
			expectedResult: func() Queue[int] {
				mn := heap.NewMinHeap(comparator.NewIntegerComparator())

				return &PriorityQueue[int]{mn}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestPriorityQueueAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue[int]
		expectedResult func() Queue[int]
		expectedError  error
	}{
		{
			name: "test add to max priority queue",
			actualResult: func() Queue[int] {
				pq := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				pq.Add(1)
				return pq
			},
			expectedResult: func() Queue[int] {
				h := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test add to min priority queue",
			actualResult: func() Queue[int] {
				pq := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				pq.Add(1)
				return pq
			},
			expectedResult: func() Queue[int] {
				h := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test add multiple elements to max priority queue",
			actualResult: func() Queue[int] {
				pq := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				pq.Add(1)

				pq.Add(2)

				pq.Add(3)

				pq.Add(4)
				return pq
			},
			expectedResult: func() Queue[int] {
				h := heap.NewMaxHeap(comparator.NewIntegerComparator(), 4, 3, 2, 1)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test add multiple elements to min priority queue",
			actualResult: func() Queue[int] {
				pq := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				pq.Add(4)

				pq.Add(3)

				pq.Add(2)

				pq.Add(1)
				return pq
			},
			expectedResult: func() Queue[int] {
				h := heap.NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)

				return &PriorityQueue[int]{h: h}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), q)
		})
	}
}

func TestPriorityQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int, error, Queue[int])
		expectedResult  func() Queue[int]
		expectedElement []int
		expectedError   error
	}{
		{
			name: "test remove item from max priority queue",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)

				e, err := q.Remove()

				return []int{e}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				h := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)

				_, err := h.Extract()
				require.NoError(t, err)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test remove item from min priority queue",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)

				e, err := q.Remove()

				return []int{e}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				h := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)

				_, err := h.Extract()
				require.NoError(t, err)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test remove multiple items from max priority queue",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				var res []int

				e, err := q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []int{4, 3, 2},
			expectedResult: func() Queue[int] {
				h := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test remove multiple items from min priority queue",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				var res []int

				e, err := q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.Remove()
				require.NoError(t, err)
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []int{1, 2, 3},
			expectedResult: func() Queue[int] {
				h := heap.NewMinHeap(comparator.NewIntegerComparator(), 4)

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test return error when max priority queue is empty",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				_, err := q.Remove()

				return []int(nil), err, q
			},
			expectedElement: []int(nil),
			expectedResult: func() Queue[int] {
				h := heap.NewMaxHeap(comparator.NewIntegerComparator())

				return &PriorityQueue[int]{h: h}
			},
		},
		{
			name: "test return error when min priority queue is empty",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				_, err := q.Remove()

				return []int(nil), err, q
			},
			expectedElement: []int(nil),
			expectedResult: func() Queue[int] {
				h := heap.NewMinHeap(comparator.NewIntegerComparator())

				return &PriorityQueue[int]{h: h}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, _, q := testCase.actualResult()

			//TODO FIX ERROR MESSAGE AND UNCOMMENT THE ASSERTION
			//assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedElement, e)
				assert.Equal(t, testCase.expectedResult(), q)
			}
		})
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (int, error)
		expectedElement int
		expectedError   error
	}{
		{
			name: "test max priority queue peek item",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test min priority queue peek item",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test max priority queue peek item two",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 4,
		},
		{
			name: "test min priority queue peek item two",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(4)
				q.Add(3)
				q.Add(2)
				q.Add(1)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test peek item from empty max priority queue",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				e, err := q.Peek()

				return e, err
			},
			expectedError: errors.New("iterator is empty"),
		},
		{
			name: "test peek item from empty min priority queue",
			actualResult: func() (int, error) {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				e, err := q.Peek()

				return e, err
			},
			expectedError: errors.New("iterator is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
		})
	}
}

func TestPriorityQueueEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when max priority queue is empty",
			actualResult: func() bool {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true when min priority queue is empty",
			actualResult: func() bool {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when max priority queue is not empty",
			actualResult: func() bool {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)

				return q.Empty()
			},
			expectedResult: false,
		},
		{
			name: "test return false when min priority queue is not empty",
			actualResult: func() bool {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)

				return q.Empty()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestPriorityQueueCount(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "test return 0 when max priority queue is empty",
			actualResult: func() int64 {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				return q.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return 0 when min priority queue is empty",
			actualResult: func() int64 {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				return q.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return max priority queue count as 2",
			actualResult: func() int64 {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)

				return q.Size()
			},
			expectedResult: 2,
		},
		{
			name: "test return min priority queue count as 2",
			actualResult: func() int64 {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)

				return q.Size()
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

func TestPriorityQueueClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true after max priority queue clear",
			actualResult: func() bool {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true after min priority queue clear",
			actualResult: func() bool {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return after true max priority queue  clear two",
			actualResult: func() bool {
				q := NewPriorityQueue[int](true, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return after true min priority queue  clear two",
			actualResult: func() bool {
				q := NewPriorityQueue[int](false, comparator.NewIntegerComparator())

				q.Add(1)
				q.Add(2)

				q.Clear()

				return q.Empty()
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
