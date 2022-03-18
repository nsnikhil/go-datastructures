package queue

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewLinkedQueue(t *testing.T) {
	actualResult := NewLinkedQueue[int]()

	ll := list.NewLinkedList[int]()

	expectedResult := &LinkedQueue[int]{
		ll: ll,
	}

	assert.Equal(t, expectedResult, actualResult)
}

func TestLinkedQueueAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue[int]
		expectedResult func() Queue[int]
		expectedError  error
	}{
		{
			name: "test add item to queue",
			actualResult: func() Queue[int] {
				q := NewLinkedQueue[int]()

				q.Add(1)
				return q
			},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)

				return &LinkedQueue[int]{ll: ll}
			},
		},
		{
			name: "test add multiple items to queue",
			actualResult: func() Queue[int] {
				q := NewLinkedQueue[int]()

				q.Add(1)

				q.Add(2)

				q.Add(3)

				q.Add(4)
				return q

			},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1, 2, 3, 4)

				return &LinkedQueue[int]{ll: ll}
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

func TestLinkedQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int, error, Queue[int])
		expectedResult  func() Queue[int]
		expectedElement []int
		expectedError   error
	}{
		{
			name: "test remove item from queue",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewLinkedQueue[int]()

				q.Add(1)

				e, err := q.Remove()

				return []int{e}, err, q
			},
			expectedElement: []int{1},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(1)
				ll.Clear()

				return &LinkedQueue[int]{ll: ll}
			},
		},
		{
			name: "test remove multiple items",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewLinkedQueue[int]()

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				var res []int

				e, err := q.Remove()
				if err != nil {
					return res, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return res, err, q
				}
				res = append(res, e)

				e, err = q.Remove()
				if err != nil {
					return res, err, q
				}
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []int{1, 2, 3},
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList(4)

				return &LinkedQueue[int]{ll: ll}
			},
		},
		{
			name: "test return error when queue is empty",
			actualResult: func() ([]int, error, Queue[int]) {
				q := NewLinkedQueue[int]()

				_, err := q.Remove()

				return []int(nil), err, q
			},
			expectedElement: []int(nil),
			expectedResult: func() Queue[int] {
				ll := list.NewLinkedList[int]()
				ll.Clear()

				return &LinkedQueue[int]{ll: ll}
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

func TestLinkedQueuePeek(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (int, error)
		expectedElement int
		expectedError   error
	}{
		{
			name: "test peek item",
			actualResult: func() (int, error) {
				q := NewLinkedQueue[int]()

				q.Add(1)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test peek item two",
			actualResult: func() (int, error) {
				q := NewLinkedQueue[int]()

				q.Add(1)
				q.Add(2)
				q.Add(3)
				q.Add(4)

				e, err := q.Peek()

				return e, err
			},
			expectedElement: 1,
		},
		{
			name: "test peek item from empty queue",
			actualResult: func() (int, error) {
				q := NewLinkedQueue[int]()

				e, err := q.Peek()

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

func TestLinkedQueueEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when queue is empty",
			actualResult: func() bool {
				q := NewLinkedQueue[int]()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when queue is not empty",
			actualResult: func() bool {
				q := NewLinkedQueue[int]()

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

func TestLinkedQueueCount(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "test return 0 when queue is empty",
			actualResult: func() int64 {
				q := NewLinkedQueue[int]()

				return q.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return count as 2",
			actualResult: func() int64 {
				q := NewLinkedQueue[int]()

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

func TestLinkedQueueClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true after clear",
			actualResult: func() bool {
				q := NewLinkedQueue[int]()

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true after clear two",
			actualResult: func() bool {
				q := NewLinkedQueue[int]()

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
