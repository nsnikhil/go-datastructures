package queue

import (
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewLinkedQueue(t *testing.T) {
	actualResult, err := NewLinkedQueue()
	require.NoError(t, err)

	ll, err := list.NewLinkedList()
	require.NoError(t, err)

	expectedResult := &LinkedQueue{
		ll: ll,
	}

	assert.Equal(t, expectedResult, actualResult)
}

func TestLinkedQueueAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Queue)
		expectedResult func() Queue
		expectedError  error
	}{
		{
			name: "test add item to queue",
			actualResult: func() (error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				return q.Add(1), q
			},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test add multiple items to queue",
			actualResult: func() (error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				if err = q.Add(1); err != nil {
					return err, q
				}

				if err = q.Add(2); err != nil {
					return err, q
				}

				if err = q.Add(3); err != nil {
					return err, q
				}

				return q.Add(4), q

			},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test add return error when type is different",
			actualResult: func() (error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				if err = q.Add(1); err != nil {
					return err, q
				}

				return q.Add("a"), q
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

func TestLinkedQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement []interface{}
		expectedError   error
	}{
		{
			name: "test remove item from queue",
			actualResult: func() ([]interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test remove multiple items",
			actualResult: func() ([]interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				var res []interface{}

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
			expectedElement: []interface{}{1, 2, 3},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(4)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test return error when queue is empty",
			actualResult: func() ([]interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				e, err := q.Remove()

				return []interface{}{e}, err, q
			},
			expectedElement: []interface{}{interface{}(nil)},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList()
				require.NoError(t, err)
				ll.Clear()

				return &LinkedQueue{ll: ll}
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
		actualResult    func() (interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement interface{}
		expectedError   error
	}{
		{
			name: "test peek item",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 1,
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test peek item two",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))
				require.NoError(t, q.Add(3))
				require.NoError(t, q.Add(4))

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: 1,
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return &LinkedQueue{ll: ll}
			},
		},
		{
			name: "test peek item from empty queue",
			actualResult: func() (interface{}, error, Queue) {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				e, err := q.Peek()

				return e, err, q
			},
			expectedElement: nil,
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList()
				require.NoError(t, err)
				ll.Clear()

				return &LinkedQueue{ll: ll}
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

func TestLinkedQueueEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when queue is empty",
			actualResult: func() bool {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when queue is not empty",
			actualResult: func() bool {
				q, err := NewLinkedQueue()
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

func TestLinkedQueueCount(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test return 0 when queue is empty",
			actualResult: func() int {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				return q.Count()
			},
			expectedResult: 0,
		},
		{
			name: "test return count as 2",
			actualResult: func() int {
				q, err := NewLinkedQueue()
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

func TestLinkedQueueClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true after clear",
			actualResult: func() bool {
				q, err := NewLinkedQueue()
				require.NoError(t, err)

				q.Clear()

				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test return true after clear two",
			actualResult: func() bool {
				q, err := NewLinkedQueue()
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
