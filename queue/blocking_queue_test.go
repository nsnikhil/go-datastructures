package queue

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateNewBlockingQueue(t *testing.T) {
	actualResult, err := NewBlockingQueue()
	require.NoError(t, err)

	lq, err := NewLinkedQueue()
	require.NoError(t, err)

	expectedResult := &BlockingQueue{lq}

	assert.Equal(t, expectedResult, actualResult)
}

func TestBlockingQueueRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement []interface{}
		expectedError   error
	}{
		{
			name: "test block remove until an element is added",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q Queue
				var e interface{}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				c := make(chan interface{})

				go func(q Queue, e interface{}, err error) {
					e, err = q.Remove()
					c <- e
				}(q, e, err)

				go func(q Queue) {
					time.AfterFunc(time.Millisecond*5, func() {
						require.NoError(t, q.Add(1))
					})
				}(q)

				return []interface{}{<-c}, err, q

			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
		{
			name: "test do not block when an element are already present added",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q Queue
				var e interface{}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				c := make(chan interface{})

				go func(q Queue, e interface{}, err error) {
					e, err = q.Remove()
					c <- e
				}(q, e, err)

				return []interface{}{<-c}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(2)
				require.NoError(t, err)

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q Queue
				var e interface{}
				var res []interface{}

				c := make(chan interface{})

				remove := func(q Queue, e interface{}, err error) {
					e, err = q.Remove()
					c <- e
				}

				add := func(q Queue, e interface{}) {
					time.AfterFunc(time.Millisecond*5, func() {
						require.NoError(t, q.Add(e))
					})
				}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				go remove(q, e, err)
				res = append(res, <-c)

				go remove(q, e, err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, err)
				res = append(res, <-c)

				go remove(q, e, err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []interface{}{1, 2, 3, 4},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, len(testCase.expectedElement), len(e))
				assert.Equal(t, testCase.expectedResult(), res)
			}
		})
	}
}

func TestBlockingQueueRemoveWithTimeout(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]interface{}, error, Queue)
		expectedResult  func() Queue
		expectedElement []interface{}
		expectedError   error
	}{
		{
			name: "test remove element with max wait of 10 milli sec",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q *BlockingQueue
				var e interface{}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				c := make(chan interface{})

				go func(q *BlockingQueue, e interface{}, err error) {
					e, err = q.RemoveWithTimeout(time.Millisecond * 10)
					c <- e
				}(q, e, err)

				go func(q Queue) {
					time.AfterFunc(time.Second, func() {
						require.NoError(t, q.Add(1))
					})
				}(q)

				return []interface{}{<-c}, err, q

			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList()
				require.NoError(t, err)

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
		{
			name: "test remove element fails when max wait is 1 milli sec",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q *BlockingQueue
				var e interface{}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				c := make(chan interface{})

				go func(q *BlockingQueue, e interface{}, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}(q, e, &err)

				go func(q Queue) {
					time.AfterFunc(time.Second, func() {
						require.NoError(t, q.Add(1))
					})
				}(q)

				return []interface{}{<-c}, err, q

			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
			expectedError: errors.New("timed out"),
		},
		{
			name: "test do not block when an element are already present added",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q *BlockingQueue
				var e interface{}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				c := make(chan interface{})

				go func(q *BlockingQueue, e interface{}, err error) {
					e, err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}(q, e, err)

				return []interface{}{<-c}, err, q
			},
			expectedElement: []interface{}{1},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(2)
				require.NoError(t, err)

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements with max wait of 10 milli sec",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q *BlockingQueue
				var e interface{}
				var res []interface{}

				c := make(chan interface{})

				remove := func(q *BlockingQueue, e interface{}, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond * 10)
					c <- e
				}

				add := func(q *BlockingQueue, e interface{}) {
					time.AfterFunc(time.Millisecond*5, func() {
						require.NoError(t, q.Add(e))
					})
				}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []interface{}{1, 2, 3, 4},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
		},
		{
			name: "test remove multiple elements fails when max wait is 1 milli sec",
			actualResult: func() ([]interface{}, error, Queue) {
				var err error
				var q *BlockingQueue
				var e interface{}
				var res []interface{}

				c := make(chan interface{})

				remove := func(q *BlockingQueue, e interface{}, err *error) {
					e, *err = q.RemoveWithTimeout(time.Millisecond)
					c <- e
				}

				add := func(q *BlockingQueue, e interface{}) {
					time.AfterFunc(time.Second, func() {
						require.NoError(t, q.Add(e))
					})
				}

				q, err = NewBlockingQueue()
				require.NoError(t, err)

				require.NoError(t, q.Add(1))
				require.NoError(t, q.Add(2))

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				go add(q, 4)
				go add(q, 3)

				go remove(q, e, &err)
				res = append(res, <-c)

				go remove(q, e, &err)
				res = append(res, <-c)

				return res, err, q
			},
			expectedElement: []interface{}{},
			expectedResult: func() Queue {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &BlockingQueue{&LinkedQueue{ll: ll}}
			},
			expectedError: errors.New("timed out"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, len(testCase.expectedElement), len(e))
				assert.Equal(t, testCase.expectedResult(), res)
			}
		})
	}
}
