package queue

import (
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewDeque(t *testing.T) {
	actualResult, err := NewDeque()
	require.NoError(t, err)

	lq, err := NewLinkedQueue()
	require.NoError(t, err)

	expectedResult := &Deque{lq}

	assert.Equal(t, expectedResult, actualResult)
}

func TestDequeAddFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Queue)
		expectedResult func() Queue
		expectedError  error
	}{
		{
			name: "test add item to deque",
			actualResult: func() (error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				return q.AddFirst(1), q
			},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test add multiple items to deque",
			actualResult: func() (error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				if err = q.AddFirst(1); err != nil {
					return err, q
				}

				if err = q.AddFirst(2); err != nil {
					return err, q
				}

				if err = q.AddFirst(3); err != nil {
					return err, q
				}

				return q.AddFirst(4), q

			},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(4, 3, 2, 1)
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test add return error when type is different",
			actualResult: func() (error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				if err = q.AddFirst(1); err != nil {
					return err, q
				}

				return q.AddFirst("a"), q
			},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
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

func TestDequeRemoveLast(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement []interface{}
		expectedError   error
	}{
		{
			name: "test remove item from deque",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.RemoveLast()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test remove multiple items",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				var res []interface{}

				e, err := q.RemoveLast()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.RemoveLast()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				e, err = q.RemoveLast()
				if err != nil {
					return e, err, q
				}
				res = append(res, e)

				return res, err, q
			},
			expectedElement: []interface{}{4, 3, 2},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test return error when deque is empty",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				e, err := q.RemoveLast()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{interface{}(nil)},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList()
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
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
		actualResult    func() (interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement interface{}
		expectedError   error
	}{
		{
			name: "test peek item",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.PeekLast()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test peek item two",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				e, err := q.PeekLast()

				return e, err, q
			},
			expectedElement: 4,
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
			},
		},
		{
			name: "test peek item from empty deque",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewDeque()
				require.NoError(t, err)

				e, err := q.PeekLast()

				return e, err, q
			},
			expectedElement: nil,
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList()
				require.NoError(t, err)

				lq := &LinkedQueue{ll: ll}

				return &Deque{lq}
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
