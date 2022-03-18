package queue

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewDeque(t *testing.T) {
	actualResult := NewDeque[int]()

	lq := NewLinkedQueue[int]()

	expectedResult := &Deque[int]{lq}

	assert.Equal(t, expectedResult, actualResult)
}

func TestDequeAddFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue[int]
		expectedResult func() Queue[int]
		expectedError  error
	}{
		{
			name: "test add item to deque",
			actualResult: func() Queue[int] {
				q := NewDeque[int]()

				q.AddFirst(1)
				return q
			},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int](1)

				lq := &LinkedQueue[int]{ll: ll}

				return &Deque[int]{lq}
			},
		},
		{
			name: "test add multiple items to deque",
			actualResult: func() Queue[int] {
				q := NewDeque[int]()

				q.AddFirst(1)

				q.AddFirst(2)

				q.AddFirst(3)

				q.AddFirst(4)
				return q

			},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int](4, 3, 2, 1)

				lq := &LinkedQueue[int]{ll: ll}

				return &Deque[int]{lq}
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

func TestDequeRemoveLast(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int, error, Queue[int])
		expectedResult  func() Queue[int]
		expectedElement []int
		expectedError   error
	}{
		{
			name: "test remove item from deque",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewDeque[int]()

				q.Add(1)

				e, err := q.RemoveLast()

				return []int{e}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int](1)
				ll.Clear()

				lq := &LinkedQueue[int]{ll: ll}

				return &Deque[int]{lq}
			},
		},
		{
			name: "test remove multiple items",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewDeque[int]()

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				var res []int

				e, err := q.RemoveLast()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.RemoveLast()
				require.NoError(t, err)
				res = append(res, e)

				e, err = q.RemoveLast()
				require.NoError(t, err)
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []int{4, 3, 2},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int](1)

				lq := &LinkedQueue[int]{ll: ll}

				return &Deque[int]{lq}
			},
		},
		{
			name: "test return error when deque is empty",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewDeque[int]()

				_, err := q.RemoveLast()

				return []int(nil), err, q
			},
			expectedElement: []int(nil),
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int]()

				lq := &LinkedQueue[int]{ll: ll}

				return &Deque[int]{lq}
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

func TestDequePeekLast(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (int, error)
		expectedElement int
		expectedError   error
	}{
		{
			name: "test peek item",
			actualResult: func() (int, error) {
				q := NewDeque[int]()

				q.Add(1)

				e, err := q.PeekLast()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test peek item two",
			actualResult: func() (int, error) {
				q := NewDeque[int]()

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				e, err := q.PeekLast()

				return e, err
			},
			expectedElement: 4,
		},
		{
			name: "test peek item from empty deque",
			actualResult: func() (int, error) {
				q := NewDeque[int]()

				e, err := q.PeekLast()

				return e, err
			},
			expectedError: errors.New("list is empty"),
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
