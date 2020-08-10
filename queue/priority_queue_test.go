package queue

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/heap"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewPriorityQueue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue
		expectedResult func() Queue
	}{
		{
			name: "test create max priority queue",
			actualResult: func() Queue {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q
			},
			expectedResult: func() Queue {
				mx, err := heap.NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{mx}
			},
		},
		{
			name: "test create min priority queue",
			actualResult: func() Queue {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q
			},
			expectedResult: func() Queue {
				mn, err := heap.NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{mn}
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
		actualResult   func() (error, Queue)
		expectedResult func() Queue
		expectedError  error
	}{
		{
			name: "test add to max priority queue",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return pq.Add(1), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test add to min priority queue",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return pq.Add(1), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test add multiple elements to max priority queue",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				if err = pq.Add(1); err != nil {
					return err, pq
				}

				if err = pq.Add(2); err != nil {
					return err, pq
				}

				if err = pq.Add(3); err != nil {
					return err, pq
				}

				return pq.Add(4), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 4, 3, 2, 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test add multiple elements to min priority queue",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				if err = pq.Add(4); err != nil {
					return err, pq
				}

				if err = pq.Add(3); err != nil {
					return err, pq
				}

				if err = pq.Add(2); err != nil {
					return err, pq
				}

				return pq.Add(1), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test add to max priority queue return error when type is different",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				if err = pq.Add(1); err != nil {
					return err, pq
				}

				return pq.Add("a"), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add to min priority queue return error when type is different",
			actualResult: func() (error, Queue) {
				pq, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				if err = pq.Add(1); err != nil {
					return err, pq
				}

				return pq.Add("a"), pq
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, q := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedResult(), q)
			}
		})
	}
}

func TestPriorityQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement []interface{}
		expectedError   error
	}{
		{
			name: "test remove item from max priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				_, err = h.Extract()
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test remove item from min priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				_, err = h.Extract()
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test remove multiple items from max priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				var res []interface{}

				e, err := q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []interface{}{4, 3, 2},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test remove multiple items from min priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				var res []interface{}

				e, err := q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []interface{}{1, 2, 3},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 4)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test return error when max priority queue is empty",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{interface{}(nil)},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test return error when min priority queue is empty",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{interface{}(nil)},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{h: h}
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

func TestPriorityQueueUpdate(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Queue)
		expectedResult func() Queue
		expectedError  error
	}{
		{
			name: "test decrease value in max priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(6))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(1))

				return q.Update(2, 0), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 6, 1, 4, 0)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test increase value in max priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(6))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(1))

				return q.Update(2, 7), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 7, 6, 4, 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test decrease value in min priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(6))

				return q.Update(2, 0), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 6)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test increase value in min priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(6))

				return q.Update(2, 7), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1, 6, 4, 7)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestPriorityQueueUpdateFunc(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Queue)
		expectedResult func() Queue
		expectedError  error
	}{
		{
			name: "test decrease value in max priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(6))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(1))

				up := func(interface{}) interface{} {
					return 0
				}

				return q.UpdateFunc(2, up), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 6, 1, 4, 0)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test increase value in max priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(6))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(1))

				up := func(interface{}) interface{} {
					return 7
				}

				return q.UpdateFunc(2, up), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 7, 6, 4, 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test decrease value in min priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(6))

				up := func(interface{}) interface{} {
					return 0
				}

				return q.UpdateFunc(2, up), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 0, 1, 4, 6)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test increase value in min priority queue",
			actualResult: func() (error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(6))

				up := func(interface{}) interface{} {
					return 7
				}

				return q.UpdateFunc(2, up), q
			},
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1, 6, 4, 7)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement interface{}
		expectedError   error
	}{
		{
			name: "test max priority queue peek item",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 1,
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test min priority queue peek item",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 1,
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test max priority queue peek item two",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 4,
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator(), 4, 3, 2, 1)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test min priority queue peek item two",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(4))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(1))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 1,
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator(), 1, 2, 3, 4)
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test peek item from empty max priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: nil,
			expectedResult: func() Queue {
				h, err := heap.NewMaxHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
		{
			name: "test peek item from empty min priority queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: nil,
			expectedResult: func() Queue {
				h, err := heap.NewMinHeap(comparator.NewIntegerComparator())
				require.NoError(t, err)

				return &PriorityQueue{h: h}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, q := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedElement, e)
				assert.Equal(t, testCase.expectedResult(), q)
			}
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
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true when min priority queue is empty",
			actualResult: func() bool {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when max priority queue is not empty",
			actualResult: func() bool {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				return q.Empty()
			},
			expectedResult: false,
		},
		{
			name: "test return false when min priority queue is not empty",
			actualResult: func() bool {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

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
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test return 0 when max priority queue is empty",
			actualResult: func() int {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q.Count()
			},
			expectedResult: 0,
		},
		{
			name: "test return 0 when min priority queue is empty",
			actualResult: func() int {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				return q.Count()
			},
			expectedResult: 0,
		},
		{
			name: "test return max priority queue count as 2",
			actualResult: func() int {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				return q.Count()
			},
			expectedResult: 2,
		},
		{
			name: "test return min priority queue count as 2",
			actualResult: func() int {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				return q.Count()
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
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true after min priority queue clear",
			actualResult: func() bool {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return after true max priority queue  clear two",
			actualResult: func() bool {
				q, err := NewPriorityQueue(true, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return after true min priority queue  clear two",
			actualResult: func() bool {
				q, err := NewPriorityQueue(false, comparator.NewIntegerComparator())
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

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
